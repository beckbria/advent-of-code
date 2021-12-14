package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/15
// Simluate a memory game.  Start with a prompt, then say either
// 0 (if the last number was new) or the number of turns since that
// number was last said.

var input = []int64{15, 12, 0, 14, 3, 1}

func main() {
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(memoryGame(input, 2020))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(memoryGame(input, 30000000))
	fmt.Println(sw.Elapsed())
}

// memoryGame simulates the memory game and returns the number stated on
// the specified turn
func memoryGame(input []int64, finalTurn int64) int64 {
	seen := make(map[int64]int64)
	for i, n := range input {
		seen[n] = int64(i)
	}
	last := seen[int64(len(seen)-1)]
	wasNew := true
	delta := int64(0)
	for i := len(seen); int64(i) < finalTurn; i++ {
		last = delta
		if wasNew {
			last = 0
		}

		oldValue, found := seen[last]
		wasNew = !found
		if found {
			delta = int64(i) - oldValue
		}
		seen[last] = int64(i)
	}
	return last
}
