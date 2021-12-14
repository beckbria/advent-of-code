package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/19
// RegEx engine

func main() {
	lines := lib.ReadFileLines("input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	rules, values := parseInput(lines)
	fmt.Println(step1(rules, values))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(rules, values))
	fmt.Println(sw.Elapsed())
}

type rule struct {
	patterns [][]int
	literal  string
}

const debug = false

func parseInput(lines []string) (map[int]*rule, []string) {
	rules := make(map[int]*rule)
	i := 0
	for ; len(lines[i]) > 0; i++ {
		r := rule{}
		tokens := strings.Split(lines[i], ": ")
		id, _ := strconv.Atoi(tokens[0])
		if tokens[1][0] == '"' {
			r.literal = tokens[1][1:2]
		} else {
			for _, p := range strings.Split(tokens[1], " | ") {
				var nums []int
				for _, n := range strings.Split(p, " ") {
					val, _ := strconv.Atoi(n)
					nums = append(nums, val)
				}
				r.patterns = append(r.patterns, nums)
			}
		}
		rules[id] = &r
	}
	return rules, lines[i+1:]
}

func compile(rules map[int]*rule, ruleID int) string {
	compiled := make(map[int]string)
	compileImpl(rules, compiled, ruleID)
	return compiled[ruleID]
}

func compileImpl(rules map[int]*rule, compiled map[int]string, ruleID int) string {
	if pat, found := compiled[ruleID]; found {
		return pat
	}

	r := rules[ruleID]
	if debug {
		fmt.Printf("Compiling rule %d [", ruleID)
		fmt.Print(r.patterns)
		fmt.Printf(" / %q]\n", r.literal)
	}
	if len(r.literal) > 0 {
		compiled[ruleID] = r.literal
		return r.literal
	}

	var options []string
	recurse := false
	for _, p := range r.patterns {
		if debug {
			fmt.Print("Pattern: ")
			fmt.Print(p)
			fmt.Println("")
		}
		for _, n := range p {
			if n == ruleID {
				continue
			}
			compileImpl(rules, compiled, n)
		}
		pattern := ""
		for _, n := range p {
			if n == ruleID {
				// Recursive patterns are fun.
				recurse = true
				pattern += "%s"
			}
			pattern += compiled[n]
		}
		options = append(options, pattern)
	}

	pat := "(" + strings.Join(options, "|") + ")"
	if debug {
		fmt.Printf("%d: %s\n", ruleID, pat)
	}
	if recurse {
		recursivePatterns := []string{pat}
		last := pat
		// Max depth = 8
		for d := 0; d < 8; d++ {
			last = fmt.Sprintf(last, pat)
			recursivePatterns = append(recursivePatterns, last)
		}
		// Remove the placeholder
		for i, s := range recursivePatterns {
			recursivePatterns[i] = strings.ReplaceAll(s, "%s", "")
		}
		// Join them together
		pat = "(" + strings.Join(recursivePatterns, "|") + ")"
		if debug {
			fmt.Printf("%d: %s\n", ruleID, pat)
		}
	}
	compiled[ruleID] = pat
	return pat
}

func step1(rules map[int]*rule, lines []string) int {
	pattern := compile(rules, 0)
	re, err := regexp.Compile("^" + pattern + "$")
	lib.Check(err)
	count := 0
	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}
	return count
}

func step2(rules map[int]*rule, lines []string) int {
	// Alter the input for step 2
	rules[8].patterns = [][]int{[]int{42}, []int{42, 8}}
	rules[11].patterns = [][]int{[]int{42, 31}, []int{42, 11, 31}}
	pattern := compile(rules, 0)
	re, err := regexp.Compile("^" + pattern + "$")
	lib.Check(err)
	count := 0
	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}
	return count
}
