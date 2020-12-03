package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/3
// Check if a skier will hit trees

const (
	tree = '#'
	open = '.'
)

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(findTrees(lines, 3, 1))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(findTrees(lines, 1, 1) * findTrees(lines, 3, 1) * findTrees(lines, 5, 1) * findTrees(lines, 7, 1) * findTrees(lines, 1, 2))
	fmt.Println(sw.Elapsed())
}

func findTrees(lines []string, xD, yD int) int {
	x := 0
	y := 0
	trees := 0
	width := len(lines[0])
	for y < len(lines) {
		if lines[y][x%width] == tree {
			trees++
		}
		x += xD
		y += yD
	}
	return trees
}
