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
				occupiedNeighbors := (scanX(layout, x, y, 1) +
					scanX(layout, x, y, -1) +
					scanY(layout, x, y, 1) +
					scanY(layout, x, y, -1) +
					scanUR(layout, x, y, 1) +
					scanUR(layout, x, y, -1) +
					scanDR(layout, x, y, 1) +
					scanDR(layout, x, y, -1))
				if occupiedNeighbors >= 5 {
					tmpLayout[x][y] = patch{occupiable: true, occupied: 0, render: "L"}
				} else if occupiedNeighbors == 0 {
					tmpLayout[x][y] = patch{occupiable: true, occupied: 1, render: "#"}
				}
			}

		}

		// printLayout(tmpLayout)
		// fmt.Println(scanX(tmpLayout, 0, 0, -1))
		// fmt.Println("")

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

// Not my best work but it gets the job done

func scanX(layout seats, curX int, curY int, sign int) int {
	xDim := len(layout)

	for x := curX + sign; x < xDim && x >= 0; x += sign {
		if layout[x][curY].occupiable {
			return layout[x][curY].occupied
		}
	}
	return 0

}

func scanY(layout seats, curX int, curY int, sign int) int {
	yDim := len(layout[0])

	for y := curY + sign; y < yDim && y >= 0; y += sign {
		if layout[curX][y].occupiable {
			return layout[curX][y].occupied
		}
	}
	return 0

}

func scanUR(layout seats, x int, y int, sign int) int {
	xDim := len(layout)
	yDim := len(layout[0])

	x += sign
	y += sign

	for x < xDim && x >= 0 && y < yDim && y >= 0 {
		if layout[x][y].occupiable {
			return layout[x][y].occupied
		}
		x += sign
		y += sign
	}
	return 0

}

func scanDR(layout seats, x int, y int, sign int) int {
	xDim := len(layout)
	yDim := len(layout[0])

	x += sign
	y += -sign

	for x < xDim && x >= 0 && y < yDim && y >= 0 {
		if layout[x][y].occupiable {
			return layout[x][y].occupied
		}
		x += sign
		y += -sign
	}
	return 0

}
