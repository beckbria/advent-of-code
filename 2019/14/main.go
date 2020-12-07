package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	f := readFormulas(input)
	fmt.Println(f.minimumRequired("ORE", "FUEL", 1))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(f.mostPossible("ORE", "FUEL"))
	fmt.Println(sw.Elapsed())
}

type formula struct {
	from map[string]int64
	to   map[string]int64
}

// To returns the number and item produced by the formula
func (f *formula) To() (string, int64) {
	if len(f.to) != 1 {
		log.Fatalf("Invalid to length: %d\n", len(f.to))
	}
	for k, v := range f.to {
		return k, v
	}
	return "", -1
}

type formulas struct {
	// lookup maps from the produced type to the source ingredients
	produces map[string]*formula
	requires map[string]*formula
}

var (
	inputRegEx = regexp.MustCompile(`^([0-9]+ [A-Z]+(, [0-9]+ [A-Z]+)*) => ([0-9]+ [A-Z]+)$`)
)

func readFormulas(input []string) formulas {
	f := formulas{produces: make(map[string]*formula), requires: make(map[string]*formula)}
	for _, s := range input {
		tokens := inputRegEx.FindStringSubmatch(s)
		from := readComponents(tokens[1])
		to := readComponents(tokens[3])

		form := formula{from: from, to: to}
		k, _ := form.To()
		f.produces[k] = &form
		for k := range form.from {
			f.requires[k] = &form
		}
	}
	return f
}

func readComponents(cl string) map[string]int64 {
	comps := make(map[string]int64)

	for _, c := range strings.Split(cl, ", ") {
		pieces := strings.Split(c, " ")
		n, err := strconv.ParseInt(pieces[0], 10, 64)
		aoc.Check(err)
		comps[pieces[1]] = n
	}

	return comps
}

func (f *formulas) minimumRequired(from, to string, quantity int64) int64 {
	need := make(map[string]int64)
	surplus := make(map[string]int64)
	need[to] = quantity
	need[from] = 0
	// Loop until only the from element remains
	for len(need) > 1 {
		for k, v := range need {
			if k == from {
				continue
			}
			producesK := f.produces[k]
			_, numK := producesK.To()

			if surp, found := surplus[k]; found {
				v -= surp
				delete(surplus, k)
			}

			quantity := int64(math.Ceil(float64(v) / float64(numK)))
			surplus[k] += numK*quantity - v

			// Add the new required ingredients
			for needK, needV := range producesK.from {
				need[needK] += quantity * needV
			}
			delete(need, k)
		}
	}

	return need[from]
}

func (f *formulas) mostPossible(from, to string) int64 {
	// Search
	fuel := int64(0)
	jump := []int64{
		1000000,
		100000,
		10000,
		1000,
		100,
		10,
		1}
	target := int64(1000000000000)
	for _, j := range jump {
		fuel = f.search(fuel, j, target)
	}
	return fuel
}

func (f *formulas) search(start, increment, target int64) int64 {
	for current := start; true; current += increment {
		cost := f.minimumRequired("ORE", "FUEL", current)
		if cost > target {
			// We've gone too far
			return current - increment
		}
	}
	return -1
}
