package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func (p *point) move(dx, dy int) {
	p.x += dx
	p.y += dy
}

func offset(p1, p2 point) (int, int) {
	return p2.x - p1.x, p2.y - p1.y
}

func (t *point) follow(h point) {
	dx, dy := offset(*t, h)
	if dx >= 2 {
		if dy == 0 {
			t.move(1, 0)
		} else if dy > 0 {
			t.move(1, 1)
		} else if dy < 0 {
			t.move(1, -1)
		}
	} else if dx <= -2 {
		if dy == 0 {
			t.move(-1, 0)
		} else if dy > 0 {
			t.move(-1, 1)
		} else if dy < 0 {
			t.move(-1, -1)
		}
	} else if dy >= 2 {
		if dx == 0 {
			t.move(0, 1)
		} else if dx > 0 {
			t.move(1, 1)
		} else if dx < 0 {
			t.move(-1, 1)
		}
	} else if dy <= -2 {
		if dx == 0 {
			t.move(0, -1)
		} else if dx > 0 {
			t.move(1, -1)
		} else if dx < 0 {
			t.move(-1, -1)
		}
	}

}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var (
		head, tail point
		dir        string
		nSteps     int
		hashSet1   = make(map[point]struct{})
		hashSet2   = make(map[point]struct{})
	)

	for _, instruction := range strings.Split(string(contents), "\n") {
		fmt.Sscanf(instruction, "%s %d", &dir, &nSteps)
		for i := 0; i < nSteps; i++ {
			if dir == "U" {
				head.move(0, 1)
			} else if dir == "D" {
				head.move(0, -1)
			} else if dir == "R" {
				head.move(1, 0)
			} else if dir == "L" {
				head.move(-1, 0)
			}

			tail.follow(head)
			hashSet1[tail] = struct{}{}
		}
	}

	fmt.Printf("Part One: %d tiles visited by the tail of the rope\n", len(hashSet1))

	// Part Two
	ropeLength := 10
	var rope = make([]point, ropeLength)
	for _, instruction := range strings.Split(string(contents), "\n") {
		fmt.Sscanf(instruction, "%s %d", &dir, &nSteps)

		for i := 0; i < nSteps; i++ {
			if dir == "U" {
				rope[0].move(0, 1)
			} else if dir == "D" {
				rope[0].move(0, -1)
			} else if dir == "R" {
				rope[0].move(1, 0)
			} else if dir == "L" {
				rope[0].move(-1, 0)
			}

			for i := 1; i < ropeLength; i++ {
				rope[i].follow(rope[i-1])
			}

			hashSet2[rope[len(rope)-1]] = struct{}{}
		}
	}

	fmt.Printf("Part Two: %d tiles visited by the tail of the rope\n", len(hashSet2))

}
