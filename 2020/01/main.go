package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/1
// Find some numbers that sum to another number
// I think we've all seen this interview problem, but for N here, O(N^3) isn't a problem

const notFound = int64(-1)

func main() {
	nums := lib.ReadFileNumbers("2020/01/input.txt")
	sw := lib.NewStopwatch()
	_, a, b := lib.FindSum2(nums, int64(2020))
	fmt.Println("Step 1:")
	fmt.Println(a * b)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	c, d, e := findSum3(nums, int64(2020))
	fmt.Println(c * d * e)

	fmt.Println(sw.Elapsed())
}

func findSum3(nums []int64, target int64) (int64, int64, int64) {
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
	return notFound, notFound, notFound
}
