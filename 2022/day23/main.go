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

	elves := Elves{}
	var xMin, xMax, yMin, yMax int

	for y, row := range strings.Split(string(contents), "\n") {
		for x, ch := range strings.Split(row, "") {
			if ch == "#" {
				elves[coord{x, y}] = true
			}
			if x < xMin {
				xMin = x
			}
			if x > xMax {
				xMax = x
			}
			if y < yMin {
				yMin = y
			}
			if y > yMax {
				yMax = y
			}
		}
	}

	directions := []coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	part := os.Args[2]
	maxRounds := 10
	if part == "2" {
		maxRounds = 10000000
	}

	for round := 0; round < maxRounds; round++ {
		contentionLog := refreshContention(xMin, xMax, yMin, yMax)
		try := map[coord]coord{}

		// stage moves
		for elf := range elves {
			if !elves.hasNeighbor(elf) {
				continue
			}
			for d := 0; d < 4; d++ {
				dirIndex := (round + d) % 4
				if elves.directionOk(elf, directions[dirIndex]) {
					tryGo := elf.plus(directions[dirIndex])
					try[elf] = tryGo
					contentionLog[tryGo].stage()
					break
				}
			}
		}

		if len(try) == 0 {
			fmt.Printf("Part Two: The elves will fully spread out in %d rounds\n", round+1)
			break
		}
		// commit moves
		for elf, tryMove := range try {
			if !contentionLog[tryMove].contended {
				delete(elves, elf)
				elves[tryMove] = true
				if tryMove.x < xMin {
					xMin = tryMove.x
				}
				if tryMove.x > xMax {
					xMax = tryMove.x
				}
				if tryMove.y < yMin {
					yMin = tryMove.y
				}
				if tryMove.y > yMax {
					yMax = tryMove.y
				}
			}
		}
	}

	// some elves could have retreated, so re-calculate xmin, xmax, etc
	xMin, xMax, yMin, yMax = (xMin+xMax)/2, (xMin+xMax)/2, (yMin+yMax)/2, (yMin+yMax)/2
	for key := range elves {
		if key.x < xMin {
			xMin = key.x
		}
		if key.x > xMax {
			xMax = key.x
		}
		if key.y < yMin {
			yMin = key.y
		}
		if key.y > yMax {
			yMax = key.y
		}
	}

	fmt.Printf("Part One: Smallest rectangle enclosing all the elves has %d empty ground tiles\n", (xMax-xMin+1)*(yMax-yMin+1)-len(elves))

}

type Elves map[coord]bool

type coord struct {
	x, y int
}

func (c coord) plus(o coord) coord {
	return coord{c.x + o.x, c.y + o.y}
}

type coordLock struct {
	staged    bool
	contended bool
}

func (c *coordLock) stage() {
	if c.staged {
		c.contended = true
	}
	c.staged = true
}

func (e Elves) hasNeighbor(c coord) bool {
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			if e[coord{c.x + dx, c.y + dy}] {
				return true
			}
		}
	}
	return false
}

func (e Elves) directionOk(c coord, d coord) bool {
	north := coord{0, -1}
	south := coord{0, +1}
	west := coord{-1, 0}
	east := coord{1, 0}

	northwest := north.plus(west)
	northeast := north.plus(east)
	southwest := south.plus(west)
	southeast := south.plus(east)

	if d == north {
		return !(e[c.plus(north)] || e[c.plus(northwest)] || e[c.plus(northeast)])
	}
	if d == south {
		return !(e[c.plus(south)] || e[c.plus(southwest)] || e[c.plus(southeast)])
	}
	if d == east {
		return !(e[c.plus(east)] || e[c.plus(northeast)] || e[c.plus(southeast)])
	}
	if d == west {
		return !(e[c.plus(west)] || e[c.plus(northwest)] || e[c.plus(southwest)])
	}

	return true
}

func refreshContention(xMin, xMax, yMin, yMax int) map[coord]*coordLock {
	m := map[coord]*coordLock{}
	for x := xMin - 1; x <= xMax+1; x++ {
		for y := yMin - 1; y <= yMax+1; y++ {
			m[coord{x, y}] = &coordLock{}
		}
	}
	return m
}
