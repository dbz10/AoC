package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	state := read(string(contents))
	for cycle := 0; cycle < 6; cycle++ {
		state = state.next()
	}

	fmt.Printf("Part One: %d active cells after 6 cycles\n", sumActive(state))

}

// Since (at least in part one) we only evolve for six rounds,
// the farthest dimension we can reach is 8+6+6
type conway3 struct {
	layout map[[3]int]int
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	zMin   int
	zMax   int
}

func (c conway3) render(padSize int) {
	// For now just render at a constant size at z = 0 for checking
	for y := c.yMax + padSize - 1; y >= c.yMin-padSize; y-- {
		for x := c.xMin - padSize; x < c.xMax+padSize; x++ {
			position := [3]int{x, y, 0}
			if c.layout[position] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (c conway3) next() conway3 {
	newLayout := map[[3]int]int{}

	for x := c.xMin - 1; x <= c.xMax+1; x++ {
		for y := c.yMin - 1; y <= c.yMax+1; y++ {
			for z := c.zMin - 1; z <= c.zMax+1; z++ {
				neighbors := sumNeighbors(x, y, z, c.layout)
				if c.layout[[3]int{x, y, z}] == 1 && (neighbors == 2 || neighbors == 3) {
					newLayout[[3]int{x, y, z}] = 1

				} else if neighbors == 3 {
					newLayout[[3]int{x, y, z}] = 1
				} else {
					newLayout[[3]int{x, y, z}] = 0
				}
			}
		}
	}
	return conway3{layout: newLayout, xMin: c.xMin - 1, xMax: c.xMax + 1, yMin: c.yMin - 1, yMax: c.yMax + 1, zMin: c.zMin - 1, zMax: c.zMax + 1}
}

func sumActive(c conway3) int {
	acc := 0
	for x := c.xMin - 1; x <= c.xMax+1; x++ {
		for y := c.yMin - 1; y <= c.yMax+1; y++ {
			for z := c.zMin - 1; z <= c.zMax+1; z++ {
				acc += c.layout[[3]int{x, y, z}]
			}
		}
	}
	return acc
}

func sumNeighbors(x int, y int, z int, l map[[3]int]int) int {
	acc := -l[[3]int{x, y, z}]
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				acc += l[[3]int{x + dx, y + dy, z + dz}]
			}
		}
	}

	return acc
}

func read(contents string) conway3 {
	// z starts at zero
	layout := map[[3]int]int{}
	lines := strings.Split(contents, "\n")
	yMax := len(lines)
	xMax := len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			position := [3]int{x, y, 0}
			if string(char) == "." {
				layout[position] = 0
			} else {
				layout[position] = 1
			}
		}
	}
	return conway3{layout: layout,
		xMin: 0,
		yMin: 0,
		zMin: 0,
		xMax: xMax,
		yMax: yMax,
		zMax: 0}
}
