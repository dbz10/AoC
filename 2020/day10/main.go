package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var vals []int
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		vals = append(vals, v)
	}

	// Sort them
	sort.Ints(vals)

	// Accumulate diff in volts
	var vDiffs = make(map[int]int)
	// The first adapter
	vDiffs[vals[0]] += 1
	for i := 1; i < len(vals); i++ {
		vDiffs[vals[i]-vals[i-1]] += 1
	}
	// Can always have that one
	vDiffs[3] += 1

	fmt.Printf("Part One: Found the following distribution of voltage differences %v, with the following key %d.\n", vDiffs, vDiffs[1]*vDiffs[3])

	// Part Two
	// Classic dynamic programming style but slight twist since tree can branch one, two, or three times.
	// Pretty happy with this solution!

	// Actually we don't need a memo here but it was useful for debugging.
	// The whole problem could be solved with three numbers l, c, r
	// And something like
	// for i := 0; i < len(vals)-3; i++ {
	// 	if vals[i]-vals[i-3] <= 3 {
	// 		l, c, r = c, r, l + c + r
	// 	} else if vals[i] - vals[i-2] <= 3 {
	// 		l, c, r = c, r, c + r
	// 	} else {
	// 		l, c, r = c, r, r
	// 	}
	// }
	// So if space was at a premium... this would have been another way I guess.

	var memo = make(map[int]int)

	// Initial conditions
	memo[0] = 1
	memo[1] = memo[0]
	if vals[1] <= 3 {
		memo[1] += 1
	}

	memo[2] = memo[1]
	if vals[2]-vals[0] <= 3 {
		memo[2] += memo[0]
	}
	if vals[2] <= 3 {
		memo[2] += 1
	}

	for i := 3; i < len(vals); i++ {
		memo[i] = memo[i-1]
		if vals[i]-vals[i-2] <= 3 {
			memo[i] += memo[i-2]
		}
		if vals[i]-vals[i-3] <= 3 {
			memo[i] += memo[i-3]
		}
	}

	fmt.Printf("Part Two: Found %d ways of arranging the adapters.\n", memo[len(vals)-1])

}
