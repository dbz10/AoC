package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// seems clumsy but ok
	var layout = make(map[int]map[int]patch)
	var tmpLayout = make(map[int]map[int]patch)

	var row int
	for scanner.Scan() {
		layout[row] = make(map[int]patch)
		tmpLayout[row] = make(map[int]patch)
		line := strings.Split(scanner.Text(), "")
		for col, char := range line {
			p, err := parse(char)
			if err != nil {
				log.Fatal(err)
			}
			layout[row][col] = p
		}
		row++
	}

	copyLayout(layout, tmpLayout)

	xDim := len(layout)
	yDim := len(layout[0])

	for {
		for x := 0; x < xDim; x++ {
			for y := 0; y < yDim; y++ {
				thisTile := layout[x][y]
				if !thisTile.occupiable {
					continue
				}
				occupiedNeighbors := (layout[x][y+1].occupied +
					layout[x][y-1].occupied +
					layout[x+1][y+1].occupied +
					layout[x+1][y-1].occupied +
					layout[x-1][y+1].occupied +
					layout[x-1][y-1].occupied +
					layout[x-1][y].occupied +
					layout[x+1][y].occupied)
				if occupiedNeighbors >= 4 {
					tmpLayout[x][y] = patch{occupiable: true, occupied: 0, render: "L"}
				} else if occupiedNeighbors == 0 {
					tmpLayout[x][y] = patch{occupiable: true, occupied: 1, render: "#"}
				}
			}

		}

		if checkLayoutsEqual(layout, tmpLayout) {
			copyLayout(tmpLayout, layout)
			break
		}

		copyLayout(tmpLayout, layout)
	}

	fmt.Println(countOccupiedSeats(layout))
}

type patch struct {
	occupiable bool
	occupied   int
	render     string
}

type seats map[int]map[int]patch

func parse(s string) (patch, error) {
	// Expect all seats to be empty in the beginning, but
	// just in case, check.
	switch s {
	case "L":
		return patch{occupiable: true, occupied: 0, render: s}, nil
	case ".":
		return patch{occupiable: false, occupied: 0, render: s}, nil
	case "#":
		return patch{occupiable: true, occupied: 1, render: s}, nil
	default:
		return patch{false, 0, ""}, fmt.Errorf("could not parse %s to a patch", s)
	}

}

func checkLayoutsEqual(l1 seats, l2 seats) bool {
	xDim := len(l1)
	yDim := len(l1[0])

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {

			if l1[x][y] != l2[x][y] {
				return false
			}
		}
	}
	return true
}

func printLayout(layout seats) {
	xDim := len(layout)
	yDim := len(layout[0])

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			fmt.Print(layout[x][y].render)
		}
		fmt.Print("\n")
	}
}

func countOccupiedSeats(layout seats) int {
	// here, location doesn't matter
	var acc int
	for _, line := range layout {
		for _, c := range line {
			acc += c.occupied
		}
	}
	return acc
}

func copyLayout(src seats, dst seats) {
	xDim := len(src)
	yDim := len(src[0])

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			dst[x][y] = src[x][y]
		}
	}
}
