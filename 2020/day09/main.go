package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const preambleSize = 25

// This one felt pretty okay.

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

	var exceptionValue, exceptionValuePosition int

	for i := preambleSize; i < len(vals); i++ {
		if !isPairSum(vals[i], vals[i-25:i]) {
			exceptionValue = vals[i]
			exceptionValuePosition = i
			fmt.Printf(
				"Part One: Found %d at position %d which is not a sum of any pairs from the previous %d numbers.\n",
				exceptionValue, exceptionValuePosition, preambleSize,
			)

			break
		}
	}

	left := 0
	right := 1
	var cs int
	for right <= len(vals) && left < right {
		cs = sum(vals[left:right])
		if cs == exceptionValue {
			fmt.Printf("Part Two: Found key %d.\n", min(vals[left:right])+max(vals[left:right]))
			break
		} else if cs > exceptionValue {
			left += 1
		} else if cs < exceptionValue {
			right += 1
		}
	}

}

func isPairSum(v int, preamble []int) bool {
	for i, s1 := range preamble {
		for _, s2 := range preamble[i:] {
			if v == s1+s2 {
				return true
			}
		}
	}
	return false
}

func sum(a []int) int {
	var acc int
	for _, v := range a {
		acc += v
	}
	return acc
}

func min(a []int) int {
	r := a[0]
	for _, v := range a {
		if v < r {
			r = v
		}
	}
	return r
}

func max(a []int) int {
	r := a[0]
	for _, v := range a {
		if v > r {
			r = v
		}
	}
	return r
}
