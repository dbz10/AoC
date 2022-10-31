package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	// What the hell is part two...

	// Ok. We have 0: 8 11
	// with 8: 42 | 42 8
	// and 11: 42 31 | 42 11 31

	// That means the message can follow the format
	// [any of 42] repeated any number of times + [42 31] nested any amount of times

	// Start by evaluating 42 and 31

	fortyTwo := evaluateSomething("42", rawRules, evaluatedRules)
	thirtyOne := evaluateSomething("31", rawRules, evaluatedRules)

	// Now reduce each message in two stages
	// 1. regexp replace [any of 42] 1 or more times to empty string
	// 2. while string not unchanged, regexp replace $[any of 42](.*)[any of 31]^/$1

	// This feels like a complete failure, but at least it got the job done.

	var validMessages int

	for _, message := range messages {
		for repetitions := 1; repetitions < 20; repetitions++ {
			reEightString := `^(?:` + strings.Join(fortyTwo, "|") + `){` + strconv.Itoa(repetitions) + `}`
			reElevenString := `(?:` + strings.Join(fortyTwo, "|") + `)` + `(.*)` + `(?:` + strings.Join(thirtyOne, "|") + `)$`
			reToEndItAllString := reEightString + reElevenString
			reEleven := regexp.MustCompile(reElevenString)
			reToEndItAll := regexp.MustCompile(reToEndItAllString)
			reducedLevelOne := reToEndItAll.FindStringSubmatch(message)
			if reducedLevelOne == nil {
				continue
			} else {
				next := reducedLevelOne[1]
				for i := 0; i < 5; i++ {
					if len(next) == 0 {
						validMessages++
						break
					} else {
						reResult := reEleven.FindStringSubmatch(next)
						if reResult == nil {
							break
						} else {
							next = reResult[1]
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part two: %d valid messages.\n", validMessages)

}

func contains(arr []string, s string) bool {
	res := false
	for _, v := range arr {
		if v == s {
			res = true
		}
	}
	return res
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
