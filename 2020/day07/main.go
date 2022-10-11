package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// This was a bit painful and I don't think my solution is very clean :(
// I struggled a bit because I felt like I wanted to pass a bag struct
//
//	type bag struct {
//		color string
//		contains []bag
//	}
//
// around by reference so that I could update it inplace.
// In the end I just used a map, and looked up bags in the map
// by their color as an id, which felt quite unsatisfying

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lines := bufio.NewScanner(file)

	var bagsSpec = make(map[string]bag)
	for lines.Scan() {
		updateBags(lines.Text(), bagsSpec)
	}

	var numBagsCanHoldShinyGold int
	for _, v := range bagsSpec {
		if canHoldShinyGold(v, bagsSpec) {
			numBagsCanHoldShinyGold++
		}
	}

	// Minus one to exclude the shiny gold bag itself.
	fmt.Printf("Part One: %d different colored bags can eventually hold a shiny gold bag\n", numBagsCanHoldShinyGold-1)
	// Minus one for the original shiny bag, as bagOfHolding essentially returns the total number of bags we are holding.
	fmt.Printf("Part Two: One shiny gold bag must contain %d other bags\n", bagOfHolding("shiny gold", bagsSpec)-1)

}

type bag struct {
	color    string
	contains map[string]int
}

func updateBags(line string, bagsDescription map[string]bag) {
	thisBagColor := strings.Split(line, " bags contain ")[0]
	var thisBagContains = make(map[string]int)
	var otherBagColor string
	for _, seg := range strings.Split(strings.Split(line, "bags contain")[1], ",") {
		colorRe := regexp.MustCompile(`\d+ (.*) bag`)
		numRe := regexp.MustCompile(`(\d+) .* bag`)
		matchResult := colorRe.FindStringSubmatch(seg)
		if len(matchResult) > 1 {
			otherBagColor = matchResult[1]
			numOfThatBag, _ := strconv.Atoi(numRe.FindStringSubmatch(seg)[1])
			thisBagContains[otherBagColor] = numOfThatBag
		}
	}
	bagsDescription[thisBagColor] = bag{thisBagColor, thisBagContains}
}

func canHoldShinyGold(b bag, allBags map[string]bag) bool {
	if b.color == "shiny gold" {
		return true
	}
	checkRest := false
	for k := range b.contains {
		if canHoldShinyGold(allBags[k], allBags) {
			checkRest = true
		}
	}
	return checkRest
}

func bagOfHolding(color string, allBags map[string]bag) int {
	var init int = 1
	for containedBag, num := range allBags[color].contains {
		init += num * bagOfHolding(containedBag, allBags)
	}
	return init
}
