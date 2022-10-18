package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	memoryPartOne := map[int]int{}

	// Part two... what the hell man...
	memoryPartTwo := map[int]int{}

	var activate, deactivate int
	var maskStr string

	for scanner.Scan() {
		line := scanner.Text()
		// If the line begins with mask, update our current masks
		if strings.HasPrefix(line, "mask") {
			maskStr = strings.Split(line, " = ")[1]
			activate = 0
			deactivate = 0
			for i, v := range maskStr {
				if string(v) == "1" {
					activate += int(math.Pow(2, float64(35-i)))
				}
				if string(v) == "0" {
					deactivate += int(math.Pow(2, float64(35-i)))
				}

			}

		} else {
			baseNum, err := strconv.Atoi(strings.Split(line, " = ")[1])
			if err != nil {
				log.Fatal(err)
			}

			re := regexp.MustCompile(`mem\[(\d*)\]`)

			memAddress, err := strconv.Atoi(re.FindStringSubmatch(line)[1])
			if err != nil {
				log.Fatal(err)
			}
			memoryPartOne[memAddress] = (baseNum | activate) & ^deactivate

			baseAddress := memAddress | activate

			for _, address := range unNestMemoryAddresses(baseAddress, maskStr) {
				memoryPartTwo[(address | activate)] = baseNum
			}

		}

	}
	var memSum int
	for _, v := range memoryPartOne {
		memSum += v
	}
	fmt.Printf("Part One: Found sum of element in bitmask %d\n", memSum)

	memSum = 0
	for _, v := range memoryPartTwo {
		memSum += v
	}
	fmt.Printf("Part Two: Found sum of element in bitmask %d\n", memSum)

}

func unNestMemoryAddresses(baseAddress int, bitmask string) []int {
	leading := string(bitmask[0])
	position := utf8.RuneCountInString(bitmask) - 1
	var addresses []int

	if position == 0 {
		if leading == "X" {
			{
				addresses = append(addresses, baseAddress|int(math.Pow(2, float64(position))))
				addresses = append(addresses, baseAddress & ^int(math.Pow(2, float64(position))))
			}
		} else {
			addresses = append(addresses, baseAddress)
		}
	} else {
		for _, addr := range unNestMemoryAddresses(baseAddress, bitmask[1:]) {
			if leading == "X" {
				{
					addresses = append(addresses, addr|int(math.Pow(2, float64(position))))
					addresses = append(addresses, addr & ^int(math.Pow(2, float64(position))))
				}
			} else {
				addresses = append(addresses, addr)
			}
		}
	}

	return addresses

}
