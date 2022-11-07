package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// In terms of data structure, it seems like the natural data structure is a bipartite graph.
	// But practically it may be easier to work with lists...

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`^([^\()]*) \(contains ([^\)]*)\)$`)
	allAllIngredients := []string{}
	ingredientAllergenMapping := map[string][][]string{}
	solved := map[string]string{}

	nonAllergens := map[string]bool{} // bit confusingly named, but it will make sense later

	for scanner.Scan() {
		contents := re.FindStringSubmatch(scanner.Text())
		ingredients := strings.Split(contents[1], " ")
		for _, i := range ingredients {
			nonAllergens[i] = true
			allAllIngredients = append(allAllIngredients, i)
		}
		allergens := strings.Split(contents[2], ", ")

		for _, a := range allergens {
			ingredientAllergenMapping[a] = append(ingredientAllergenMapping[a], ingredients)
		}
	}

	// Now, iteratively prune the possibilities, trading off between set intersections
	// over possibilities, and updating known values.

	for len(ingredientAllergenMapping) > 0 {
		for key, value := range ingredientAllergenMapping {
			mem := threeWayIntersect(value[0], value[0], solved)
			for _, n := range value[1:] {
				mem = threeWayIntersect(mem, n, solved)
			}

			if len(mem) == 1 {
				solved[mem[0]] = key
				delete(ingredientAllergenMapping, key)
			}
		}
	}

	for ingredient, _ := range nonAllergens {
		_, exists := solved[ingredient]
		if exists {
			delete(nonAllergens, ingredient)
		}
	}

	var nonAllergenAppearances int
	for _, ingredient := range allAllIngredients {
		_, exists := nonAllergens[ingredient]
		if exists {
			nonAllergenAppearances++
		}
	}

	fmt.Printf("Part One: Found %d non allergen appearances\n", nonAllergenAppearances)

	// Part two requires jumping through 1.5 hoops but its not so bad.
	// Invert the ingredient - allergen mapping, load out all allergens, sort them,
	// then retrieve the keys in sorted order.
	allergenIngredientMapping := map[string]string{}
	allergens := []string{}
	for ingredient, allergen := range solved {
		allergenIngredientMapping[allergen] = ingredient
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)
	ingredientsSortedByAllergen := []string{}
	for _, a := range allergens {
		ingredientsSortedByAllergen = append(ingredientsSortedByAllergen, allergenIngredientMapping[a])
	}

	fmt.Println("Part Two: Canonical something or other")
	fmt.Println(strings.Join(ingredientsSortedByAllergen, ","))

}

type void struct{}

func threeWayIntersect(this []string, other []string, solved map[string]string) []string {
	var nothing void
	otherHashMap := map[string]void{}

	for _, v := range other {
		otherHashMap[v] = nothing
	}

	out := []string{}
	for _, v := range this {
		_, exists := otherHashMap[v]
		_, alreadyUsed := solved[v]
		if exists && (!alreadyUsed) {
			out = append(out, v)
		}
	}

	return out
}
