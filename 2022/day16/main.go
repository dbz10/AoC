package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type node struct {
	name          string
	flowRate      int
	connections   []*node
	exclusionZone []string
}

func (n node) String() string {
	return fmt.Sprint(n.name)
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	nodes := parseConnections(string(contents))
	pairwiseDistances := enumerateDistances(nodes)

	for _, n := range nodes {
		nnn := map[string]struct{}{}
		for _, x := range n.connections {
			for _, y := range x.connections {
				nnn[y.name] = struct{}{}
			}
		}
		exclusionZone := []string{}
		for key := range nnn {
			exclusionZone = append(exclusionZone, key)
		}
		n.exclusionZone = exclusionZone
	}

	// This seems kind of nontrivial. I'm not sure what to do
	// other than DFS
	nRounds := 30
	candidateNodes := []node{}
	for _, v := range nodes {
		if v.flowRate > 0 {
			candidateNodes = append(candidateNodes, *v)
		}
	}

	t0 := time.Now()

	dfsMaxScoreMaybe, path := dfsSearch("AA", nRounds, candidateNodes, pairwiseDistances)
	fmt.Printf("Part One: %d pressure released\n", dfsMaxScoreMaybe)
	fmt.Println(path)
	fmt.Println("took time:", time.Since(t0))

	// Part Two, may god have mercy on my soul
	dfsMaxScoreMaybePartTwo, pathPartTwo := pairDfsSearch("AA", "AA", 0, 0, nRounds-4, candidateNodes, candidateNodes, pairwiseDistances)
	fmt.Printf("Part Two: %d pressure released\n", dfsMaxScoreMaybePartTwo)
	fmt.Println(pathPartTwo)
	fmt.Println("took time:", time.Since(t0))

}

type tail struct {
	name           string
	score          int
	succeedingPath []string
}

func dfsSearch(currentPosition string, remainingTurns int, candidateNodes []node, distances map[string]map[string]int) (int, []string) {

	if len(candidateNodes) == 0 {
		return 0, []string{}
	}

	candidateScores := []tail{}
	for i, next := range candidateNodes {
		timeToOpen := distances[currentPosition][next.name] + 1
		// if we cannot reach that node in time, nothing to consider
		if timeToOpen > remainingTurns {
			continue
		}
		// otherwise, consider moving to that node and continuing
		thisNodeScore := next.flowRate * (remainingTurns - timeToOpen)

		// I think something about go slicing was biting me, so I just create
		// a brand new slice
		nextCandidates := removeFromSlice(candidateNodes, i)

		remainingScore, remainingPath := dfsSearch(
			next.name,
			remainingTurns-timeToOpen,
			nextCandidates,
			distances,
		)

		candidateScores = append(candidateScores, tail{next.name, thisNodeScore + remainingScore, remainingPath})

	}

	if len(candidateScores) == 0 {
		return 0, []string{}
	}

	sort.Slice(candidateScores, func(i, j int) bool {
		return candidateScores[i].score > candidateScores[j].score
	})

	fullPath := append([]string{candidateScores[0].name}, candidateScores[0].succeedingPath...)
	// return the max found
	return candidateScores[0].score, fullPath
}

