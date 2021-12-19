package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/#
// Reverse-engineer the format of a ticket

func main() {
	lines := lib.ReadFileLines("2020/16/input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	rules, yours, others := parseInput(lines)
	answer1, invalidIndexes := step1(rules, others)
	fmt.Println(answer1)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(rules, yours, others, invalidIndexes))
	fmt.Println(sw.Elapsed())
}

type ticket []int

type interval struct {
	min, max int
}

type rule struct {
	name string
	r    []interval
}

func (r *rule) valid(i int) bool {
	for _, ra := range r.r {
		if i >= ra.min && i <= ra.max {
			return true
		}
	}
	return false
}

func parseInput(lines []string) ([]rule, ticket, []ticket) {
	rules := []rule{}
	i := 0
	// Parse the rules
	for ; i < len(lines); i++ {
		l := lines[i]
		if len(l) == 0 {
			i += 2 // empty line, "your ticket:"
			break
		}
		r := rule{}
		tok := strings.Split(l, ": ")
		r.name = tok[0]
		for _, iv := range strings.Split(tok[1], " or ") {
			t := strings.Split(iv, "-")
			min, _ := strconv.Atoi(t[0])
			max, _ := strconv.Atoi(t[1])
			r.r = append(r.r, interval{min: min, max: max})
		}
		rules = append(rules, r)
	}

	yours := parseTicket(lines[i])
	i += 3 // empty line, "nearby tickets:"

	others := []ticket{}
	for ; i < len(lines); i++ {
		others = append(others, parseTicket(lines[i]))
	}

	return rules, yours, others
}

func parseTicket(s string) ticket {
	t := ticket{}
	for _, n := range strings.Split(s, ",") {
		i, _ := strconv.Atoi(n)
		t = append(t, i)
	}
	return t
}

func step1(rules []rule, others []ticket) (int, []int) {
	sum := 0
	invalidIdx := []int{}
	for i, t := range others {
		for _, n := range t {
			valid := false
			for _, r := range rules {
				valid = r.valid(n)
				if valid {
					break
				}
			}
			if !valid {
				sum += n
				invalidIdx = append(invalidIdx, i)
				break
			}
		}
	}
	return sum, invalidIdx
}

func step2(rules []rule, yours ticket, others []ticket, invalidIndexes []int) int {
	tickets := pruneInvalid(others, invalidIndexes)
	candidates := findCandidates(rules, tickets)
	product := int(1)
	for name, val := range candidates {
		if len(name) >= 9 && name[0:9] == "departure" {
			product *= yours[val]
		}
	}
	return product
}

// pruneInvalid removes the specified array indices from a slice
func pruneInvalid(orig []ticket, invalidIndexes []int) []ticket {
	sort.Ints(invalidIndexes)
	tickets := make([]ticket, len(orig))
	copy(tickets, orig)
	for i := len(invalidIndexes) - 1; i >= 0; i-- {
		split := invalidIndexes[i]
		tickets = append(tickets[0:split], tickets[split+1:]...)
	}
	return tickets
}

// buildCandidateMap returns a map of rule name to all indexes which would be valid for that rule
func buildCandidateMap(rules []rule, tickets []ticket) map[string]map[int]bool {
	c := make(map[string]map[int]bool)
	for _, r := range rules {
		for i := range tickets[0] {
			valid := true
			for _, t := range tickets {
				valid = r.valid(t[i])
				if !valid {
					break
				}
			}
			if valid {
				if c[r.name] == nil {
					c[r.name] = make(map[int]bool)
				}
				c[r.name][i] = true
			}
		}
	}
	return c
}

// findCandidates returns a map of rule name to the index corresponding to that rule
func findCandidates(rules []rule, tickets []ticket) map[string]int {
	allCandidates := buildCandidateMap(rules, tickets)
	finalValues := make(map[string]int)
	for len(allCandidates) > 0 {
		for name, values := range allCandidates {
			if len(values) == 1 {
				val := 0
				for k := range values {
					val = k
				}

				finalValues[name] = val
				// Remove that value from all other candidates
				for k := range allCandidates {
					delete(allCandidates[k], val)
				}
				delete(allCandidates, name)
				// Start over because we may have altered traversal of the map
				break
			}
		}
	}
	return finalValues
}
