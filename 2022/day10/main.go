package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var scanner = bufio.NewScanner(file)
	var signal = []int{}
	var currentValue int = 1
	var instruction string

	for scanner.Scan() {
		instruction = scanner.Text()
		if instruction == "noop" {
			signal = append(signal, currentValue)
		} else {
			delta, err := strconv.Atoi(strings.Split(instruction, " ")[1])
			if err != nil {
				log.Fatal(err)
			}
			signal = append(signal, currentValue, currentValue)
			currentValue += delta
		}
	}

	result := (signal[19]*20 +
		signal[59]*60 +
		signal[99]*100 +
		signal[139]*140 +
		signal[179]*180 +
		signal[219]*220)
	fmt.Printf("Part One: sum of specified signal strengths is %d\n", result)

	// Part Two
	fmt.Println(len(signal))
	screenWidth := 40
	for i, x := range signal {
		hor := i % screenWidth
		if hor >= x-1 && hor <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if hor == 39 {
			fmt.Print("\n")
		}
	}
}
