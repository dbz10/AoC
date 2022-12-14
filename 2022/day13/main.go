package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// A packet can have EITHER a value OR children, but not both
type packet struct {
	value    []int // just for the ability to have an empty slice, though it will have only 1 or zero elements
	children []*packet
	parent   *packet
}

func (p packet) isTerminal() bool {
	return len(p.children) == 0
}

func (p packet) valueNonEmpty() bool {
	return len(p.value) > 0
}

func (p packet) valueEmpty() bool {
	return !p.valueNonEmpty()
}

func (p packet) String() string {
	// Depth First
	ll := len(p.children)
	outputString := ""
	if p.isTerminal() && p.valueNonEmpty() {
		outputString += fmt.Sprint(p.value[0])
	} else if p.isTerminal() && p.valueEmpty() {
		outputString += ""
	} else {
		outputString += "["
		for i, child := range p.children {
			outputString += child.String()
			if i < ll-1 {
				outputString += ","
			}
		}
		outputString += "]"
	}
	return outputString
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	allPairs := strings.Split(string(contents), "\n\n")
	packets := [][2]packet{}
	for _, p := range allPairs {
		pairs := strings.Split(p, "\n")
		packets = append(packets, [2]packet{parsePacket(pairs[0]), parsePacket(pairs[1])})
	}

	results := []bool{}
	for _, p := range packets {
		res := compare(p[0], p[1])
		// Because I know 0 will not happen
		if res == 1 {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}

	var sum int
	for i := range results {
		if results[i] {
			sum += i + 1
		}
	}
	fmt.Println("Part One:", sum)

	// Part Two
	// is easy thanks to my great pain in part one
	flattened := []packet{}
	for _, p := range packets {
		flattened = append(flattened, p[0], p[1])
	}
	flattened = append(flattened, parsePacket("[[2]]"))
	flattened = append(flattened, parsePacket("[[6]]"))

	sort.Slice(flattened, func(i, j int) bool { return compare(flattened[i], flattened[j]) == 1 })
	var prod int = 1
	for i, f := range flattened {
		if fmt.Sprint(f) == "[[2]]" || fmt.Sprint(f) == "[[6]]" {
			prod *= (i + 1)
		}
	}
	fmt.Println("Part Two:", prod)
}

func parsePacket(s string) packet {
	p := packet{}
	segments := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, "[", "[,"), "]", ",]"), ",") // vomit
	currentPointer := &p
	for _, block := range segments {
		if block == "[" {
			childPointer := &packet{parent: currentPointer}
			currentPointer.children = append(currentPointer.children, childPointer)
			currentPointer = childPointer
		} else if block == "]" {
			currentPointer = currentPointer.parent
		} else if block == "" {
			childPointer := &packet{value: []int{}, parent: currentPointer}
			currentPointer.children = append(currentPointer.children, childPointer)
		} else {
			// then it had better be a number
			newVal, err := strconv.Atoi(block)
			if err != nil {
				log.Fatal(err)
			}
			childPointer := &packet{value: []int{newVal}, parent: currentPointer}
			currentPointer.children = append(currentPointer.children, childPointer)
		}
	}
	// Since we always have the outer containing list
	return *p.children[0]
}

func compare(l, r packet) int {
	// This function is absolutely horrible
	// Overall idea: As soon as we get a definitive comparison,
	// break out of everything and return it. Until then, continue.
	// It seems that structurally I have to allow for the possibility
	// That the entire comparison is inconclusive, even though it
	// might not ever happen in practice
	// So the output will be -1, 0, 1

	// If L and R are both terminal nodes, compare the values,
	// after checking that the values are not empty
	if l.isTerminal() && r.isTerminal() {
		return compareTerminalValues(l, r)
	} else if l.isTerminal() && !r.isTerminal() {
		// Then l is a terminal node and we need to compare the value to r's child
		if l.valueEmpty() {
			return 1
		}
		// otherwise, compare the value to (up to) r's first child-value
		cv := compare(l, *r.children[0])
		if cv == 1 {
			return 1
		} else if cv == 0 && len(r.children) > 1 {
			return 1
		} else {
			return -1
		}
	} else if !l.isTerminal() && r.isTerminal() {
		// Then r is a terminal node and we need to compare the value to l's child
		if r.valueEmpty() {
			return -1
		}
		// otherwise, compare the value to (up to) l's first child-value
		cv := compare(*l.children[0], r)
		if cv == 1 {
			return 1
		} else if cv == 0 && len(l.children) > 1 {
			return -1
		} else {
			return -1
		}
	} else {
		// Otherwise, we start comparing children one by one
		ll := len(l.children)
		rr := len(r.children)
		for i := 0; i < max(ll, rr); i++ {
			if ll == i {
				// Then L has run out of values
				return 1
			} else if rr == i {
				// Then R has run out of values
				return -1
			} else {
				cv := compare(*l.children[i], *r.children[i])
				if cv == 0 {
					continue
				}
				return cv
			}
		}
		return 0
	}
}

func compareTerminalValues(l, r packet) int {
	if l.valueNonEmpty() && r.valueNonEmpty() {
		return tril(l.value[0], r.value[0])
	} else if l.valueNonEmpty() && r.valueEmpty() {
		return -1
	} else if l.valueEmpty() && r.valueNonEmpty() {
		return 1
	}
	return 0
}

func tril(a, b int) int {
	if a < b {
		return 1
	} else if b < a {
		return -1
	}
	return 0
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
