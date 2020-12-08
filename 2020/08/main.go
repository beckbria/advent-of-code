package main

import (
	"fmt"
	"log"

	"../../aoc"
	"../cpu"
)

// https://adventofcode.com/2020/day/8
// Another year, another assembly language computer.

func main() {
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	instructions := cpu.ReadProgram("input.txt")
	fmt.Println(step1(instructions))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(instructions))
	fmt.Println(sw.Elapsed())
}

// Step 1: Value of the accumulator when it finds an infinite loop
func step1(inst []cpu.Instruction) int {
	c := cpu.NewComputer(inst)
	c.FindInfiniteLoop()
	return c.Acc
}

// Step 2: How many bags musta shiny gold bag contain?
func step2(inst []cpu.Instruction) int {
	for idx, i := range inst {
		if i.Op == "acc" {
			continue
		}
		newProgram := append([]cpu.Instruction{}, inst...)
		newProgram[idx].Op = inverseOp(i.Op)
		c := cpu.NewComputer(newProgram)
		looped := c.FindInfiniteLoop()
		if !looped {
			return c.Acc
		}
	}
	return -1
}

func inverseOp(op string) string {
	switch op {
	case "nop":
		return "jmp"
	case "jmp":
		return "nop"
	}
	log.Fatalf("Unexpected inverse op: %q", op)
	return ""
}
