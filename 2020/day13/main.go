package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	eeta, err := strconv.Atoi(strings.Split(string(contents), "\n")[0])
	if err != nil {
		log.Fatal(err)
	}

	busLines := []int{}
	for _, ch := range strings.Split(strings.Split(string(contents), "\n")[1], ",") {
		if ch != "x" {
			lineNo, err := strconv.Atoi(ch)
			if err != nil {
				log.Fatal(err)
			}
			busLines = append(busLines, lineNo)
		}
	}

	shortestWait := (busLines[0]*(eeta/busLines[0]+1) - eeta)
	earliestTimestampId := busLines[0] * shortestWait
	for _, v := range busLines {
		wait := v*(eeta/v+1) - eeta
		if wait < shortestWait {
			shortestWait = wait
			earliestTimestampId = wait * v
		}
	}
	fmt.Printf("Part One: %d.\n", earliestTimestampId)

	// Part two is somewhat nontrivial....
	// Brute force?
	sort.Sort(sort.Reverse(sort.IntSlice(busLines)))
	var trial int
	for n := 1; n < 1e10; n++ {
		trial = n*busLines[0] + busLines[0] - busLines[len(busLines)-1]
		if checkSequenceCongruence(trial, busLines[1:]) {
			fmt.Printf("Found an answer! %d.\n", trial)
			break
		}
	}

	log.Fatal("Part 2 by brute force... did not get there.")

}

func checkSequenceCongruence(n int, s []int) bool {
	for _, v := range s {
		if int(math.Mod(float64(n), float64(v))) != v-s[len(s)-1] {
			return false
		}
	}
	return true
}
