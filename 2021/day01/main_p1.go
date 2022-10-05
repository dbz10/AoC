package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContents := readFileToInt("inputs/day01_p1.txt")
	fmt.Println("Part 1")
	fmt.Println(countIncreases(fileContents, 1))
	fmt.Println("Part 2")
	fmt.Println(countIncreases(fileContents, 3))

}

func readFileToInt(filePath string) []int {
	fileContents, _ := os.ReadFile(filePath)
	b := make([]int, len(fileContents))

	lines := strings.Split(string(fileContents), "\n")
	for i, value := range lines {
		b[i], _ = strconv.Atoi(value)
	}
	return b

}

func countIncreases(arr []int, windowSize int) int {

	counter := 0
	for i := 0; i < len(arr)-windowSize-1; i++ {
		if arraySum(arr[i+1:i+1+windowSize]) > arraySum(arr[i:i+windowSize]) {
			counter += 1
		}
	}

	return counter

}

func arraySum(arr []int) int {
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}
