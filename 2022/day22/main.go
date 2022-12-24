package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	puzzle := strings.Split(string(contents), "\n\n")[0]
	path := strings.Split(string(contents), "\n\n")[1]

	lines := strings.Split(puzzle, "\n")
	var lx, ly int
	var startX = 10000
	ly = len(lines)
	for _, l := range lines {
		if len(l) > lx {
			lx = len(l)
		}
	}

	grid := map[coord]string{}
	for y, l := range lines {
		for x, ch := range fmt.Sprintf("%-*s ", lx, l) {
			if y == 0 && string(ch) == "." && x < startX {
				startX = x
			}
			grid[coord{x, y}] = string(ch)
		}
	}

	nodes := map[coord]*node{}
	for key := range grid {
		nodes[key] = &node{loc: key}
	}
	for key := range grid {
		n := nodes[key]
		n.north = nodes[findNext(grid, key, coord{0, -1}, lx, ly)]
		n.west = nodes[findNext(grid, key, coord{-1, 0}, lx, ly)]
		n.east = nodes[findNext(grid, key, coord{1, 0}, lx, ly)]
		n.south = nodes[findNext(grid, key, coord{0, 1}, lx, ly)]
	}

	// well this is a little bit hacky but ok
	// they also don't come in pairs in the problem statement
	// but i'm gonna handle that because the alternative
	// is that it's a bit harder to parse this string...
	instructions := []instruction{}
	for _, sa := range strings.Split(path, "R") {
		v, err := strconv.Atoi(sa)
		if err == nil {
			instructions = append(instructions, instruction{v, "R"})
		} else {
			subsegs := strings.Split(sa, "L")
			for i := 0; i < len(subsegs)-1; i++ {
				v, _ = strconv.Atoi(subsegs[i])
				instructions = append(instructions, instruction{v, "L"})
			}
			v, _ = strconv.Atoi(subsegs[len(subsegs)-1])
			instructions = append(instructions, instruction{v, "R"})
		}
	}

	cur := coordWithOrientation{coord{startX, 0}, "east"}
	for i := 0; i < len(instructions)-1; i++ {
		move := instructions[i]
		cur.c = nodes[cur.c].walk(move.distance, cur.facing).loc
		cur.facing = cur.turn(move.turn)
	}
	// no turning on the last move
	cur.c = nodes[cur.c].walk(instructions[len(instructions)-1].distance, cur.facing).loc

	finalScore := 1000*(cur.c.y+1) + 4*(cur.c.x+1)
	if cur.facing == "north" {
		finalScore += 3
	}
	if cur.facing == "west" {
		finalScore += 2
	}
	if cur.facing == "south" {
		finalScore += 1
	}
	fmt.Printf("Part One: Final password after following the monkey's path is %d\n", finalScore)

}

type instruction struct {
	distance int
	turn     string
}

type coord struct {
	x, y int
}

type coordWithOrientation struct {
	c      coord
	facing string
}

func (c coordWithOrientation) turn(dir string) string {
	if c.facing == "east" {
		if dir == "R" {
			return "south"
		} else {
			return "north"
		}
	}
	if c.facing == "north" {
		if dir == "R" {
			return "east"
		} else {
			return "west"
		}
	}
	if c.facing == "west" {
		if dir == "R" {
			return "north"
		} else {
			return "south"
		}
	}
	if dir == "R" {
		return "west"
	} else {
		return "east"
	}
}

func (c1 coord) plus(c2 coord, lx, ly int) coord {
	return coord{wrap(c1.x+c2.x, lx), wrap(c1.y+c2.y, ly)}
}

type node struct {
	loc                      coord
	west, east, north, south *node
}

func (n node) step(facing string) *node {
	if facing == "north" {
		return n.north
	}
	if facing == "west" {
		return n.west
	}
	if facing == "east" {
		return n.east
	}
	return n.south
}

func (n *node) walk(m int, facing string) *node {
	for i := 0; i < m; i++ {
		n = n.step(facing)
	}
	return n
}

func wrap(a, b int) int {
	// only designed for taking one step at a time
	if a == -1 {
		return b - 1
	}
	if a == b {
		return 0
	}
	return a
}

// func wrapCube(cur coordWithOrientation, dir coord) coordWithOrientation {
// 	// ok this is only going to work for my puzzle input,
// 	// and was derived based on a cut out piece paper
// 	sideLength := 50
// 	var sideIndex int
// 	x := cur.c.x
// 	y := cur.c.y
// 	if x < 50 {
// 		if y < 150 {
// 			sideIndex = 5
// 		} else {
// 			sideIndex = 6
// 		}
// 	} else if x < 100 {
// 		if y < 50 {
// 			sideIndex = 2
// 		} else if y < 100 {
// 			sideIndex = 3
// 		} else {
// 			sideIndex = 4
// 		}
// 	} else {
// 		sideIndex = 1
// 	}

// 	if sideIndex == 1 {
// 		if x+dir.x == 150 {
// 			return coordWithOrientation{coord{99, y}, "west"}
// 		}
// 		if y+dir.y == -1 {
// 			return coordWithOrientation{coord{x, 199}, "north"}
// 		}
// 		if y+dir.y == 50 {
// 			return coordWithOrientation{coord{99, x - 50}, "west"} // my god
// 		}
// 	}
// }

func findNext(m map[coord]string, start, dir coord, lx, ly int) coord {
	remember := start
	for m[start.plus(dir, lx, ly)] == " " {
		start = start.plus(dir, lx, ly)
	}
	if m[start.plus(dir, lx, ly)] == "#" {
		return remember
	} else {
		// then the only remaining option is "."
		return start.plus(dir, lx, ly)
	}
}
