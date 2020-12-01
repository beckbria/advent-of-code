package main

import (
	"fmt"
	"strconv"

	"../../aoc"
)

// https://adventofcode.com/2020/day/1
// Find some numbers that sum to another number
// I think we've all seen this interview problem, but for N here, O(N^3) isn't a problem

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	nums := parseLines(lines)
	a, b := findSum2(nums, 2020)
	fmt.Println("Step 1:")
	fmt.Println(a * b)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	c, d, e := findSum3(nums, 2020)
	fmt.Println(c * d * e)

	fmt.Println(sw.Elapsed())
}

func parseLines(lines []string) []int {
	nums := make([]int, 0)
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}
	return nums
}

func findSum2(nums []int, target int) (int, int) {
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			if a+b == target {
				return a, b
			}
		}
	}
	return -1, -1
}

func findSum3(nums []int, target int) (int, int, int) {
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			for k := j + 1; k < len(nums); k++ {
				c := nums[k]
				if a+b+c == target {
					return a, b, c
				}
			}
		}
	}
	return -1, -1, -1
}