func pairDfsSearch(
	selfDestination, elephantDestination string,
	selfTimeToArrive, elephantTimeToArrive int,
	remainingTurns int,
	selfCandidates, elephantCandidates []node,
	distances map[string]map[string]int,
) (int, []string) {
	// Basically my solution consists of DFS over interleaves of
	// mine and the elephant's path. If I was going to try to optimize
	// this further, I would include some simple path pruning measures.
	// One idea I had is to maintain different lists of next node candidates
	// for me and the elephant. When the elephant selects destination node X,
	// I would remove not only node X from both candidate lists, but
	// also remove the vicinity of X from my own candidate list, since
	// logically it cannot be optimal for me and the elephant to be
	// opening nodes in the same area. Anyways...

	// Ok I did end up implementing this with a radius 2 exclusion zone
	// and it reduced the time from ~3 minutes down to but
	// I'm not sure how general it is. With very highly connected
	// graph it would almost certainly fail.
	// And somehow a radius 3 exclusion zone becomes slower again.
	// To get further performance, probably requires a slightly
	// different approach?

	// For further optimizations, I would keep track of the maximum
	// score found so far and prune based on that.

	// If only one person is free at the time, then it's a single
	// person depth first search as before.
	// If both I and the elephant are free to explore, I will go first

	// Since I will fast forward to the earlier of whichever person
	// reaches their node in a given round, one of the cooldowns will
	// always be zero. Both cooldowns will be zero on the first turn.

	// I think by some simple pruning criteria I could have
	// drastically reduced the search space but it finishes in ~3 minutes
	// and I'm already behind so I must go on!

	if max(len(selfCandidates), len(elephantCandidates)) == 0 {
		return 0, []string{}
	}

	candidateScores := []tail{}

	if selfTimeToArrive == 0 {
		// Then I have arrived, I am current at selfDestination
		// and I am going to check for the next node

		for i, next := range selfCandidates {
			timeToOpen := distances[selfDestination][next.name] + 1
			if timeToOpen > remainingTurns {
				continue
			}
			thisNodeScore := next.flowRate * (remainingTurns - timeToOpen)
			nextCandidates := removeFromSlice(selfCandidates, i)
			nextElephantCandidates := removeExclusionZone(elephantCandidates, next)
			selfTimeToArrive = timeToOpen
			timeUntilNextEvent := min(selfTimeToArrive, elephantTimeToArrive)
			remainingScore, remainingPath := pairDfsSearch(
				next.name,
				elephantDestination,
				selfTimeToArrive-timeUntilNextEvent,
				elephantTimeToArrive-timeUntilNextEvent,
				remainingTurns-timeUntilNextEvent,
				nextCandidates,
				nextElephantCandidates,
				distances,
			)
			candidateScores = append(candidateScores, tail{next.name, thisNodeScore + remainingScore, remainingPath})
		}
	} else if elephantTimeToArrive == 0 {
		// Then the elephant has arrived at elephantDestination
		// and he will go to check for the next node

		for i, next := range elephantCandidates {
			timeToOpen := distances[elephantDestination][next.name] + 1
			if timeToOpen > remainingTurns {
				continue
			}
			thisNodeScore := next.flowRate * (remainingTurns - timeToOpen)
			nextCandidates := removeFromSlice(elephantCandidates, i)
			nextSelfCandidates := removeExclusionZone(selfCandidates, next)
			elephantTimeToArrive = timeToOpen
			timeUntilNextEvent := min(selfTimeToArrive, elephantTimeToArrive)
			remainingScore, remainingPath := pairDfsSearch(
				selfDestination,
				next.name,
				selfTimeToArrive-timeUntilNextEvent,
				elephantTimeToArrive-timeUntilNextEvent,
				remainingTurns-timeUntilNextEvent,
				nextSelfCandidates,
				nextCandidates,
				distances,
			)
			candidateScores = append(candidateScores, tail{next.name, thisNodeScore + remainingScore, remainingPath})
		}
	}

	if len(candidateScores) == 0 {
		return 0, []string{}
	}

	sort.Slice(candidateScores, func(i, j int) bool {
		return candidateScores[i].score > candidateScores[j].score
	})

	fullPath := append([]string{candidateScores[0].name}, candidateScores[0].succeedingPath...)
	// return the max found
	return candidateScores[0].score, fullPath
}

func removeFromSlice(nn []node, i int) []node {
	out := []node{}
	for j, n := range nn {
		if j == i {
			continue
		}
		out = append(out, n)
	}
	return out
}

func removeExclusionZone(nn []node, n node) []node {
	// construct the set of next nearest neighbors

	out := []node{}
	for _, other := range nn {
		if !contains(n.exclusionZone, other.name) {
			out = append(out, other)
		}

	}
	return out
}

func contains(ss []string, s string) bool {
	for _, sp := range ss {
		if sp == s {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func enumerateDistances(nodes map[string]*node) map[string]map[string]int {
	var pairwiseDistances = map[string]map[string]int{}
	for k1, n1 := range nodes {
		pairwiseDistances[k1] = map[string]int{}
		for k2, n2 := range nodes {
			if k1 == k2 {
				continue
			}
			pairwiseDistances[k1][k2] = calculateDistance(*n1, *n2)
		}
	}
	return pairwiseDistances
}

func calculateDistance(n1, n2 node) int {
	// BFS
	visited := map[string]struct{}{
		n1.name: {},
	}
	queue := []*node{}
	queue = append(queue, n1.connections...)

	var intraLevelCounter, level int
	var levelSize int = len(queue)

	for len(queue) > 0 {
		if intraLevelCounter == levelSize {
			intraLevelCounter = 0
			level++
			levelSize = len(queue)
		}

		head := queue[0]
		if head.name == n2.name {
			level++
			return level
		}
		visited[head.name] = struct{}{}
		for _, connection := range head.connections {
			_, exists := visited[connection.name]
			if !exists {
				queue = append(queue, connection)
			}
		}
		queue = queue[1:]

		intraLevelCounter++
	}
	return level

}

func parseConnections(contents string) map[string]*node {
	lines := strings.Split(contents, "\n")

	re := regexp.MustCompile(`[A-Z, ]*$`)
	nodes := map[string]*node{}

	// As usual, have to go through two loops. 1 to create all the nodes, 2 to add the connections
	for _, line := range lines {
		var valve string
		var flowRate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &valve, &flowRate)
		nodes[valve] = &node{name: valve, flowRate: flowRate}
	}
	for _, line := range lines {
		var valve string
		var flowRate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d; ", &valve, &flowRate)
		connections := re.FindString(line)
		for _, other := range strings.Split(connections, ", ") {
			nodes[valve].connections = append(nodes[valve].connections, nodes[strings.TrimSpace(other)])
		}
	}
	return nodes
}
