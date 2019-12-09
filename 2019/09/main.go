package main

import (
	"fmt"

	"../aoc"
	"../intcode"
)

// https://adventofcode.com/2019/day/9

func main() {
	program := intcode.ReadIntCode("input.txt")
	sw := aoc.NewStopwatch()
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
