package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/21
// Determine which ingredients cause which allergeies

type set map[string]bool

type recipe struct {
	ingredients set
	allergens   set
}

func newRecipe() *recipe {
	var r recipe
	r.ingredients = make(set)
	r.allergens = make(set)
	return &r
}

func main() {
	lines := lib.ReadFileLines("input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	r := parseRecipes(lines)
	count, cause := step1(r)
	fmt.Println(count)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(cause))
	fmt.Println(sw.Elapsed())
}

var (
	inputRegex = regexp.MustCompile(`^([a-z ]+) \(contains ([a-z, ]+)\)$`)
)

func parseRecipes(lines []string) []*recipe {
	var r []*recipe
	for _, l := range lines {
		r = append(r, parseRecipe(l))
	}
	return r
}

func parseRecipe(line string) *recipe {
	r := newRecipe()
	tok := inputRegex.FindStringSubmatch(line)
	for _, i := range strings.Split(tok[1], " ") {
		r.ingredients[i] = true
	}
	for _, a := range strings.Split(tok[2], ", ") {
		r.allergens[a] = true
	}
	//fmt.Println(*r)
	return r
}

const debug = true

func step1(recipes []*recipe) (int64, map[string]string) {
	allergens := make(map[string]set)
	allIngredients := make(set)
	safeIngredients := make(set)
	// Record which allergens are produced by which ingredients.
	for _, r := range recipes {
		for a := range r.allergens {
			if _, found := allergens[a]; found {
				// Check all of the ingredients against the candidate list
				for i := range allergens[a] {
					if _, found2 := r.ingredients[i]; !found2 {
						// Reject this ingredient as it isn't in this recipe
						delete(allergens[a], i)
					}
				}
			} else {
				allergens[a] = make(set)
				for i := range r.ingredients {
					allergens[a][i] = true
				}
			}
			for i := range r.ingredients {
				allIngredients[i] = true
				safeIngredients[i] = true
			}
		}
	}

	// Determine which ingredients cause which allergens
	cause := make(map[string]string)
ALLERGENS:
	for len(allergens) > 0 {
		for a, is := range allergens {
			if len(is) < 1 {
				log.Fatalf("Impossible allergen: %s", a)
			} else if len(is) == 1 {
				i := ""
				for k := range is {
					i = k
				}
				cause[a] = i
				// Reject this cause in every other allergen
				for k := range allergens {
					delete(allergens[k], i)
				}
				delete(safeIngredients, i)
				delete(allergens, a)
				continue ALLERGENS
			}
		}
	}

	// Count the safe allergens
	count := int64(0)
	for _, r := range recipes {
		for i := range r.ingredients {
			if _, found := safeIngredients[i]; found {
				count++
			}
		}
	}
	return count, cause
}

// Takes map from allergen to effect
func step2(cause map[string]string) string {
	var allergens []string
	for k := range cause {
		allergens = append(allergens, k)
	}
	sort.Strings(allergens)
	var answer []string
	for _, a := range allergens {
		answer = append(answer, cause[a])
	}
	return strings.Join(answer, ",")
}
