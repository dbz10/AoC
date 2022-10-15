package main

import (
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

	eeta, err := strconv.Atoi(strings.Split(string(contents), "\n")[0])
	if err != nil {
		log.Fatal(err)
	}

	busLines := []int{}
	for _, ch := range strings.Split(strings.Split(string(contents), "\n")[1], ",") {
		if ch != "x" {
			lineNo, err := strconv.Atoi(ch)
			if err != nil {
				log.Fatal(err)
			}
			busLines = append(busLines, lineNo)
		}
	}

	shortestWait := (busLines[0]*(eeta/busLines[0]+1) - eeta)
	earliestTimestampId := busLines[0] * shortestWait
	for _, v := range busLines {
		wait := v*(eeta/v+1) - eeta
		if wait < shortestWait {
			shortestWait = wait
			earliestTimestampId = wait * v
		}
	}
	println(earliestTimestampId)

	// Part two is somewhat nontrivial....
}
