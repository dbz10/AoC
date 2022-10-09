package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	xStep    = 3
	yStep    = 1
	treeChar = "#"
)

func main() {
	content, err := readFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	treesEncounteredPartOne := partOne(content)

	fmt.Printf("Part One: Found %d trees on the way through the forest\n", treesEncounteredPartOne)

	var treesEncounteredPartTwo int = 1

	for _, slope := range []scooterSlope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2}} {
		treesEncounteredPartTwo *= partTwo(content, slope.xStep, slope.yStep)
	}

	fmt.Printf("Part Two: Found %d trees on the way through the forest\n", treesEncounteredPartTwo)

}

type scooterSlope struct {
	xStep int
	yStep int
}

func partOne(content [][]string) int {
	nRows := len(content)
	nCols := len(content[0])

	var xPos, treesEncountered int

	for yPos := 0; yPos < nRows; yPos += yStep {

		if content[yPos][xPos] == treeChar {
			treesEncountered++
		}
		xPos = wrapIndex(xPos+xStep, nCols)
	}
	return treesEncountered
}

func partTwo(content [][]string, xStep int, yStep int) int {
	nRows := len(content)
	nCols := len(content[0])

	var xPos, treesEncountered int

	for yPos := 0; yPos < nRows; yPos += yStep {

		if content[yPos][xPos] == treeChar {
			treesEncountered++
		}
		xPos = wrapIndex(xPos+xStep, nCols)
	}
	return treesEncountered
}

func readFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := bufio.NewScanner(file)

	var b [][]string
	for lines.Scan() {
		row := []string{}
		for _, char := range lines.Text() {
			row = append(row, string(char))
		}
		b = append(b, row)
	}
	return b, nil
}

func wrapIndex(idx int, dimSize int) int {
	return int(math.Mod(float64(idx), float64(dimSize)))
	// more simply for integers could also be idx - (idx/dimSize) * dimSize
}
