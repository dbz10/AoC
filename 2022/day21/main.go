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

	monkeys := MonkeyOps{
		monkeySee:         map[string]int{},
		monkeyDo:          map[string]string{},
		branchTouchedHumn: map[string]bool{},
	}

	for _, line := range strings.Split(string(contents), "\n") {
		monkey := strings.TrimSpace(strings.Split(line, ":")[0])
		monkeyDo := strings.TrimSpace(strings.Split(line, ":")[1])
		monkeys.monkeyDo[monkey] = monkeyDo
	}
	monkeys.branchTouchedHumn["humn"] = true
	freshMonkeys := copyMonkeys(monkeys) // lol.... making a copy for part 2

	fmt.Printf("Part One: root monkey is going to yell %d!\n", monkeys.monkey("root"))

	// Part Two, holy shit
	// okay, my initial attempt to find a pattern
	// in the derivative of the total failed.
	// well this might be stupid but i'm just
	// to try implementing an alternating jump and scan approach...?

	// it worked great!!!

	var left, right, humanSide, nonHumanSide string
	left = strings.Split(monkeys.monkeyDo["root"], " ")[0]
	right = strings.Split(monkeys.monkeyDo["root"], " ")[2]
	if monkeys.branchTouchedHumn[left] {
		humanSide = left
		nonHumanSide = right
	} else {
		humanSide = right
		nonHumanSide = left
	}

	want := monkeys.monkey(nonHumanSide)
	var got int
	alternateUniverse := copyMonkeys(freshMonkeys)
	alternateUniverse.monkeyDo["humn"] = "0"
	acc := []int{}
	scanLength := 100
	start := 0
outer:
	for {
		for i := start; i < start+scanLength; i++ {
			alternateUniverse := copyMonkeys(freshMonkeys)
			alternateUniverse.monkeyDo["humn"] = fmt.Sprint(i)
			got = alternateUniverse.monkey(humanSide)
			acc = append(acc, got)
			if got == want {
				fmt.Printf("Part Two: A starting value of %d will pass the equality test!\n", i)
				break outer
			}
		}
		slope := (float64(acc[len(acc)-1] - acc[0])) / float64(scanLength)
		guess := int(float64(want-got) / slope)
		start += guess
		acc = []int{}
		fmt.Println(start, want, got)
	}

	// and...
	alternateUniverse = copyMonkeys(freshMonkeys)
	alternateUniverse.monkeyDo["humn"] = "3451534022348"
	alternateUniverse.monkey("root")

}

// what a silly data structure
type MonkeyOps struct {
	monkeySee         map[string]int
	monkeyDo          map[string]string
	branchTouchedHumn map[string]bool
}

type Monkey struct {
	name              string
	value             int
	branchTouchedHumn bool
}

func (m MonkeyOps) monkey(monkey string) int {
	// DFS with memoization

	// if the description is just an int,
	// then return it's value
	v, exists := m.monkeySee[monkey]
	if exists {
		return v
	}

	monkeyDo := m.monkeyDo[monkey]

	var out int
	v, err := strconv.Atoi(strings.TrimSpace(monkeyDo))
	if err == nil {
		out = v
	} else {
		var m1, m2, op string
		fmt.Sscanf(monkeyDo, "%s %s %s", &m1, &op, &m2)
		if op == "+" {
			out = m.monkey(m1) + m.monkey(m2)
		} else if op == "-" {
			out = m.monkey(m1) - m.monkey(m2)
		} else if op == "*" {
			out = m.monkey(m1) * m.monkey(m2)
		} else {
			out = m.monkey(m1) / m.monkey(m2)
		}
		m.branchTouchedHumn[monkey] = m.branchTouchedHumn[m1] || m.branchTouchedHumn[m2]
		if monkey == "root" {
			left := Monkey{m1, m.monkeySee[m1], m.branchTouchedHumn[m1]}
			right := Monkey{m2, m.monkeySee[m2], m.branchTouchedHumn[m2]}
			fmt.Printf("Root Values: %v, %v\n", left, right)
		}
	}
	m.monkeySee[monkey] = out
	return out
}

func copyMonkeys(m MonkeyOps) MonkeyOps {
	out := MonkeyOps{
		monkeySee:         map[string]int{},
		monkeyDo:          map[string]string{},
		branchTouchedHumn: map[string]bool{},
	}
	for key := range m.monkeyDo {
		out.monkeyDo[key] = m.monkeyDo[key]
	}
	out.branchTouchedHumn["humn"] = true
	return out
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
