package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rules, messages := strings.Split(strings.Split(string(contents), "\n\n")[0], "\n"), strings.Split(strings.Split(string(contents), "\n\n")[1], "\n")
	rawRules := map[string]string{}

	for _, v := range rules {
		key, value := strings.Split(v, ":")[0], strings.TrimSpace(strings.Split(v, ":")[1])
		rawRules[key] = value
	}

	evaluatedRules := map[string][]string{}

	// Evaluate the entire chain
	evaluateSomething("0", rawRules, evaluatedRules)

	fmt.Println(evaluatedRules["42"])

	validMessages := 0
	for _, message := range messages {
		for _, isOk := range evaluatedRules["0"] {
			if message == isOk {
				validMessages += 1
				continue
			}
		}
	}

	fmt.Printf("Part one: %d valid messages.\n", validMessages)

}

func evaluateSomething(key string, raw map[string]string, parsed map[string][]string) []string {

	// first check if the key has already been evaluated
	value, exists := parsed[key]
	if exists {
		return value
	}

	// if not, proceed recursively with memo-ization

	// according to the problem statement, some rules consist of just a single letter "a" or "b"
	// and looking at the puzzle input, looks like I have one rule "a" and one rule "b"
	// so this is the end of the rabbit hole

	rv := strings.ReplaceAll(raw[key], "\"", "")

	if (rv == "a") || (rv == "b") {
		parsed[key] = []string{rv}
	} else {
		// make an array of strings representing a possible rule that a message can match
		out := []string{}
		// if there's an or, go over all possibilities
		for _, subPart := range strings.Split(rv, " | ") {
			// for lhs and rhs of the `or`, if they contain multiple keys, need to form
			// all combinations of possibilities from the first and second key.
			// combine by concatenating strings.

			// conceptually, it is easier for me to just check the two cases whether
			// there is one key or two keys
			if len(strings.Split(subPart, " ")) == 1 {
				out = append(out, evaluateSomething(subPart, raw, parsed)...)
			} else {
				// form all concatenations of items from parsed left and right
				left, right := strings.Split(subPart, " ")[0], strings.Split(subPart, " ")[1]
				for _, lv := range evaluateSomething(left, raw, parsed) {
					for _, rv := range evaluateSomething(right, raw, parsed) {
						out = append(out, lv+rv)
					}
				}
			}
		}
		parsed[key] = out
	}

	return parsed[key]
}
