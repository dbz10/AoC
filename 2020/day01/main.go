package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const targetSum = 2020

func main() {
	arr := readFileToInt("inputs/day01_p1.txt")
	v1, v2, err := findPair(arr, targetSum)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found values %d and %d, which multiply to %d\n", v1, v2, v1*v2)
	}

	triplets, err := findThree(arr, targetSum)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found values %d, which multiply to %d\n", triplets, triplets[0]*triplets[1]*triplets[2])
	}
}

func findPair(arr []int, targetSum int) (int, int, error) {
	for i, val1 := range arr {
		for _, val2 := range arr[i:] {
			if val1+val2 == targetSum {
				return val1, val2, nil
			}
		}
	}
	return 0, 0, errors.New("Could not find a pair satisfying the condition")
}

func findThree(arr []int, targetSum int) ([3]int, error) {
	// for fun, consider if the function would be extended even further to say findN
	// findN could be implemented recursively, something along the lines of

	// func findN(arr, n targetSum) int[], error {
	// 	for i, v := range arr {
	// 		res, err := findN(arr[i:], n-1, targetSum - v)
	// 		if err == nil {
	// 			return append(res, v)
	// 		}
	// 	}
	// }

	// but would need to work out the exact details

	// here we can reuse the previous function
	for i, v1 := range arr {
		v2, v3, err := findPair(arr[i:], targetSum-v1)
		if err == nil {
			return [3]int{v1, v2, v3}, nil
		}
	}
	return [3]int{}, errors.New("Could not find values to satisfy the condition :(")
}

func readFileToInt(filePath string) []int {
	fileContents, _ := os.ReadFile(filePath)
	lines := strings.Split(string(fileContents), "\n")
	b := make([]int, len(lines))
	for i, value := range lines {
		b[i], _ = strconv.Atoi(value)
	}
	return b
}
