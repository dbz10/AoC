package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var charValue = map[string]int{"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for key := range charValue {
		charValue[strings.ToUpper(key)] = charValue[key] + 26
	}

	var totalPriority, badgePriority int
	var register = make(map[string]bool)
	var tripletIterator int

outer:
	for _, line := range strings.Split(string(contents), "\n") {
		// Part Two
		if tripletIterator == 0 {
			// create a new register to track common items
			register = map[string]bool{}
			for _, ch := range line {
				register[string(ch)] = true
			}
		} else {
			// if we are the second or third elf in the triplet,
			// eliminate items that we are not carrying from the register
			for key := range register {
				if !strings.Contains(line, key) {
					delete(register, key)
				}
			}
		}
		tripletIterator++

		if tripletIterator == 3 {
			// The triplet is complete, so evaluate what is left.
			// And then get ready to start the next triplet.
			if len(register) > 1 {
				fmt.Fprintf(os.Stderr, "Found more than one element left in the register. That was unexpected: %v", register)
			}
			for key := range register {
				badgePriority += charValue[key]
			}
			tripletIterator = 0
		}

		// Part One
		l := len(line)
		left := line[:l/2]
		right := line[l/2:]
		for _, ch := range left {
			char := string(ch)
			if strings.Contains(right, char) {
				// problem statement says there is only one such item type per backpack
				// however the item type can be present more than one time. So we do
				// have to break the loop.
				totalPriority += charValue[char]
				continue outer
			}
		}

	}

	fmt.Printf("Part One: Sum of priorities of mispacked items: %d\n", totalPriority)
	fmt.Printf("Part One: Sum of security badge priorities: %d\n", badgePriority)

}
