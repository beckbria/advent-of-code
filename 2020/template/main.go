package main

import (
	"fmt"
	"regexp"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// TODO: Description

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

var (
	inputRegex = regexp.MustCompile(`^([a-z ]+) bags contain ([a-z0-9, ]+)\.$`)
)

func step1(lines []string) int {
	for _, l := range lines {
		tokens := inputRegex.FindStringSubmatch(l)
		first := tokens[1]
	}
	return -1
}

// Step 2: How many bags musta shiny gold bag contain?
func step2(lines []string) int64 {
	return -1
}
