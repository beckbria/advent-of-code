package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/6

const debug = false

func main() {
	input := lib.ReadFileLines("2022/06/input.txt")[0]
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(input))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(input))
	fmt.Println(sw.Elapsed())
}

func allUnique[T comparable](s []T) bool {
	seen := make(map[T]bool)
	for _, k := range s {
		if _, found := seen[k]; found {
			return false
		}
		seen[k] = true
	}
	return true
}

func findUniqueSubstringEndIndex(input string, length int) int64 {
	prog := []rune(input)
	buffer := make([]rune, length)
	copy(buffer, prog[:length])
	replaceIterator := length - 1
	for i := replaceIterator; i < len(input); i++ {
		buffer[replaceIterator] = prog[i]
		replaceIterator = (replaceIterator + 1) % length
		if allUnique(buffer) {
			return int64(i + 1)
		}
	}
	return -1
}

func step1(input string) int64 {
	return findUniqueSubstringEndIndex(input, 4)
}

func step2(input string) int64 {
	return findUniqueSubstringEndIndex(input, 14)
}
