package main

import (
	"fmt"
	"sort"

	"../../aoc"
)

// https://adventofcode.com/2020/day/10
// Daisy chaining power adapters

func main() {
	nums := aoc.ReadFileNumbers("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(nums))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(nums))
	fmt.Println(sw.Elapsed())
}

func step1(nums aoc.Int64Slice) int {
	sort.Sort(nums)
	ones := 0
	threes := 1
	prev := int64(0)
	for _, n := range nums {
		diff := n - prev
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
		prev = n
	}
	return ones * threes
}

const outlet = int64(0)

func step2(nums aoc.Int64Slice) int64 {
	sort.Sort(nums)
	pre := make(map[int64][]int64)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n <= int64(3) {
			pre[n] = append(pre[n], outlet)
		}
		for j := i - 1; j >= 0 && nums[i]-nums[j] <= int64(3) && nums[i]-nums[j] >= int64(0); j-- {
			pre[n] = append(pre[n], nums[j])
		}
	}
	paths := make(map[int64]int64)
	paths[outlet] = int64(1)
	for _, n := range nums {
		for _, p := range pre[n] {
			paths[n] += paths[p]
		}
	}

	return paths[nums[len(nums)-1]]
}
