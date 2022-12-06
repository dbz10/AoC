package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var stop int = -1
	for i := 3; i < len(code); i++ {
		if checkSequenceNonRepeating(code, i, 4) {
			stop = i + 1
			break
		}
	}
	fmt.Printf("Part One: Start packet read starting at position %d\n", stop)

	stop = -1
	for i := 13; i < len(code); i++ {
		if checkSequenceNonRepeating(code, i, 14) {
			stop = i + 1
			break
		}
	}
	fmt.Printf("Part Two: Start message read starting at position %d\n", stop)

}

func checkSequenceNonRepeating(code []byte, start, n int) bool {
	var hashSet = make(map[byte]struct{})
	for i := 0; i < n; i++ {
		hashSet[code[start-i]] = struct{}{}
	}
	return len(hashSet) == n
}
