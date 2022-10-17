package main

import (
	"errors"
	"fmt"
	"log"
	"math"
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
	lineNumber := []int{}
	var ln int
	for _, ch := range strings.Split(strings.Split(string(contents), "\n")[1], ",") {
		if ch != "x" {
			lineNo, err := strconv.Atoi(ch)
			if err != nil {
				log.Fatal(err)
			}
			busLines = append(busLines, lineNo)
			lineNumber = append(lineNumber, ln)
		}
		ln++
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
	fmt.Printf("Part One: %d.\n", earliestTimestampId)

	// Part two is somewhat nontrivial....
	// I think this is roughly how the CRT approach goes
	// Man this was no joke.

	mAll := float64(1)
	baseFactors := []int64{}
	for i := range busLines {
		mAll *= float64(busLines[i])
		base := int64(1)
		for _, v := range busLines[:i] {
			base *= int64(v)
		}
		for _, v := range busLines[i+1:] {
			base *= int64(v)
		}
		baseFactors = append(baseFactors, base)

	}

	solvedFactors := []int64{}

	// Now for each after the first element, brute force search
	// for multiple to get the desired mod after division by the value.

	// should have t + l mod b = 0
	// in other words t mod b = -l
	// in other words t mod b = b - l
	for i := range busLines {

		multiple, err := searchForModMultiplier(
			baseFactors[i],
			busLines[i],
			busLines[i]-lineNumber[i])
		if err != nil {
			log.Fatal(err)
		}
		solvedFactors = append(solvedFactors, baseFactors[i]*int64(multiple))
	}

	var answerMaybe float64
	for _, v := range solvedFactors {
		answerMaybe += float64(v)
	}

	answerMaybe = answerMaybe - math.Floor(answerMaybe/float64(mAll))*float64(mAll)

	fmt.Printf("Part two, could it be... %d\n", int(answerMaybe))

}

func searchForModMultiplier(start int64, divisor int, targetModulus int) (int, error) {
	// Of course this can be done more efficiently with modular arithmetic
	// but let me get any kind of working solution first

	decomposedTargetModulus := int(math.Mod(float64(targetModulus), float64(divisor)))
	if decomposedTargetModulus < 0.0 {
		decomposedTargetModulus += divisor
	}

	for i := int64(1); i <= int64(divisor); i++ {
		if int(math.Mod(float64(start*i), float64(divisor))) == decomposedTargetModulus {
			return int(i), nil
		}
	}
	return -1, errors.New("could not find a value, for some reason")
}
