package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

// https://adventofcode.com/2019/day/5
// Implement new instructions in the IntCode computer, run a program,
// return the diagnostic code it emits immediately prior to terminating

func main() {
	program := intcode.ReadIntCode("2019/05/input.txt")
	sw := lib.NewStopwatch()
	io1 := intcode.NewConstantInputOutput(1)
	c := intcode.NewComputer(program)
	c.Io = io1
	c.Run()
	fmt.Print("Part 1: ")
	fmt.Println(io1.Outputs[len(io1.Outputs)-1])

	c.Reset()
	io2 := intcode.NewConstantInputOutput(5)
	c.Io = io2
	c.Run()
	fmt.Print("Part 2: ")
	fmt.Println(io2.Outputs[len(io2.Outputs)-1])

	fmt.Println(sw.Elapsed())
}
