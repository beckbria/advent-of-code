package main

import (
	"fmt"
	"regexp"

	"github.com/beckbria/advent-of-code/2021/lib"
)

// https://adventofcode.com/2020/day/#
// TODO: Description

func main() {
	lines := lib.ReadFileNumbers("input.txt")
	sw := lib.NewStopwatch()
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

func step1(lines []int64) int64 {
	count := int64(0)
	prev := int64(9999999999)
	for _, n := range lines {
		if n > prev {
			count++
		}
		prev = n
	}
	return count
}

func step2(lines []int64) int64 {
	count := int64(0)
	prevSum := int64(9999999999)
	for i := range lines {
		newSum := lines[i] + lines[(i+1)%len(lines)] + lines[(i+2)%len(lines)]
		if newSum > prevSum {
			count++
		}
		prevSum = newSum
	}
	return count
}
