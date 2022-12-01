package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	carries := strings.Split(string(contents), "\n\n")
	calories := []int{}
	maxIndex := 0
	maxValue := 0

	for index, backpack := range carries {
		total := 0
		for _, item := range strings.Split(backpack, "\n") {
			cals, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			total += cals
		}
		calories = append(calories, total)
		if total > maxValue {
			maxIndex = index
			maxValue = total
		}
	}

	fmt.Printf("Part One: Elf %d is carrying the most calories with %d\n", maxIndex, maxValue)

	sort.Ints(calories)

	fmt.Printf("Part Two: The top 3 elves are carrying %d calories respectively, %d in sum\n", calories[len(calories)-3:], calories[len(calories)-3]+calories[len(calories)-2]+calories[len(calories)-1])
}
