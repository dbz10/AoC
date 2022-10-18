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
	startingNumbers := []int{}
	for _, char := range strings.Split(string(contents), ",") {
		v, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		startingNumbers = append(startingNumbers, v)
	}

	memo := map[int]int{}
	var nowSaid, previousSaid int
	for i := 0; i < 30000000; i++ {
		if i < len(startingNumbers) {
			nowSaid = startingNumbers[i]
		} else {
			res, exists := memo[previousSaid]
			if !exists {
				nowSaid = 0
			} else {
				nowSaid = i - res
			}
		}

		memo[previousSaid] = i
		previousSaid = nowSaid

	}

	fmt.Printf("Part One or Two: The Nth number is %d\n", nowSaid)
}
