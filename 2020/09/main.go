package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/9
// TODO: Find numbers that sum to other numbers

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

func step1(nums []int64, depth int) int64 {
	for i := 25; i < len(nums); i++ {
		first, second := findSum2(nums[i-25:i], nums[i])
		if first == second {
			return nums[i]
		}
	}
	return -1
}

// Step 2: How many bags musta shiny gold bag contain?
func step2(nums []int64, target int64) int64 {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				min, max := minAndMax(nums[i : j+1])
				return min + max
			} else if sum > target {
				break
			}
		}
	}
	return int64(-1)
}

func minAndMax(nums []int64) (int64, int64) {
	min := int64(9999999999999)
	max := int64(-99999999999999)
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

const notFound = -1

func findSum2(nums []int64, target int64) (int64, int64) {
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			if a+b == target {
				return a, b
			}
		}
	}
	return notFound, notFound
}
