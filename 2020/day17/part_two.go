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
	state := read4(string(contents))
	for cycle := 0; cycle < 6; cycle++ {
		state = state.next()
	}

	fmt.Printf("Part Two: %d active cells after 6 cycles\n", sumActive4(state))

}

// Since (at least in part one) we only evolve for six rounds,
// the farthest dimension we can reach is 8+6+6
type conway4 struct {
	layout map[[4]int]int
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	zMin   int
	zMax   int
	wMin   int
	wMax   int
}

func (c conway4) render(padSize int) {
	// For now just render at a constant size at z = 0 for checking
	for y := c.yMax + padSize - 1; y >= c.yMin-padSize; y-- {
		for x := c.xMin - padSize; x < c.xMax+padSize; x++ {
			position := [4]int{x, y, 0, 0}
			if c.layout[position] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (c conway4) next() conway4 {
	newLayout := map[[4]int]int{}

	for x := c.xMin - 1; x <= c.xMax+1; x++ {
		for y := c.yMin - 1; y <= c.yMax+1; y++ {
			for z := c.zMin - 1; z <= c.zMax+1; z++ {
				for w := c.wMin - 1; w <= c.wMax+1; w++ {
					neighbors := sumNeighbors4(x, y, z, w, c.layout)
					if c.layout[[4]int{x, y, z, w}] == 1 && (neighbors == 2 || neighbors == 3) {
						newLayout[[4]int{x, y, z, w}] = 1
					} else if neighbors == 3 {
						newLayout[[4]int{x, y, z, w}] = 1
					} else {
						newLayout[[4]int{x, y, z, w}] = 0
					}
				}
			}
		}
	}
	return conway4{layout: newLayout,
		xMin: c.xMin - 1,
		xMax: c.xMax + 1,
		yMin: c.yMin - 1,
		yMax: c.yMax + 1,
		zMin: c.zMin - 1,
		zMax: c.zMax + 1,
		wMin: c.wMin - 1,
		wMax: c.wMax + 1}
}

func sumActive4(c conway4) int {
	acc := 0
	for x := c.xMin - 1; x <= c.xMax+1; x++ {
		for y := c.yMin - 1; y <= c.yMax+1; y++ {
			for z := c.zMin - 1; z <= c.zMax+1; z++ {
				for w := c.wMin - 1; w <= c.wMax+1; w++ {
					acc += c.layout[[4]int{x, y, z, w}]
				}
			}
		}
	}
	return acc
}

func sumNeighbors4(x int, y int, z int, w int, l map[[4]int]int) int {
	acc := -l[[4]int{x, y, z, w}]
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				for _, dw := range []int{-1, 0, 1} {
					acc += l[[4]int{x + dx, y + dy, z + dz, w + dw}]

				}
			}
		}
	}

	return acc
}

func read4(contents string) conway4 {
	// z and w start at zero
	layout := map[[4]int]int{}
	lines := strings.Split(contents, "\n")
	yMax := len(lines)
	xMax := len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			position := [4]int{x, y, 0, 0}
			if string(char) == "." {
				layout[position] = 0
			} else {
				layout[position] = 1
			}
		}
	}
	return conway4{layout: layout,
		xMin: 0,
		yMin: 0,
		zMin: 0,
		wMin: 0,
		xMax: xMax,
		yMax: yMax,
		zMax: 0,
		wMax: 0}
}
