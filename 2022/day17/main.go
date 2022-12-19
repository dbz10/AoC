package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func (c coord) plus(other coord) coord {
	return coord{c.x + other.x, c.y + other.y}
}

type block struct {
	origin  coord   // bottom left segment of the square which encloses the block. does not have to contain a rock
	offsets []coord // displacement to all the other segments in the block
}

type state struct {
	// blockStateEncoding int
	gustIndex int
	nextRock  int
}

func (b block) String() string {
	base := fmt.Sprintf("Origin: %v, Positions: ", b.origin)
	for _, o := range b.offsets {
		base += fmt.Sprintf("%v, ", b.origin.plus(o))
	}
	return base
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rockTypes := []block{
		{
			origin:  coord{0, 0},
			offsets: []coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		},
		{
			origin:  coord{0, 0},
			offsets: []coord{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		},
		{
			origin:  coord{0, 0},
			offsets: []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			origin:  coord{0, 0},
			offsets: []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			origin:  coord{0, 0},
			offsets: []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	wind := strings.Split(string(contents), "")
	var maxRockHeight, rockIterator, landedRocks int
	allRock := map[coord]bool{}
	for x := 0; x < 7; x++ {
		allRock[coord{x, 0}] = true
	}
	rock := newRock(rockIterator, maxRockHeight, rockTypes)
	rockIterator++

	for round := 0; round < 1000000; round++ {
		if landedRocks == 2022 {
			break
		}

		windDirection := wind[round%len(wind)]
		rock = tryMoveLR(rock, allRock, windDirection)
		rock, err = tryMoveDown(rock, allRock)

		if err != nil {
			// update the set with all rock locations and max rock height
			landedRocks++
			for _, offset := range rock.offsets {
				allRock[rock.origin.plus(offset)] = true
				maxRockHeight = max(rock.origin.plus(offset).y, maxRockHeight)
			}
			// check if we've reached a cycle. we only need to measure this once.

			// generate a new rock
			rock = newRock(rockIterator%len(rockTypes), maxRockHeight, rockTypes)
			rockIterator++
		}

	}
	fmt.Printf("Part One: After 2022 rocks have landed, the rock has reached height %d\n", maxRockHeight)

	// Part Two
	// Clearly we cannot simulate 1000000000000 falling rocks.
	// So we have to rely on the fact that after infinite time,
	// the pattern of rocks is guaranteed to converge into a cycle.

	// So to solve the problem we have to find
	// a) the time (measured in amount of rocks landed) to reach the beginning of a cycle
	// b) the length of the cycle (in rocks landed)

	// A state can be represented by three labels
	// 1. An encoding of the state of the top 4 rock layer
	// 2. The index in the "gust" array
	// 3. The next rock about to fall

	type memoInfo struct {
		rocksLanded, rockHeight int
	}
	stateMemoization := map[state]memoInfo{}

	rockIterator = 0
	landedRocks = 0
	maxRockHeight = 0
	allRock = map[coord]bool{}
	for x := 0; x < 7; x++ {
		allRock[coord{x, 0}] = true
	}
	rock = newRock(rockIterator, maxRockHeight, rockTypes)
	rockIterator++

	firstRepetitionSeen := false
	firstRepeatedState := state{8, 2}

	cycleMemoization := map[int]int{}
	cycleCounter := 0
	inCycle := false

	for round := 0; round < 1000000; round++ {
		windDirection := wind[round%len(wind)]
		rock = tryMoveLR(rock, allRock, windDirection)
		rock, err = tryMoveDown(rock, allRock)

		if err != nil {
			// update the set with all rock locations and max rock height
			landedRocks++
			for _, offset := range rock.offsets {
				allRock[rock.origin.plus(offset)] = true
				maxRockHeight = max(rock.origin.plus(offset).y, maxRockHeight)
			}

			currentState := state{
				// blockStateEncoding: rockState,
				gustIndex: round % len(wind),
				nextRock:  rockIterator % len(rockTypes),
			}

			if firstRepetitionSeen && currentState == firstRepeatedState {
				firstRepeatedState = currentState
				fmt.Printf("Since the last %d was seen, %d more rocks landed and %d additional height was added\n",
					currentState, landedRocks-stateMemoization[currentState].rocksLanded, maxRockHeight-stateMemoization[currentState].rockHeight)
				break
			}

			if stateMemoization[currentState].rocksLanded > 0 && !firstRepetitionSeen {
				firstRepeatedState = currentState
				firstRepetitionSeen = true
				inCycle = true
				fmt.Println("Numbers at first repetition")
				fmt.Println(maxRockHeight, "height")
				fmt.Println(landedRocks, "landed rocks")
			}

			if inCycle {
				cycleMemoization[cycleCounter] = maxRockHeight
				cycleCounter++
			}

			stateMemoization[currentState] = memoInfo{landedRocks, maxRockHeight}

			// generate a new rock
			rock = newRock(rockIterator%len(rockTypes), maxRockHeight, rockTypes)
			rockIterator++
		}
	}
	fmt.Println(cycleMemoization[1188] - cycleMemoization[0])
}

func render(allRock map[coord]bool, maxHeight int, currentRock block) string {
	aug := map[coord]int{}

	for key := range allRock {
		aug[key] = 1
	}
	for _, o := range currentRock.offsets {
		aug[currentRock.origin.plus(o)] = -1
	}

	output := ""
	for y := maxHeight + 7; y > 0; y-- {
		output += "|"
		for x := 0; x < 7; x++ {
			if aug[coord{x, y}] == 1 {
				output += "#"
			} else if aug[coord{x, y}] == -1 {
				output += "@"
			} else {
				output += "."
			}
		}
		output += "|\n"
	}
	output += "|"
	for x := 0; x < 7; x++ {
		output += "_"
	}
	output += "|\n"
	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newRock(rockIterator, maxRockHeight int, rockTypes []block) block {
	rock := rockTypes[rockIterator]
	rock.origin = coord{2, maxRockHeight + 4}
	return rock
}

func tryMoveLR(rock block, allRock map[coord]bool, gust string) block {
	var dx int
	if gust == ">" {
		dx = 1
	} else {
		dx = -1
	}

	for _, offset := range rock.offsets {
		newCoord := rock.origin.plus(offset)
		newCoord.x += dx
		if newCoord.x < 0 || newCoord.x > 6 || allRock[newCoord] {
			// then it cannot move left to right
			return rock
		}
	}
	rock.origin.x += dx
	return rock
}

func tryMoveDown(rock block, allRock map[coord]bool) (block, error) {
	for _, offset := range rock.offsets {
		nc := rock.origin.plus(offset)
		nc.y--
		if allRock[nc] || nc.y <= 0 {
			return rock, errors.New("hit the floor")
		}
	}
	rock.origin.y--
	return rock, nil
}
