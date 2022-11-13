package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// instantiate all of our cups
	// in two passes for simplicity
	cups := map[int]cup{}
	cs := strings.Split(string(contents), "")
	for _, s := range cs {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ec := cup{id: id, next: &cup{}}
		cups[id] = ec
	}

	for i, _ := range cs {
		idThis, _ := strconv.Atoi(cs[i])
		idNext, _ := strconv.Atoi(cs[int(math.Mod(float64(i+1), 9))])
		*cups[idThis].next = cups[idNext]
	}

	// starting cup
	startingCup, _ := strconv.Atoi(cs[0])
	currentCup := cups[startingCup]

	for i := 0; i < 100; i++ {
		pickUpIds := []int{currentCup.next.id,
			currentCup.next.next.id,
			currentCup.next.next.next.id}
		destinationCupId := modToOne(currentCup.id-1, 9)

		for contains(pickUpIds, destinationCupId) {
			destinationCupId = modToOne(destinationCupId-1, 9)
		}

		// now reroute links
		*currentCup.next = cups[currentCup.next.next.next.next.id]
		*cups[pickUpIds[2]].next = cups[cups[destinationCupId].next.id]
		*cups[destinationCupId].next = cups[pickUpIds[0]]
		currentCup = *currentCup.next
	}

	currentCup = cups[1]

	fmt.Println("Part One: ")
	for i := 0; i < 9; i++ {
		fmt.Print(currentCup.id)
		currentCup = *currentCup.next
	}
	fmt.Print("\n")

}

// a linked list seems okay here
type cup struct {
	id   int
	next *cup
}

func contains(arr []int, v int) bool {
	for _, c := range arr {
		if c == v {
			return true
		}
	}
	return false
}

func modToOne(a, b int) int {
	if a <= 0 {
		a = a + b
	}
	return int(math.Mod(float64(a-1), float64(b)) + 1)
}
