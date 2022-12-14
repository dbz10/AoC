package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const initTd = 100000000

type node struct {
	value            int
	validNeighbors   []point
	reverseNeighbors []point
}

func (n *node) addValidNeighbor(other point) {
	n.validNeighbors = append(n.validNeighbors, other)
}

type point struct {
	x, y int
}

var charValue = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"S": 1,
	"E": 26,
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	nodes := map[point]*node{}
	rows := strings.Split(string(contents), "\n")

	letterGrid := [][]string{}
	for _, row := range rows {
		letterRow := []string{}
		for _, char := range row {
			letterRow = append(letterRow, string(char))
		}
		letterGrid = append(letterGrid, letterRow)
	}

	xShape := len(letterGrid)
	yShape := len(letterGrid[0])

	var start, end point

	// Construct graph in two stages.
	// First, initialize all of the nodes
	for x, row := range letterGrid {
		for y, v := range row {
			nodes[point{x, y}] = &node{value: charValue[v]}
			if v == "S" {
				start = point{x, y}
			}
			if v == "E" {
				end = point{x, y}
			}
		}
	}

	// Now, add connections
	for loc := range nodes {
		allNeighbors := getNeighborIndices(loc.x, loc.y, xShape, yShape)
		for _, neighbor := range allNeighbors {
			if nodes[neighbor].value <= nodes[loc].value+1 {
				nodes[loc].addValidNeighbor(neighbor)
			}
		}
	}

	pathLength := dijkstra(nodes, start, end)
	fmt.Printf("Part One: %d steps are required to reach the high point\n", pathLength)

	// Part Two
	// I know this is highly inefficient but I just want to move on.

	startingPointCandidates := []point{}
	for point, node := range nodes {
		if node.value == 1 {
			startingPointCandidates = append(startingPointCandidates, point)
		}
	}

	optimizedPathLength := pathLength
	for _, candidate := range startingPointCandidates {
		distance := dijkstra(nodes, candidate, end)
		if distance < optimizedPathLength {
			optimizedPathLength = distance
		}
	}

	fmt.Printf("Part Two: The shortest path from base elevation to the optimal reception point takes %d steps\n", optimizedPathLength)

}

func getNeighborIndices(x, y, xShape, yShape int) []point {
	neighboringIndices := []point{}
	for _, dx := range []int{-1, 1} {
		if x+dx >= 0 && x+dx < xShape {
			neighboringIndices = append(neighboringIndices, point{x + dx, y})
		}
	}

	for _, dy := range []int{-1, 1} {
		if y+dy >= 0 && y+dy < yShape {
			neighboringIndices = append(neighboringIndices, point{x, y + dy})
		}
	}

	return neighboringIndices
}

func dijkstra(graph map[point]*node, start, end point) int {
	// Set up a set of all unvisited nodes
	unvisitedNodes := []point{}
	for key := range graph {
		unvisitedNodes = append(unvisitedNodes, key)
	}
	distanceToNodes := map[point]int{}
	// set initial distance super high, technically infinity but for a limited size graph, this is more than enough
	for key := range graph {
		distanceToNodes[key] = initTd
	}
	distanceToNodes[start] = 0
	sort.Slice(unvisitedNodes, func(i, j int) bool { return distanceToNodes[unvisitedNodes[i]] < distanceToNodes[unvisitedNodes[j]] })

	// ok it's not the most clear interface with jumping back and forth between
	// point <-> node but anyways, it works

	for contains(unvisitedNodes, end) && distanceToNodes[unvisitedNodes[0]] < initTd {
		current := unvisitedNodes[0]
		unvisitedNodes = unvisitedNodes[1:]
		for _, neighbor := range graph[current].validNeighbors {
			if contains(unvisitedNodes, neighbor) && distanceToNodes[neighbor] > distanceToNodes[current]+1 {
				distanceToNodes[neighbor] = distanceToNodes[current] + 1
			}
		}
		sort.Slice(unvisitedNodes, func(i, j int) bool { return distanceToNodes[unvisitedNodes[i]] < distanceToNodes[unvisitedNodes[j]] })
		if len(unvisitedNodes) == 0 {
			break
		}
	}

	return distanceToNodes[end]
}

func contains(unvisitedNodes []point, p point) bool {
	for _, pp := range unvisitedNodes {
		if pp == p {
			return true
		}
	}
	return false
}
