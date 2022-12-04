package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/4

var (
	inputRegex = regexp.MustCompile(`^([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)$`)
)

type span struct {
	start, end int64
}

func (s1 *span) totallyContains(s2 *span) bool {
	return s2.start >= s1.start && s2.end <= s1.end
}

func (s *span) set(a, b int64) {
	if a > b {
		s.start = b
		s.end = a
	} else {
		s.start = a
		s.end = b
	}
}

type spanPair struct {
	first, second span
}

func (sp *spanPair) eitherTotallyContained() bool {
	return sp.first.totallyContains(&sp.second) || sp.second.totallyContains(&sp.first)
}

func (sp *spanPair) overlap() bool {
	return sp.first.start <= sp.second.end && sp.first.end >= sp.second.start
}

func parsePair(s string) *spanPair {
	sp := spanPair{}
	tokens := inputRegex.FindStringSubmatch(s)
	if len(tokens) != 5 {
		fmt.Printf("Bad input '%s'\n", s)
	}
	var start, end int64
	start, _ = strconv.ParseInt(tokens[1], 10, 64)
	end, _ = strconv.ParseInt(tokens[2], 10, 64)
	sp.first.set(start, end)
	start, _ = strconv.ParseInt(tokens[3], 10, 64)
	end, _ = strconv.ParseInt(tokens[4], 10, 64)
	sp.second.set(start, end)
	return &sp
}

func main() {
	lines := lib.ReadFileLines("2022/04/input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

func step1(lines []string) int64 {
	pairs := []*spanPair{}
	for _, line := range lines {
		pairs = append(pairs, parsePair(line))
	}

	totallyContained := int64(0)
	for _, p := range pairs {
		if p.eitherTotallyContained() {
			totallyContained++
		}
	}

	return totallyContained
}

func step2(lines []string) int64 {
	pairs := []*spanPair{}
	for _, line := range lines {
		pairs = append(pairs, parsePair(line))
	}

	overlap := int64(0)
	for _, p := range pairs {
		if p.overlap() {
			overlap++
		}
	}

	return overlap
}
