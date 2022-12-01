package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/1

func main() {
	lines := lib.ReadFileLines("2022/01/input.txt")
	input := parseLines(lines)
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(input))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(input))
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

func step1(elves [][]int64) int64 {
	max := int64(0)
	for _, elf := range elves {
		sum := int64(0)
		for _, n := range elf {
			sum = sum + n
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func step2(elves [][]int64) int64 {
	sums := lib.Int64Slice{}
	for _, elf := range elves {
		sum := int64(0)
		for _, n := range elf {
			sum = sum + n
		}
		sums = append(sums, sum)
	}
	sort.Sort(sums)
	n := len(sums)
	total := sums[n-1] + sums[n-2] + sums[n-3]

	return total
}
