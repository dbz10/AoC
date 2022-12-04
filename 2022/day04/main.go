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

	var completeContainments, partialOverlaps int

	for _, pair := range strings.Split(string(contents), "\n") {
		leftRange := strings.Split(pair, ",")[0]
		rightRange := strings.Split(pair, ",")[1]

		// bad practice but for the sake of LoC, not checking the error
		leftStart, _ := strconv.Atoi(strings.Split(leftRange, "-")[0])
		leftEnd, _ := strconv.Atoi(strings.Split(leftRange, "-")[1])
		rightStart, _ := strconv.Atoi(strings.Split(rightRange, "-")[0])
		rightEnd, _ := strconv.Atoi(strings.Split(rightRange, "-")[1])

		if (leftStart <= rightStart && leftEnd >= rightEnd) || (rightStart <= leftStart && rightEnd >= leftEnd) {
			completeContainments++
		}

		if (leftStart <= rightStart && leftEnd >= rightStart) || (rightStart <= leftStart && rightEnd >= leftStart) {
			partialOverlaps++
		}

	}

	fmt.Printf("Part One: Found %d instances where one of the elves assignment completely contained the others\n", completeContainments)
	fmt.Printf("Part Two: Found %d instances where the pair's assignments partially overlapped\n", partialOverlaps)

}
