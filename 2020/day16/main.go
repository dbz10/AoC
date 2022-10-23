package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rules, myTicket, nearbyTickets := read(os.Args[1])
	invalid := []int{}

	validTickets := []ticket{}

	for _, ticket := range nearbyTickets {
		ticketOk := true
		for _, number := range ticket {
			ok := false
			for _, rule := range rules {
				if rule.checkNumber(number) {
					ok = true
				}
			}
			if !ok {
				invalid = append(invalid, number)
				ticketOk = false
			}
		}
		if ticketOk {
			validTickets = append(validTickets, ticket)
		}
	}

	scanningErrorRate := 0
	for _, v := range invalid {
		scanningErrorRate += v
	}

	fmt.Printf("Part One: Scanning error rate of %d\n", scanningErrorRate)

	fmt.Println("Number of valid tickets", len(validTickets))

	// Part Two: For all rules, basically we want to look for in which position,
	// all tickets have a valid entry. Though in practice this could be a little bit tricky...

	// To start, let me hope there is a unique fit for each field.
	// Spoiler, there isn't.

	fieldPossiblePositions := map[string][]int{}

	for _, rule := range rules {
		for position := 0; position < len(myTicket); position++ {
			positionOk := true
			for _, ticket := range append(validTickets, myTicket) {
				if !rule.checkNumber(ticket[position]) {
					positionOk = false
					break
				}
			}
			if positionOk {
				fieldPossiblePositions[rule.field] = append(fieldPossiblePositions[rule.field], position)
			}
		}
	}

	resolvedRules := []rule{}
	for _, r := range rules {
		resolvedRules = append(resolvedRules,
			rule{r.field,
				r.lowerRange,
				r.upperRange,
				fieldPossiblePositions[r.field]})
	}

	sort.Slice(resolvedRules, func(i, j int) bool {
		return len(resolvedRules[i].possiblePositions) < len(resolvedRules[j].possiblePositions)
	})

	available := map[int]bool{}
	for p := range myTicket {
		available[p] = true
	}

	fieldPositionMap := map[string]int{}

	for _, rule := range resolvedRules {
		for _, possiblePosition := range rule.possiblePositions {
			if available[possiblePosition] {
				fieldPositionMap[rule.field] = possiblePosition
				available[possiblePosition] = false
				break
			}
		}
	}

	departureCode := 1
	for key, index := range fieldPositionMap {
		if strings.HasPrefix(key, "departure") {
			departureCode *= myTicket[index]
		}
	}

	fmt.Printf("Part two: departure code is %d\n", departureCode)

}

func read(path string) ([]rule, ticket, []ticket) {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	splits := strings.Split(string(contents), "\n\n")

	rules := []rule{}
	myTicket := ticket{}
	nearbyTickets := []ticket{}

	for _, line := range strings.Split(splits[0], "\n") {
		rules = append(rules, ruleFrom(line))
	}

	for _, v := range strings.Split(strings.Split(splits[1], "\n")[1], ",") {
		v, _ := strconv.Atoi(v)
		myTicket = append(myTicket, v)
	}

	for _, l := range strings.Split(splits[2], "\n")[1:] {
		placeholder := ticket{}
		for _, v := range strings.Split(l, ",") {
			v, _ := strconv.Atoi(v)
			placeholder = append(placeholder, v)
		}
		nearbyTickets = append(nearbyTickets, placeholder)
	}

	return rules, myTicket, nearbyTickets
}

type fieldRange struct {
	min int
	max int
}

type rule struct {
	field             string
	lowerRange        fieldRange
	upperRange        fieldRange
	possiblePositions []int
}

type ticket []int

// class: 1-3 or 5-7

func ruleFrom(s string) rule {
	field := strings.Split(s, ":")[0]
	rhs := strings.Split(s, ":")[1]
	lowRange := strings.TrimSpace(strings.Split(rhs, "or")[0])
	highRange := strings.TrimSpace(strings.Split(rhs, "or")[1])

	lowMin, _ := strconv.Atoi(strings.Split(lowRange, "-")[0])
	lowMax, _ := strconv.Atoi(strings.Split(lowRange, "-")[1])
	highMin, _ := strconv.Atoi(strings.Split(highRange, "-")[0])
	highMax, _ := strconv.Atoi(strings.Split(highRange, "-")[1])

	return rule{field, fieldRange{lowMin, lowMax}, fieldRange{highMin, highMax}, []int{}}
}

func (r rule) checkNumber(n int) bool {
	return (n >= r.lowerRange.min && n <= r.lowerRange.max) || (n >= r.upperRange.min && n <= r.upperRange.max)
}
