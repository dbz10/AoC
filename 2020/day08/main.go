package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Not sure if my solution is that "nice"
// But it gets the job done.

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var instructions = strings.Split(string(contents), "\n")

	var visited = make(map[int]bool)
	var curLine, acc, delta, jump int

	for {
		_, exists := visited[curLine]
		if exists {
			break
		}
		delta, jump, err = interpret(parse(instructions[curLine]))
		if err != nil {
			log.Fatal(err)
		}
		acc += delta
		visited[curLine] = true
		curLine += jump
	}
	fmt.Printf("Part One: The value of the accumulator right before re-visiting an instruction is %d.\n", acc)

	// Part Two
	// Start from the beginning again.

	// We know that either
	// * we can change a jmp <-> nop and end right away
	// * we can change a jmp <-> nop and move to a instruction which can eventually terminate
	// In the latter case, we know that we didn't visit during part one, otherwise the program could have completed.
	// So then I think we just need to explore the rest of the path every time we encounter scenario B.

	curLine = 0
	acc = 0
	delta = 0
	jump = 0
	nInstructions := len(instructions)
	var i instruction
	// Not totally clear to me whether I can re-use the visited nodes from part one
	// so to be safe I will just start from scratch
	var visitedPartTwo = make(map[int]bool)

	// just to prevent an infinite loop
	for niter := 0; niter < 1e6; niter++ {
		i = parse(instructions[curLine])
		delta, jump, err = interpret(i)
		if err != nil {
			log.Fatal(err)
		}
		acc += delta

		// Check if we can escape by exchanging jmp <-> nop
		if i.action == "jmp" {

			if curLine == nInstructions-1 {
				// Then we can change this jmp to a nop and get to the next line,
				// which is off the end of the program
				break
			}

			// Otherwise, check whether going to the next line instead of jumping could result in us leaving the program
			hypotheticalNextLine := curLine + 1
			_, exists := visitedPartTwo[hypotheticalNextLine]
			if !exists {
				ok, delta := canEnd(hypotheticalNextLine, instructions, visitedPartTwo)
				if ok {
					acc += delta
					break
				}
			}
		} else if i.action == "nop" {
			if curLine+i.num == nInstructions {
				// Then we can change this nop to a jump and jump to the line
				// right after the end of the program
				break
			}

			// Otherwise, check whether jumping instead of going to the next line could result in us leaving the program
			hypotheticalNextLine := curLine + jump
			_, exists := visitedPartTwo[hypotheticalNextLine]
			if !exists {
				ok, delta := canEnd(hypotheticalNextLine, instructions, visitedPartTwo)
				if ok {
					acc += delta
					break
				}
			}

		}
		visited[curLine] = true
		curLine += jump
	}

	fmt.Printf("Part Two: The value of the accumulator after completing the (fixed) program %d.\n", acc)

}

type instruction struct {
	action string
	num    int
}

func parse(s string) instruction {
	split := strings.Split(s, " ")
	action := split[0]
	jump, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	return instruction{action: action, num: jump}
}

func interpret(i instruction) (int, int, error) {
	// Returns by how much to modify the accumulator, and how many lines to jump
	switch i.action {
	case "acc":
		return i.num, 1, nil
	case "nop":
		return 0, 1, nil
	case "jmp":
		return 0, i.num, nil
	default:
		return 0, 0, errors.New("supplied action did not match a known action: [jmp, acc, nop]")
	}
}

func canEnd(lineNumber int, instructions []string, scan map[int]bool) (bool, int) {
	// If true, return the additional amount of accumulator acquired along the way
	// By keeping the values in scan, in theory we never have to re-check a particular
	// line.
	var accDelta int
	for {
		_, exists := scan[lineNumber]
		if exists {
			return false, 0
		}
		delta, jump, err := interpret(parse(instructions[lineNumber]))
		if err != nil {
			log.Fatal(err)
		}
		accDelta += delta
		scan[lineNumber] = true

		if lineNumber+jump == len(instructions) {
			return true, accDelta
		}
		lineNumber += jump
	}

}
