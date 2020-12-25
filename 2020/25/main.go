package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/25
// RFID Encryption

var data = []int64{8458505, 16050997}

func main() {
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	_, _, key := step1(data)
	fmt.Println(key)
	fmt.Println(sw.Elapsed())
}

func step1(input []int64) (int64, int64, int64) {
	cardPublic := input[0]
	doorPublic := input[1]

	cardLoop, doorLoop := int64(-1), int64(-1)

	val := int64(1)
	for it := int64(1); cardLoop < 0 || doorLoop < 0; it++ {
		val = (val * 7) % 20201227
		if cardPublic == val {
			cardLoop = it
		}
		if doorPublic == val {
			doorLoop = it
		}
	}

	key := cardPublic
	for i := int64(1); i < doorLoop; i++ {
		key = (key * cardPublic) % 20201227
	}

	return cardLoop, doorLoop, key
}
