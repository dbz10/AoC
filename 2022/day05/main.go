package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Utilize slices as a stack
type stack []string

func (s *stack) pop(n ...int) []string {
	if len(n) > 0 {
		nEle := n[0]
		if len(*s) < nEle {
			log.Fatalf("cannot pop %d items from a stack which only contains %d", n, len(*s))
		}
		ele := (*s)[len(*s)-nEle:]
		*s = (*s)[:len(*s)-nEle]
		return ele
	}

	ele := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return []string{ele}
}

func moveOneByOne(from, to *stack, n int) {
	if len(*from) < n {
		log.Fatalf("cannot move %d items from a stack which only contains %d", n, len(*from))
	}
	for i := 0; i < n; i++ {
		*to = append(*to, from.pop()...)
	}
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	initialArrangement := strings.Split(string(contents), "\n\n")[0]
	instructions := strings.Split(string(contents), "\n\n")[1]
	stacksPartOne := parseStartingPositions(initialArrangement)
	stacksPartTwo := parseStartingPositions(initialArrangement)

	for _, instr := range strings.Split(instructions, "\n") {
		var n, from, to int
		fmt.Sscanf(instr, "move %d from %d to %d", &n, &from, &to)

		moveOneByOne(&stacksPartOne[from-1], &stacksPartOne[to-1], n)
		stacksPartTwo[to-1] = append(stacksPartTwo[to-1], stacksPartTwo[from-1].pop(n)...)
	}

	fmt.Print("Part One: After the crane operation, the top of each stack is ")
	for _, s := range stacksPartOne {
		fmt.Print(s[len(s)-1])
	}
	fmt.Print("\n")

	fmt.Print("Part Two: After the crane operation, the top of each stack is ")
	for _, s := range stacksPartTwo {
		fmt.Print(s[len(s)-1])
	}
	fmt.Print("\n")

}

func parseStartingPositions(cf string) []stack {
	// Expect something like
	//     [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3

	// So after dropping the last line, we have [X] separated by spaces

	lines := strings.Split(cf, "\n")
	containerLines := lines[:len(lines)-1]
	headerLine := lines[len(lines)-1]

	// Figure out how many stacks we need
	nStacks, err := strconv.Atoi(string(headerLine[len(headerLine)-2]))
	var stacks = make([]stack, nStacks)

	if err != nil {
		log.Fatal(err)
	}

	// By manual inspection, the label for box i occurs at character position 4i+1
	// Fill up the stacks, in reverse order
	for i := len(containerLines) - 1; i >= 0; i-- {
		line := containerLines[i]
		for j := 0; j < nStacks; j++ {
			if string(line[4*j+1]) != " " {
				stacks[j] = append(stacks[j], string(line[4*j+1]))
			}
		}
	}

	return stacks

}
