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

	lines := strings.Split(string(contents), "\n")
	arr := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, ch := range line {
			i, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, i)
		}
		arr = append(arr, row)
	}

	var totalVisible int

	for x := range arr {
		for y := range arr[x] {
			if isVisible(arr, x, y) {
				totalVisible++
			}
		}
	}

	fmt.Printf("Part One: %d trees are visible\n", totalVisible)

	// Part Two
	var maxScenicScore, ss int

	for x := range arr {
		for y := range arr[x] {
			ss = scenicScore(arr, x, y)
			if ss > maxScenicScore {
				maxScenicScore = ss
			}
		}
	}

	fmt.Printf("Part Two: Maximum scenic score found was %d\n", maxScenicScore)

}

func isVisible(arr [][]int, x, y int) bool {
	// Not the most efficient since we continue scanning even
	// if it was visible in one direction, but okay.

	// scan up
	var visU, visD, visL, visR = true, true, true, true
	for xt := 0; xt < x; xt++ {
		if arr[xt][y] >= arr[x][y] {
			visU = false
		}
	}
	// scan down
	for xt := x + 1; xt < len(arr); xt++ {
		if arr[xt][y] >= arr[x][y] {
			visD = false
		}
	}
	// scan left
	for yt := 0; yt < y; yt++ {
		if arr[x][yt] >= arr[x][y] {
			visL = false
		}
	}
	// scan right
	for yt := y + 1; yt < len(arr[0]); yt++ {
		if arr[x][yt] >= arr[x][y] {
			visR = false
		}
	}
	return visU || visD || visL || visR
}

func scenicScore(arr [][]int, x, y int) int {
	// Not the most efficient since we continue scanning even
	// if it was visible in one direction, but okay.

	// scan up
	var canSeeU, canSeeD, canSeeL, canSeeR int
	for xt := x - 1; xt >= 0; xt-- {
		canSeeU++
		if arr[xt][y] >= arr[x][y] {
			break
		}
	}
	// scan down
	for xt := x + 1; xt < len(arr); xt++ {
		canSeeD++
		if arr[xt][y] >= arr[x][y] {
			break
		}
	}
	// scan left
	for yt := y - 1; yt >= 0; yt-- {
		canSeeL++
		if arr[x][yt] >= arr[x][y] {
			break
		}
	}
	// scan right
	for yt := y + 1; yt < len(arr[0]); yt++ {
		canSeeR++
		if arr[x][yt] >= arr[x][y] {
			break
		}
	}
	return canSeeU * canSeeD * canSeeL * canSeeR
}
