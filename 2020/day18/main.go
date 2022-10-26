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
	equations := strings.Split(string(contents), "\n")

	var acc int
	var acc2 int
	for _, eq := range equations {
		acc += evaluate(eq)
		acc2 += evaluate2(eq)
	}
	fmt.Printf("Part One: Sum of all evaluted equations is %d\n", acc)
	fmt.Printf("Part Two: Sum of all evaluted equations is %d\n", acc2)
}

func evaluate(eq string) int {
	// I can think of one or two ways to do this but a simple approach
	// could be to use regexp's to iteratively evaluate innermost parentheses

	re := regexp.MustCompile(`\(([^\(\)]*)\)`) // Matches a `(` which reaches a `)` without any other `(` or `)` in between.
	for {
		if !strings.Contains(eq, "(") {
			break
		}
		subExpr := re.FindStringSubmatch(eq)[1]
		subResult := evaluateSimple(subExpr)
		eq = strings.Replace(eq, "("+subExpr+")", subResult, 1)
	}

	fin := evaluateSimple(eq)

	res, err := strconv.Atoi(fin)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func evaluateSimple(subexpr string) string {

	runes := strings.Split(subexpr, " ")
	acc, err := strconv.Atoi(runes[0])
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(runes)-1; i += 2 {
		nv, err := strconv.Atoi(runes[i+1])
		if err != nil {
			log.Fatal(err)
		}
		op := runes[i]

		if op == "+" {
			acc += nv
		} else if op == "-" {
			acc -= nv
		} else if op == "*" {
			acc *= nv
		} else if op == "/" {
			acc /= nv
		}
	}

	return fmt.Sprint(acc)
}

func evaluateSimple2(eq string) string {
	// Now we have to evaluate addition with precedence over multiplication.
	// We can implement another level of iterative regexp replacement.
	for {
		if !strings.Contains(eq, "+") {
			return evaluateSimple(eq)
		}
		re := regexp.MustCompile(`(\d* \+ \d*)`) // Matches "a + b" for a and b integers.
		subExpr := re.FindStringSubmatch(eq)[1]
		eq = strings.Replace(eq, subExpr, evaluateSimple(subExpr), 1)
	}

}

func evaluate2(eq string) int {
	re := regexp.MustCompile(`\(([^\(\)]*)\)`) // Matches a `(` which reaches a `)` without any other `(` or `)` in between.
	for {
		if !strings.Contains(eq, "(") {
			break
		}
		subExpr := re.FindStringSubmatch(eq)[1]
		subResult := evaluateSimple2(subExpr)
		eq = strings.Replace(eq, "("+subExpr+")", subResult, 1)
	}

	fin := evaluateSimple2(eq)

	res, err := strconv.Atoi(fin)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
