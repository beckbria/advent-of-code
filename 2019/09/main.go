package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

// https://adventofcode.com/2019/day/9
// Add a new parameter mode to the intcode computer.  Run a new program.

func main() {
	program := intcode.ReadIntCode("2019/09/input.txt")
	sw := lib.NewStopwatch()
	// Part 1
	io := intcode.NewConstantInputOutput(int64(1))
	c := intcode.NewComputer(program)
	c.Io = io
	c.Run()
	fmt.Println(io.Outputs)
	// Part 2
	c.Reset()
	io = intcode.NewConstantInputOutput(int64(2))
	c.Io = io
	c.Run()
	fmt.Println(io.Outputs)
	fmt.Println(sw.Elapsed())
}
