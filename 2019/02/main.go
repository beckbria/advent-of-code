package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

// https://adventofcode.com/2019/day/2
// Write a computer for a made-up language.  Run a program.
// Part 2: Find a set of inputs to match the target value

const add = int64(1)
const multiply = int64(2)
const terminate = int64(99)
const debug = false

func main() {
	program := intcode.ReadIntCode("input.txt")
	sw := lib.NewStopwatch()

	// Step 1 but with an IntCode computer
	step1 := make([]int64, len(program))
	// Corrections from instruction
	copy(step1, program)
	step1[1] = int64(12)
	step1[2] = int64(2)
	c := intcode.NewComputer(step1)
	c.Run()
	fmt.Println("Step 1:")
	fmt.Println(c.Memory[0])
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	a, b := findTarget(program, int64(19690720))
	fmt.Printf("%d %d -> %d\n", a, b, (100*a)+b)
	fmt.Println(sw.Elapsed())
}

func findTarget(program []int64, target int64) (int64, int64) {
	searchStart := int64(0)
	searchEnd := int64(250)
	c := intcode.NewComputer(program)
	for a := searchStart; a < searchEnd; a++ {
		for b := searchStart; b < searchEnd; b++ {
			p := make([]int64, len(program))
			copy(p, program)
			p[1] = a
			p[2] = b
			c.LoadProgram(p)
			succeeded := c.Run()
			if succeeded && c.Memory[0] == target {
				return a, b
			}
		}
	}
	return -1, -1
}
