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

	tiles := []tile{}
	for _, cell := range strings.Split(string(contents), "\n\n") {
		parsed, err := parse(cell)
		if err != nil {
			log.Fatal(err)
		}
		tiles = append(tiles, parsed)
	}

	// I really hope border matchings are unique...
	// If border matchings are unique, then a shortcut to part 1
	// can be to search for the four tiles whose edges only have two
	// other matches among all other tiles

	corners := []tile{}

	for _, thisTile := range tiles {
		var matchedEdges int

		for _, otherTile := range tiles {
			if thisTile.id == otherTile.id {
				continue
			}

			thisEdgesOriginal := thisTile.getAllEdgePermutations()[:4]
			otherEdgePossibilities := otherTile.getAllEdgePermutations()

			for _, thisEdge := range thisEdgesOriginal {
				for _, otherEdge := range otherEdgePossibilities {
					if equals(thisEdge, otherEdge) {
						matchedEdges++
						break
					}
				}
			}

		}
		if matchedEdges == 2 {
			corners = append(corners, thisTile)
		}
	}

	res := 1
	for _, c := range corners {
		res *= c.id
	}

	fmt.Printf("Part One: Product of all corner tile ids is %d\n", res)

}

// Actually we don't need need the whole layout, we only need the tile edges
// but it's nice to have for.... some reason
type tile struct {
	id     int
	layout [][]int
}

func parse(s string) (tile, error) {
	lines := strings.Split(s, "\n")
	header := lines[0]
	body := lines[1:]
	if len(lines) != 11 {
		return tile{}, fmt.Errorf("string contained %d lines rather than expected 11", len(lines))
	}
	if len(body[0]) != 10 {
		return tile{}, fmt.Errorf("line contained %d chars rather than expected 10", len(lines[0]))
	}

	var id int
	fmt.Sscanf(header, "Tile %d:", &id)

	layout := [][]int{}
	for _, line := range body {
		lr := []int{}
		for _, char := range line {
			if string(char) == "." {
				lr = append(lr, 0)
			} else {
				lr = append(lr, 1)
			}
		}
		layout = append(layout, lr)
	}

	return tile{id, layout}, nil
}

func (t tile) render() {
	s := ""
	for _, row := range t.layout {
		for _, v := range row {
			if v == 0 {
				s += "."
			} else {
				s += "#"
			}
		}
		s += "\n"
	}
	fmt.Println(strings.TrimSpace(s))
}

func (t tile) getAllEdgePermutations() [][]int {

	out := [][]int{}

	left := []int{}
	right := []int{}

	for _, row := range t.layout {
		for c, v := range row {
			if c == 0 {
				left = append(left, v)
			} else if c == 9 {
				right = append(right, v)
			}
		}
	}
	out = append(out, t.layout[0])
	out = append(out, left)
	out = append(out, right)
	out = append(out, t.layout[9])
	out = append(out, reverseArray(t.layout[0]))
	out = append(out, reverseArray(left))
	out = append(out, reverseArray(right))
	out = append(out, reverseArray(t.layout[9]))

	return out

}

func reverseArray(arr []int) []int {
	new := []int{}
	for i := range arr {
		new = append(new, arr[len(arr)-i-1])
	}
	return new
}

func equals(this []int, other []int) bool {
	// only being used for length 10 arrays! so don't need to check length
	eq := true
	for i := range this {
		if this[i] != other[i] {
			eq = false
		}
	}
	return eq
}
