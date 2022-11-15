package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(f)

	tiles := map[vector]int{}
	for scanner.Scan() {
		loc := vector{0, 0}
		// add a buffer
		line := scanner.Text()
		lineWithBuffer := line + "  " // just a stupid hack
		for i := 0; i < len(line); {
			if lineWithBuffer[i:i+2] == "se" {
				loc = loc.Add(SE)
				i += 2
			} else if lineWithBuffer[i:i+2] == "ne" {
				loc = loc.Add(NE)
				i += 2
			} else if lineWithBuffer[i:i+2] == "sw" {
				loc = loc.Add(SW)
				i += 2
			} else if lineWithBuffer[i:i+2] == "nw" {
				loc = loc.Add(NW)
				i += 2
			} else if lineWithBuffer[i:i+1] == "e" {
				loc = loc.Add(E)
				i += 1
			} else if lineWithBuffer[i:i+1] == "w" {
				loc = loc.Add(W)
				i += 1
			}
		}

		tiles[loc.round()] = 1 - tiles[loc.round()]
	}

	flipped := 0
	for _, value := range tiles {
		flipped += value
	}

	fmt.Printf("Part One: In total have %v flipped tiles\n", flipped)

	// Part Two
	for day := 1; day <= 100; day++ {
		toFlip := map[vector]bool{}
		// annoyingly we have to check not only all visited tiles but all neighbors of visited tiles
		// to expand at the boundary

		// some redundancy but its ok, go is fast
		// though this solution can be sped up by approximately 7x if
		// really needed by keeping track of where the boundary is,
		// i think? though clearly there is some cost to tracking the boundary
		// the exact cost is not obvious

		for tile := range tiles {
			toFlip[tile.round()] = tile.checkTile(tiles)
			toFlip[tile.Add(E).round()] = tile.Add(E).checkTile(tiles)
			toFlip[tile.Add(W).round()] = tile.Add(W).checkTile(tiles)
			toFlip[tile.Add(NE).round()] = tile.Add(NE).checkTile(tiles)
			toFlip[tile.Add(SE).round()] = tile.Add(SE).checkTile(tiles)
			toFlip[tile.Add(NW).round()] = tile.Add(NW).checkTile(tiles)
			toFlip[tile.Add(SW).round()] = tile.Add(SW).checkTile(tiles)
		}

		for tile, flip := range toFlip {
			if flip {
				tiles[tile] = 1 - tiles[tile]
			}
		}

		flipped = 0
		for _, value := range tiles {
			flipped += value
		}

	}

	flipped = 0
	for _, value := range tiles {
		flipped += value
	}

	fmt.Printf("Part Two: In total have %v flipped tiles\n", flipped)

}

type vector struct {
	x float64
	y float64
}

func (v vector) Add(other vector) vector {
	return vector{v.x + other.x, v.y + other.y}
}

func (v vector) round() vector {
	// To avoid issues of using floating points as keys?
	return vector{math.Round(v.x*1000) / 1000, math.Round(v.y*1000) / 1000}
}

func (v vector) sumAdjacent(t map[vector]int) int {
	out := 0
	out += t[v.Add(E).round()]
	out += t[v.Add(W).round()]
	out += t[v.Add(NE).round()]
	out += t[v.Add(SE).round()]
	out += t[v.Add(NW).round()]
	out += t[v.Add(SW).round()]
	return out
}

func (v vector) checkTile(tiles map[vector]int) bool {
	color := tiles[v.round()]
	if color == 1 && (v.sumAdjacent(tiles) == 0 || v.sumAdjacent(tiles) > 2) {
		return true
	} else if color == 0 && v.sumAdjacent(tiles) == 2 {
		return true
	} else {
		return false
	}
}

// Define primitive translations
var c60 float64 = math.Cos(math.Pi / 3)
var s60 float64 = math.Sin(math.Pi / 3)
var E vector = vector{1, 0}
var W vector = vector{-1, 0}
var NE vector = vector{c60, s60}
var NW vector = vector{-c60, s60}
var SE vector = vector{c60, -s60}
var SW vector = vector{-c60, -s60}
