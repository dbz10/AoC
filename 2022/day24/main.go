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

	blizzards := []blizzard{}
	walls := map[coord]bool{}
	north := coord{0, -1}
	south := coord{0, 1}
	west := coord{-1, 0}
	east := coord{1, 0}

	for y, row := range strings.Split(string(contents), "\n") {
		for x, ch := range strings.Split(row, "") {
			if ch == "v" {
				blizzards = append(blizzards, blizzard{loc: coord{x, y}, direction: south})
			}
			if ch == ">" {
				blizzards = append(blizzards, blizzard{loc: coord{x, y}, direction: east})
			}
			if ch == "<" {
				blizzards = append(blizzards, blizzard{loc: coord{x, y}, direction: west})
			}
			if ch == "^" {
				blizzards = append(blizzards, blizzard{loc: coord{x, y}, direction: north})
			}
			if ch == "#" {
				walls[coord{x, y}] = true
			}
		}
	}

	ly := len(strings.Split(string(contents), "\n"))
	lx := len(strings.Split(string(contents), "\n")[0])

	start := coord{1, 0}
	end := coord{lx - 2, ly - 1}

	now := []coord{start}
	goal := end
	haveForgottenSnacks := false

outer:
	for round := 1; round < 100000; round++ {
		next := []coord{}

		// The blizzards move. After they move,
		// check where I am allowed to be standing on the next turn.

		for i := range blizzards {
			blizzards[i].move(lx, ly)
		}

		blocked := isBlocked(walls, blizzards)

		for _, cur := range now {
			for _, test := range []coord{
				cur,
				cur.plus(north),
				cur.plus(south),
				cur.plus(east),
				cur.plus(west),
			} {
				if !blocked[test] && !contains(next, test) && test.y >= 0 && test.y < ly {
					next = append(next, test)

				}
			}
		}

		if goal == end && contains(next, goal) && !haveForgottenSnacks {
			fmt.Printf("Part One: Took %d turns to traverse the plateau. Need to return for the snacks though...\n", round)
			next = []coord{end}
			goal = start
		}
		if contains(next, end) && haveForgottenSnacks {
			fmt.Printf("Part Two: Took %d turns to traverse the plateau three times to come bak with the snacks.\n", round)
			break outer
		}

		if contains(next, goal) {
			fmt.Printf("Aquired the left behind snacks on round %d\n", round)
			haveForgottenSnacks = true
			next = []coord{start}
			goal = end
		}

		now = next
	}

}

type coord struct {
	x, y int
}

func contains(arr []coord, a coord) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func (b *blizzard) move(lx, ly int) {
	// by manual inspedtion, there is no blizzard heading up or down on the entry or end row
	lxl := lx - 2
	lyl := ly - 2

	b.loc.x = (b.loc.x+b.direction.x-1+lxl)%lxl + 1
	b.loc.y = (b.loc.y+b.direction.y-1+lyl)%lyl + 1
}

func (c coord) plus(o coord) coord {
	return coord{c.x + o.x, c.y + o.y}
}

type blizzard struct {
	loc       coord
	direction coord
}

func isBlocked(walls map[coord]bool, blizzards []blizzard) map[coord]bool {
	// map: coord -> is the space blocked or not

	out := map[coord]bool{}

	for wall := range walls {
		out[wall] = true
	}

	for _, blizzard := range blizzards {
		out[blizzard.loc] = true
	}

	return out
}
