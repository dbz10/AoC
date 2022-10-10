package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var highestSeatId, mySeat int
	var occupiedSeats []int

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lines := bufio.NewScanner(file)

	for lines.Scan() {
		stringBinary := convertLetterStringToBinaryString(lines.Text())
		rowNumber, err := strconv.ParseInt(stringBinary[:7], 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		columnNumber, err := strconv.ParseInt(stringBinary[7:], 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		seatId := rowNumber*8 + columnNumber
		occupiedSeats = append(occupiedSeats, int(seatId)) // for part 2
		if seatId > int64(highestSeatId) {
			highestSeatId = int(seatId)
		}
	}
	fmt.Printf("Part One: Highest seat id found was %d\n", highestSeatId)

	// Part two
	// Maybe there is a cleaner way to do this but this at least is
	// easy
	for possiblyMySeat := 1; possiblyMySeat < highestSeatId; possiblyMySeat++ {
		if !arrayContains(occupiedSeats, possiblyMySeat) {
			mySeat = possiblyMySeat
		}
	}

	fmt.Printf("Part Two: My seat should be... %d\n", mySeat)

}

func convertLetterStringToBinaryString(str string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(str, "F", "0"),
				"B", "1",
			),
			"L", "0",
		), "R", "1",
	)
}

func arrayContains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
