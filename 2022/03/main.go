package main

import (
	"fmt"
	"strconv"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/3

const debug = false

func main() {
	lines := lib.ReadFileLines("2022/03/input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

func parseLines(lines []string) [][]int64 {
	elves := [][]int64{}
	elf := []int64{}

	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, elf)
			elf = []int64{}
		} else {
			n, _ := strconv.ParseInt(line, 10, 64)
			elf = append(elf, n)
		}
	}
	return elves
}

func score(r rune) int64 {
	if r >= 'a' && r <= 'z' {
		return int64(r - 'a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int64(r - 'A') + 27
	}
	if debug {
		fmt.Printf("Invalid score: '%s'\n", string(r))
	}
	return 0
}

func step1(lines []string) int64 {
	sum := int64(0)
	for _, line := range lines {
		half := len(line)/2
		first := line[0:half]
		second := line[half:]
		if (debug && len(first) != len(second)) {
			fmt.Printf("Unequal compartments: '%s' '%s'\n", first, second)
		}
		chars1 := lib.FrequencyCount(first)
		chars2 := lib.FrequencyCount(second)
		for key, _ := range chars1 {
			if _, found := chars2[key]; found {
				// Add the score of any item found in both halves
				sum += score(key)
			}
		}
	}
	return sum
}

func step2(lines []string) int64 {
	sum := int64(0)
	for i := 0; i < len(lines); i += 3 {
		first := lib.FrequencyCount(lines[i])
		second := lib.FrequencyCount(lines[i + 1])
		third := lib.FrequencyCount(lines[i + 2])
		for key, _ := range first {
			if _, found2 := second[key]; found2 {
				if _, found3 := third[key]; found3 {
					// Add the score of any item found in all 3 packs
					sum += score(key)
				}
			}
		}
	}
	return sum
}
