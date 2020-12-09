package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/9
// Find numbers that sum to other numbers

func main() {
	nums := aoc.ReadFileNumbers("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	step1Answer := step1(nums, 25)
	fmt.Println(step1Answer)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(nums, step1Answer))
	fmt.Println(sw.Elapsed())
}

// Step 1: Find a number which is not the sum of any two of the n values preceding it
func step1(nums []int64, depth int) int64 {
	for i := depth; i < len(nums); i++ {
		found, _, _ := aoc.FindSum2(nums[i-depth:i], nums[i])
		if !found {
			return nums[i]
		}
	}
	return -1
}

// Step 2: Find a contiguous string of numbers which sum to a value.
// There are more efficient ways, but N=1000, so O(N^2) runs in microseconds
func step2(nums []int64, target int64) int64 {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				min, max := aoc.MinAndMax(nums[i : j+1])
				return min + max
			} else if sum > target {
				break
			}
		}
	}
	return int64(-1)
}
