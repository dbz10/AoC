package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var lines = strings.Split(string(contents), "\n\n")
	var totalAnyYes int

	// Part One

	type void struct{}
	var nothing void

	for _, line := range lines {
		combinedLine := strings.ReplaceAll(line, "\n", "")
		qSet := make(map[string]void)
		for _, char := range combinedLine {
			qSet[string(char)] = nothing
		}
		totalAnyYes += len(qSet)
	}

	fmt.Printf("Part One: Found %d total questions answered yes\n", totalAnyYes)

	// Part Two
	// Maybe there's a standard library function to do this... :thinking:
	// (and part one while we would have been at it)
	var totalAllYes int

	for _, line := range lines {
		numPeopleInParty := len(strings.Split(line, "\n"))
		combinedLine := strings.ReplaceAll(line, "\n", "")
		qSet := make(map[string]int)
		for _, char := range combinedLine {
			qSet[string(char)] += 1
		}

		for _, nYeses := range qSet {
			if nYeses == numPeopleInParty {
				totalAllYes++
			}
		}
	}

	fmt.Printf("Part Two: Found %d total questions for which all party members answered yes\n", totalAllYes)
}
