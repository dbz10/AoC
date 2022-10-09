package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var lines = strings.Split(string(contents), "\n\n")
	var validLinesPartOne int
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, line := range lines {
		if allKeysPresent(line, requiredKeys) {
			validLinesPartOne++
		}
	}

	fmt.Printf("Part One: Found %d identification documents containing all required keys\n", validLinesPartOne)

	// Part Two
	validityConditions := map[string]func(input string) bool{
		"byr": validByr,
		"iyr": validIyr,
		"eyr": validEyr,
		"hgt": validHgt,
		"hcl": validHcl,
		"ecl": validEcl,
		"pid": validPid,
	}

	var validLinesPartTwo int
	for _, line := range lines {
		if allKeysPresentAndValid(line, validityConditions) {
			validLinesPartTwo++
		}
	}

	fmt.Printf("Part Two: Found %d identification documents containing all required keys with valid entries\n", validLinesPartTwo)

}

func allKeysPresent(line string, expectedKeys []string) bool {
	var kvMap = make(map[string]string)
	var key, value string

	reformattedLine := strings.Join(strings.Split(line, "\n"), " ")
	for _, keyVal := range strings.Split(reformattedLine, " ") {
		_, err := fmt.Sscanf(strings.Join(strings.Split(keyVal, ":"), " "), "%s %s", &key, &value)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}
		kvMap[key] = value
	}

	for _, key := range expectedKeys {
		_, exists := kvMap[key]
		if !exists {
			return false
		}
	}
	return true
}

func allKeysPresentAndValid(line string, validityConditions map[string]func(string) bool) bool {
	var kvMap = make(map[string]string)
	var key, value string

	reformattedLine := strings.Join(strings.Split(line, "\n"), " ")
	for _, keyVal := range strings.Split(reformattedLine, " ") {
		_, err := fmt.Sscanf(strings.Join(strings.Split(keyVal, ":"), " "), "%s %s", &key, &value)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}
		kvMap[key] = value
	}

	for key, condition := range validityConditions {
		value, exists := kvMap[key]
		if !exists || !condition(value) {
			return false
		}
	}
	return true
}

func validByr(byr string) bool {
	year, _ := strconv.Atoi(byr)
	return year >= 1920 && year <= 2002
}

func validIyr(iyr string) bool {
	year, _ := strconv.Atoi(iyr)
	return year >= 2010 && year <= 2020
}

func validEyr(eyr string) bool {
	year, _ := strconv.Atoi(eyr)
	return year >= 2020 && year <= 2030
}

func validHgt(height string) bool {
	// here, I found out or realized later that I could have used
	// Sscanf(height, "%s%d"), which would have been able to handle
	// suffixes of different length
	if strings.HasSuffix(height, "cm") {
		heightCm, _ := strconv.Atoi(height[:len(height)-2])
		return heightCm >= 150 && heightCm <= 193
	} else if strings.HasSuffix(height, "in") {
		heightIn, _ := strconv.Atoi(height[:len(height)-2])
		return heightIn >= 59 && heightIn <= 76
	} else {
		return false
	}
}

func validHcl(hcl string) bool {
	matched, _ := regexp.Match(`^#[0-9a-f]{6}$`, []byte(hcl))
	return matched
}

func validEcl(ecl string) bool {
	matched, _ := regexp.Match(`^amb|blu|brn|gry|grn|hzl|oth$`, []byte(ecl))
	return matched
}

func validPid(pid string) bool {
	matched, _ := regexp.Match(`^\d{9}$`, []byte(pid))
	return matched
}
