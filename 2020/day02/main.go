package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// after completing this, I learned about the function strings.Count!
// and about fmt.Sscanf which can parse a string into variables by format!
// and we could have used a struct to handle the data,
// which might have been a nicer approach.

func main() {
	file, err := os.Open("inputs/day02_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := bufio.NewScanner(file)

	validPasswordsPartOne := 0
	validPasswordsPartTwo := 0
	for lines.Scan() {
		splitText := strings.Split(lines.Text(), " ")
		minOccurrences, _ := strconv.Atoi(strings.Split(splitText[0], "-")[0])
		maxOccurrences, _ := strconv.Atoi(strings.Split(splitText[0], "-")[1])
		key := string(splitText[1][0])
		phrase := splitText[2]
		if partOne(phrase, key, minOccurrences, maxOccurrences) {
			validPasswordsPartOne++
		}
		// it's a bit unclear reusing the same variable "minOccurrences" and "maxOccurrences"
		// despite the different meaning in part two but anyways
		if partTwo(phrase, key, minOccurrences, maxOccurrences) {
			validPasswordsPartTwo++
		}
	}

	fmt.Printf("Found %d valid passwords for part one \n", validPasswordsPartOne)
	fmt.Printf("Found %d valid passwords for part two \n", validPasswordsPartTwo)

}

func partOne(phrase string, key string, minOccurrences int, maxOccurrences int) bool {
	counter := 0
	for _, char := range phrase {
		if string(char) == key {
			counter++
		}
	}
	return (counter >= minOccurrences) && (counter <= maxOccurrences)
}

func partTwo(phrase string, key string, indexOne int, indexTwo int) bool {
	// kinda XOR
	left := string(phrase[indexOne-1])
	right := string(phrase[indexTwo-1])
	return (left == key) != (right == key)
}
