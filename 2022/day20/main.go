package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// maybe a doubly linked list?
	// going to need another copy for part 2
	nodes := map[int]*node{}
	nodes2 := map[int]*node{}
	// parse the input
	ints := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, v)
	}
	// into a DLL. remember where zero is
	zeroIndex := -1
	for i, v := range ints {
		nodes[i] = &node{value: v}
		nodes2[i] = &node{value: v * 811589153}
		if v == 0 {
			zeroIndex = i
		}
	}
	var lindex, rindex int
	for i := range ints {
		if i == 0 {
			lindex = len(ints) - 1
			rindex = 1
		} else if i == len(ints)-1 {
			lindex = len(ints) - 2
			rindex = 0
		} else {
			lindex = i - 1
			rindex = i + 1
		}
		nodes[i].l = nodes[lindex]
		nodes[i].r = nodes[rindex]

		nodes2[i].l = nodes2[lindex]
		nodes2[i].r = nodes2[rindex]
	}

	for i, v := range ints {
		// yeah i needed a little bit of a hint
		// to realize about the wrapping aspect -
		// no wrapping in the sample but there
		// is wrapping around in the input :(
		nodes[i].stitch(v % (len(ints) - 1))
	}

	var coordinateSum int
	cur := nodes[zeroIndex]
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			cur = cur.r
		}
		fmt.Printf("at position %d, value %d\n", i*1000, cur.value)
		coordinateSum += cur.value
	}
	fmt.Printf("Part One: Sum of grove coordinates is %d\n", coordinateSum)

	// Part Two. ok, but ok
	for round := 0; round < 10; round++ {
		for i := range ints {
			nodes2[i].stitch(nodes2[i].value % (len(ints) - 1))
		}
	}
	cur = nodes2[zeroIndex]
	coordinateSum = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			cur = cur.r
		}
		fmt.Printf("at position %d, value %d\n", i*1000, cur.value)
		coordinateSum += cur.value
	}
	fmt.Printf("Part One: Sum of decrypted grove coordinates is %d\n", coordinateSum)

}

type node struct {
	value int
	l, r  *node
}

func (n *node) String() string {
	return fmt.Sprintf("Left: %d, Value: %d, Right: %d", n.value, n.l.value, n.r.value)
}

func (self *node) stitch(n int) {
	// interleave self between the appropriate
	// neighbors n steps to the right or left
	// for positive or negative n, respectively

	// careful accounting to avoid off by
	// one errors:
	// when moving right, move becomes self's
	// new left, and move.r becomes self's
	// when moving left, to have the same convention,
	// we actually have to go one step further

	// 1 2 -3 -> 2 1 -3
	// 5 -1 2 -> -1 5 2 = 2 -1 5
	// 2 is 2 left of -1 in the original setup
	if n == 0 {
		return
	}
	move := self
	if n > 0 {
		for i := 0; i < n; i++ {
			move = move.r
		}
	} else if n < 0 {
		for i := 0; i < -n+1; i++ {
			move = move.l
		}

	}
	// little bit tricky here
	// first connect across where self used to be
	self.l.r = self.r
	self.r.l = self.l

	// now move self into its new place
	self.l = move
	self.r = move.r

	// now tell neighbors that self has entered
	self.l.r = self
	self.r.l = self
}
