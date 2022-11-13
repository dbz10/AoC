package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var deckOne, deckTwo []int

	for _, v := range strings.Split(strings.Split(string(contents), "\n\n")[0], "\n")[1:] {
		value, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		deckOne = append(deckOne, value)
	}

	for _, v := range strings.Split(strings.Split(string(contents), "\n\n")[1], "\n")[1:] {
		value, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		deckTwo = append(deckTwo, value)
	}

	for {
		if len(deckOne) == 0 || len(deckTwo) == 0 {
			break
		}

		drawOne := deckOne[0]
		drawTwo := deckTwo[0]

		// They don't say anything about tie breaks so I assume a tie would not be possible... in part one
		if drawOne > drawTwo {
			deckOne = append(deckOne[1:], drawOne, drawTwo)
			deckTwo = deckTwo[1:]
		} else {
			deckOne = deckOne[1:]
			deckTwo = append(deckTwo[1:], drawTwo, drawOne)
		}
	}

	var winningDeck []int
	var score int
	if len(deckOne) > 0 {
		winningDeck = deckOne
	} else {
		winningDeck = deckTwo
	}

	for i, v := range winningDeck {
		score += v * (len(winningDeck) - i)
	}

	fmt.Printf("Part One: Score of winning deck is %d\n", score)
}
