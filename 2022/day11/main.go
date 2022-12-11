package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type queue []int

func (q *queue) push(v int) {
	*q = append(*q, v)
}

func (q *queue) pop() int {
	if len(*q) == 0 {
		fmt.Fprintf(os.Stderr, "cannot pop from empty queue")
	}
	ele := (*q)[0]
	(*q) = (*q)[1:]
	return ele
}

type monkey struct {
	q                queue
	itemsInspected   int
	operation        func(int) int
	chooseNextMonkey func(int) int
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	monkeys := loadMonkeys(contents)

	rounds := 20
	for i := 0; i < rounds; i++ {
		for m := range monkeys {
			for len(monkeys[m].q) > 0 {
				monkeys[m].itemsInspected++
				v := monkeys[m].q.pop()
				v = monkeys[m].operation(v)
				nm := monkeys[m].chooseNextMonkey(v)
				monkeys[nm].q.push(v)
			}
		}
	}

	// Sort Descending
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].itemsInspected > monkeys[j].itemsInspected })
	monkeyBusiness := monkeys[0].itemsInspected * monkeys[1].itemsInspected
	fmt.Printf("Part One: monkey business level after 20 rounds is %d\n", monkeyBusiness)

	// Part Two

}

func loadMonkeys(contents []byte) []monkey {
	monkeys := []monkey{}

	for _, block := range strings.Split(string(contents), "\n\n") {
		m := monkey{}
		blockLines := strings.Split(block, "\n")
		startingItemsLine := blockLines[1]
		operationLine := blockLines[2]
		testLine := blockLines[3]
		ifTrue := blockLines[4]
		ifFalse := blockLines[5]

		for _, i := range strings.Split(strings.Split(startingItemsLine, ":")[1], ",") {
			item, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				log.Fatal(err)
			}
			m.q = append(m.q, item)
		}

		var op string
		var constant int
		if strings.TrimSpace(operationLine) == "Operation: new = old * old" {
			m.operation = func(i int) int { return (i * i / 3) }
		} else {
			fmt.Sscanf(strings.TrimSpace(operationLine), "Operation: new = old %s %d", &op, &constant)
			if op == "*" {
				m.operation = func(i int) int { return (i * constant / 3) }
			} else {
				m.operation = func(i int) int { return (i + constant) / 3 }
			}
		}

		var divisor int
		fmt.Sscanf(strings.TrimSpace(testLine), "Test: divisible by %d", &divisor)
		var monkeyT, monkeyF int
		fmt.Sscanf(strings.TrimSpace(ifTrue), "If true: throw to monkey %d", &monkeyT)
		fmt.Sscanf(strings.TrimSpace(ifFalse), "If false: throw to monkey %d", &monkeyF)

		m.chooseNextMonkey = func(i int) int {
			if i%divisor == 0 {
				return monkeyT
			}
			return monkeyF
		}

		monkeys = append(monkeys, m)
	}

	return monkeys
}
