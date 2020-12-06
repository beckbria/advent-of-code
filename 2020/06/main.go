package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/6
// Parse survey inputs

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

// step1 counts the number of distinct characters seem in each section.
// Sections are separated by blank lines.
func step1(lines []string) int {
	sum := 0
	seen := make(map[rune]bool)
	for _, l := range lines {
		if len(l) == 0 {
			sum += len(seen)
			seen = make(map[rune]bool)
		}
		for _, c := range l {
			seen[c] = true
		}
	}
	sum += len(seen)
	return sum
}

// step2 counts the number of characters found on every line of a section.
// Sections are separated by blank lines.
func step2(lines []string) int {
	sum := 0
	seen := make(map[rune]int)
	count := 0
	for _, l := range lines {
		if len(l) == 0 {
			for _, c := range seen {
				if c == count {
					sum++
				}
			}
			count = 0
			seen = make(map[rune]int)
		} else {
			for _, c := range l {
				seen[c] = seen[c] + 1
			}
			count++
		}
	}
	for _, c := range seen {
		if c == count {
			sum++
		}
	}
	return sum
}
