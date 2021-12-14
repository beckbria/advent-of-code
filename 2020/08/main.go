package main

import (
	"fmt"
	"log"

	"github.com/beckbria/advent-of-code/2020/lib"
	. "github.com/beckbria/advent-of-code/2020/cpu"
)

// https://adventofcode.com/2020/day/8
// Another year, another assembly language computer.

func main() {
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	instructions := ReadProgram("input.txt")
	fmt.Println(step1(instructions))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(instructions))
	fmt.Println(sw.Elapsed())
}

// Step 1: Value of the accumulator when it finds an infinite loop
func step1(inst []Instruction) Data {
	c := NewComputer(inst)
	c.FindInfiniteLoop()
	return c.Acc
}

// Step 2: One instruction must be altered to make the program terminate.
// Which? What's the value of the accumulator when it terminates?
func step2(inst []Instruction) Data {
	for idx, i := range inst {
		if i.Op == OpAcc {
			continue
		}
		newProgram := append([]Instruction{}, inst...)
		newProgram[idx].Op = inverseOp(i.Op)
		c := NewComputer(newProgram)
		looped := c.FindInfiniteLoop()
		if !looped {
			return c.Acc
		}
	}
	return -1
}

func inverseOp(op OpCode) OpCode {
	switch op {
	case OpNop:
		return OpJmp
	case OpJmp:
		return OpNop
	}
	log.Fatalf("Unexpected inverse op: %d", op)
	return OpNop
}
