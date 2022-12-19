package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	blueprints := []blueprint{}
	for _, line := range strings.Split(string(contents), "\n") {
		costsString := strings.TrimSpace(strings.Split(line, ":")[1])
		var oreOreCost, clayOreCost, obsidianClayCost, obsidianOreCost, geodeObsidianCost, geodeOreCost int
		fmt.Sscanf(costsString,
			"Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&oreOreCost, &clayOreCost, &obsidianOreCost, &obsidianClayCost, &geodeOreCost, &geodeObsidianCost,
		)
		blueprints = append(blueprints, blueprint{
			oreConstruction:      resources{ore: oreOreCost},
			clayConstruction:     resources{ore: clayOreCost},
			obsidianConstruction: resources{ore: obsidianOreCost, clay: obsidianClayCost},
			geodeConstruction:    resources{ore: geodeOreCost, obsidian: geodeObsidianCost},
		})
	}

	now := time.Now()
	blueprintQuality := []int{}
	for _, b := range blueprints {
		blueprintQuality = append(blueprintQuality, geodeCapacity(b, 24))
	}

	var qualitySum int
	for i, v := range blueprintQuality {
		qualitySum += (i + 1) * v
	}
	fmt.Printf("Part One: Sum of weighted blueprint quality is %d\n", qualitySum)
	fmt.Println(time.Since(now))

	// Part Two
	now = time.Now()
	blueprintQuality = []int{}
	for _, b := range blueprints[:3] {
		blueprintQuality = append(blueprintQuality, geodeCapacity(b, 32))
	}
	var qualityProd int = 1
	for _, v := range blueprintQuality {
		qualityProd *= v
	}
	fmt.Printf("Part Two: Product of geode capacity of first three blueprints in 32 minutes is %d\n", qualityProd)
	fmt.Println(time.Since(now))

}

type resources struct {
	ore, clay, obsidian, geodes int // how much of each something we have. could be ore, could be robots... lol
}

func (r resources) harvest(amount resources) resources {
	r.ore += amount.ore
	r.clay += amount.clay
	r.obsidian += amount.obsidian
	r.geodes += amount.geodes
	return r
}

func (r resources) spend(amount resources) resources {
	// I check that the resources exist before calling the function
	r.ore -= amount.ore
	r.clay -= amount.clay
	r.obsidian -= amount.obsidian
	// we don't spend geodes
	return r
}

type blueprint struct {
	oreConstruction, clayConstruction, obsidianConstruction, geodeConstruction resources // how much each one costs to build
}

func (r resources) String() string {
	return fmt.Sprintf("%d ore, %d clay, %d obsidian, %d geodes", r.ore, r.clay, r.obsidian, r.geodes)
}

func (b blueprint) String() string {
	return fmt.Sprintf("Ore Robot: %v, Clay Robot: %v, Obsidian Robot: %v, Geode Robot: %v", b.oreConstruction, b.clayConstruction, b.obsidianConstruction, b.geodeConstruction)
}

type state struct {
	robots    resources
	backpack  resources
	cooldowns resources // this one seemed to be the key.
}

func (s state) String() string {
	return fmt.Sprintf("Robots: %v, Backpack: %v\n", s.robots, s.backpack)
}

func geodeCapacity(b blueprint, nRounds int) int {
	// how many geodes the blueprint can extract in n rounds
	// breadth first search maybe...
	current := []state{{resources{ore: 1}, resources{}, resources{}}}

	// i'm pretty sure that after we have built any of the resources,
	// we never need to consider a state which doesn't have that resource
	// in the tree anymore
	// well this turned out to not be true for any resource other than
	// geode
	have := resources{ore: 1}

	for depth := 0; depth < nRounds; {

		// some kind of breadth first search over possible actions at the current step
		// [nothing, build ore, build clay, build obsidian, build geode]
		next := []state{}
		for _, c := range current {
			harvested := c.backpack.harvest(c.robots)
			if c.backpack.canBuild(b.geodeConstruction) {
				// if we can build a geode, i'm almost positive this is the
				// optimal choice at any point in time since we don't get points
				// for anything else
				next = append(next, state{robots: c.robots.harvest(resources{geodes: 1}), backpack: harvested.spend(b.geodeConstruction), cooldowns: resources{}})
				have.geodes = max(have.geodes, c.robots.geodes+1)
			}
			if keep(c.robots, have, b) {
				cd := resources{}
				if c.backpack.canBuild(b.clayConstruction) {
					cd.clay = 1
				}
				if c.backpack.canBuild(b.obsidianConstruction) {
					cd.ore = 1
				}
				if c.backpack.canBuild(b.oreConstruction) {
					cd.ore = 1
				}
				next = append(next, state{robots: c.robots, backpack: harvested, cooldowns: cd})
			}
			if c.backpack.canBuild(b.clayConstruction) && keep(c.robots.harvest(resources{clay: 1}), have, b) && c.cooldowns.clay != 1 {
				next = append(next, state{robots: c.robots.harvest(resources{clay: 1}), backpack: harvested.spend(b.clayConstruction), cooldowns: resources{}})
			}
			if c.backpack.canBuild(b.obsidianConstruction) && keep(c.robots.harvest(resources{obsidian: 1}), have, b) && c.cooldowns.obsidian != 1 {
				next = append(next, state{robots: c.robots.harvest(resources{obsidian: 1}), backpack: harvested.spend(b.obsidianConstruction), cooldowns: resources{}})

			}
			if c.backpack.canBuild(b.oreConstruction) && keep(c.robots.harvest(resources{ore: 1}), have, b) && c.cooldowns.ore != 1 {
				next = append(next, state{robots: c.robots.harvest(resources{ore: 1}), backpack: harvested.spend(b.oreConstruction), cooldowns: resources{}})
			}
		}

		current = next
		// sorting seems to help performance, since it keeps the queue as small as possible
		sort.Slice(current, func(i, j int) bool {
			return current[i].backpack.canBuild(b.geodeConstruction) && !current[j].backpack.canBuild(b.geodeConstruction)
		})
		depth++
	}

	sort.Slice(current, func(i, j int) bool {
		return current[i].backpack.geodes > current[j].backpack.geodes
	})
	return current[0].backpack.geodes
}

func (backpack resources) canBuild(cost resources) bool {
	return backpack.ore >= cost.ore && backpack.clay >= cost.clay && backpack.obsidian >= cost.obsidian
}

func keep(r, have resources, b blueprint) bool {
	// these -5 are just a totally random heuristic, to try to find some pruning condition that works
	maxClayCost := b.obsidianConstruction.clay
	maxObsidianCost := b.geodeConstruction.obsidian
	maxOreCost := max(max(max(b.geodeConstruction.ore, b.obsidianConstruction.ore), b.clayConstruction.ore), b.oreConstruction.ore)
	return r.geodes >= have.geodes-1 && r.clay <= maxClayCost && r.obsidian <= maxObsidianCost && r.ore <= maxOreCost
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
