package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) d(other Point) Point {
	return Point{other.x - p.x, other.y - p.y}
}

type Sand Point

func (s *Sand) move(dx, dy int) {
	s.x += dx
	s.y += dy
}

func (s Sand) findMove(grid map[Point]int) (Point, error) {
	// Down is positive...
	if grid[Point{s.x, s.y + 1}] == open {
		return Point{0, 1}, nil
	} else if grid[Point{s.x - 1, s.y + 1}] == open {
		return Point{-1, 1}, nil
	} else if grid[Point{s.x + 1, s.y + 1}] == open {
		return Point{1, 1}, nil
	}
	return Point{}, errors.New("nowhere to go")
}

const (
	open = iota
	sand
	rock
	source
)

var rendering = map[int]string{
	open:   ".",
	sand:   "o",
	rock:   "#",
	source: "+",
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	grid, xMin, xMax, yMin, yMax := parseGrid(string(contents))
	fmt.Printf("Logging all boundaries so the compiler won't complain: xMin: %d xMax: %d yMin: %d yMax: %d\n", xMin, xMax, yMin, yMax)
	var totalGrains int
	var grain = Sand{500, 0}
	for grain.y < yMax {
		grid[Point(grain)] = sand
		move, err := grain.findMove(grid)
		if err != nil {
			grid[Point(grain)] = sand
			grain = Sand{500, 0}
			totalGrains++
		}
		grid[Point(grain)] = open
		grain.move(move.x, move.y)
	}

	fmt.Printf("Part One: %d grains of sand piled up before they start falling endlessly into the void\n", totalGrains)

	// Part Two, can just modify the loop I think.

	// reinitialize the grid.
	grid, _, _, _, yMax = parseGrid(string(contents))
	totalGrains = 0
	grain = Sand{500, 0}
	for {
		grid[Point(grain)] = sand
		if grain.y == yMax+1 {
			// Then there is an invisible infinite horizontal floor at the square below
			grid[Point(grain)] = sand
			grain = Sand{500, 0}
			totalGrains++
		}
		move, err := grain.findMove(grid)
		if err != nil {
			// Now we have a little bit more logic
			if grain.y == 0 {
				totalGrains++
				break
			}
			grid[Point(grain)] = sand
			grain = Sand{500, 0}
			totalGrains++
		}
		grid[Point(grain)] = open
		grain.move(move.x, move.y)
	}

	fmt.Printf("Part Two: %d grains of sand piled up before they blocked the source\n", totalGrains)

}

func parseGrid(block string) (map[Point]int, int, int, int, int) {
	grid := map[Point]int{}
	lines := strings.Split(block, "\n")
	for _, line := range lines {
		positions := strings.Split(line, " -> ")
		for i := 0; i < len(positions)-1; i++ {
			l := tupleToPoint(positions[i])
			r := tupleToPoint(positions[i+1])
			delta := l.d(r)

			// Rock goes in straight lines so only one of dx or dy is nonzero
			if delta.x != 0 {
				direction := sign(delta.x)
				for j := 0; j <= delta.x*direction; j++ {
					grid[Point{l.x + j*direction, l.y}] = rock
				}
			}
			if delta.y != 0 {
				direction := sign(delta.y)
				for j := 0; j <= delta.y*direction; j++ {
					grid[Point{l.x, l.y + j*direction}] = rock
				}
			}

		}
	}

	grid[Point{500, 0}] = source
	keys := make([]Point, 0, len(grid))
	for p := range grid {
		keys = append(keys, p)
	}
	xMin := keys[0].x
	xMax := keys[0].x
	yMin := keys[0].y
	yMax := keys[0].y
	for p := range grid {
		if p.x <= xMin {
			xMin = p.x
		}
		if p.x >= xMax {
			xMax = p.x
		}
		if p.y <= yMin {
			yMin = p.y
		}
		if p.y >= yMax {
			yMax = p.y
		}
	}

	// though only yMax is needed, extras are returned just for visualization for fun
	return grid, xMin, xMax, yMin, yMax
}

func tupleToPoint(s string) Point {
	sp := strings.Split(s, ",")
	x, err := strconv.Atoi(sp[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(sp[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}

func sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}
	return 0
}

func renderMap(grid map[Point]int, xMin, xMax, yMin, yMax int) string {
	output := ""
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			output += rendering[grid[Point{x, y}]]
		}
		output += "\n"
	}
	return output
}
