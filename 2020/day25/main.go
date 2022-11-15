package main

import (
	"errors"
	"fmt"
	"log"
)

// I hope go is fast because I have no clue how to
// approach this other than brute force
func main() {

	cardPublicKey := 14012298
	doorPublicKey := 74241

	cardLoopSecret, err := bruteForceCrack(7, cardPublicKey)
	if err != nil {
		log.Panic(err)
	}

	encryptionKey := encrypt(doorPublicKey, cardLoopSecret)

	fmt.Printf("Part One: The encryption key is %d\n", encryptionKey)

}

func bruteForceCrack(subjectNumber int, publicKey int) (int, error) {
	// brute force to find the loop size
	// despite having the encrypt function to encrypt for a known loop size,
	// this cycle is way faster for cracking the secret loop count
	loopCounter := 0
	loopValue := 1
	var t int
	for loopCounter < 10000000 {
		if loopValue == publicKey {
			return loopCounter, nil
		} else {
			t = loopValue * subjectNumber
			loopValue = t - (t/20201227)*20201227
		}
		loopCounter++
	}
	return -1, errors.New("could not crack the public key within loop limit")
}

func encrypt(subjectNumber int, loopCounter int) int {
	loopValue := 1
	var t int
	for loop := 0; loop < loopCounter; loop++ {
		t = loopValue * subjectNumber
		loopValue = t - (t/20201227)*20201227
	}
	return loopValue
}
