package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"../aoc"
)

const debug = false

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	f := readFormulas(input)
	fmt.Println(f.convert("ORE", "FUEL"))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(sw.Elapsed())
}

type formula struct {
	from map[string]int
	to   map[string]int
}

// To returns the number and item produced by the formula
func (f *formula) To() (string, int) {
	if len(f.to) != 1 {
		log.Fatalf("Invalid to length: %d\n", len(f.to))
	}
	for k, v := range f.to {
		return k, v
	}
	return "", -1
}

type formulas struct {
	f []*formula
	// lookup maps from the produced type to the source ingredients
	lookup map[string]*formula
}

var (
	inputRegEx = regexp.MustCompile("^([0-9]+ [A-Z]+(, [0-9]+ [A-Z]+)*) => ([0-9]+ [A-Z]+)$")
)

func readFormulas(input []string) formulas {
	f := formulas{f: []*formula{}, lookup: make(map[string]*formula)}
	for _, s := range input {
		tokens := inputRegEx.FindStringSubmatch(s)
		from := readComponents(tokens[1])
		to := readComponents(tokens[3])

		form := formula{from: from, to: to}
		f.f = append(f.f, &form)
		k, _ := form.To()
		f.lookup[k] = &form
	}
	return f
}

func readComponents(cl string) map[string]int {
	comps := make(map[string]int)

	for _, c := range strings.Split(cl, ", ") {
		pieces := strings.Split(c, " ")
		n, err := strconv.Atoi(pieces[0])
		aoc.Check(err)
		comps[pieces[1]] = n
	}

	return comps
}

func (f formulas) convert(from, to string) int64 {

	need := make(map[string]int64)
	need[to] = 1
	need[from] = 0
	// Loop until only the from element remains
	for len(need) > 1 {
		for k, v := range need {
			if k == from {
				continue
			}
			producesK := f.lookup[k]
			_, numK := producesK.To()

			quantity := int64(math.Ceil(float64(v) / float64(numK)))
			if debug {
				fmt.Printf("Producing %d %s from %d ", int64(numK)*quantity, k, quantity)
				fmt.Println(producesK.from)
			}

			// Add the new required ingredients
			for needK, needV := range producesK.from {
				need[needK] += quantity * int64(needV)
			}
			delete(need, k)
		}
	}

	return need[from]
}
