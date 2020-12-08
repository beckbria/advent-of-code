package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/8

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	instructions := readInstructions(lines)
	fmt.Println(step1(instructions))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(instructions))
	fmt.Println(sw.Elapsed())
}

type Instruction struct {
	op  string
	num int
}

func readInstructions(lines []string) []Instruction {
	inst := make([]Instruction, 0)
	for _, l := range lines {
		tokens := strings.Split(l, " ")
		i := Instruction{op: tokens[0]}
		val, err := strconv.Atoi(tokens[1])
		aoc.Check(err)
		i.num = val
		inst = append(inst, i)
	}
	return inst
}

type Computer struct {
	ip, acc int
	inst    []Instruction
}

func NewComputer(program []Instruction) Computer {
	return Computer{
		ip:   0,
		acc:  0,
		inst: program,
	}
}

func (c *Computer) Step() bool {
	i := &c.inst[c.ip]
	switch i.op {
	case "nop":
		c.ip++
	case "acc":
		c.acc += i.num
		c.ip++
	case "jmp":
		c.ip += i.num
	}
	return c.ip != len(c.inst)
}

func (c *Computer) FindInfiniteLoop() bool {
	seen := make(map[int]bool)
	for !seen[c.ip] {
		seen[c.ip] = true
		stillRunning := c.Step()
		if !stillRunning {
			return false
		}
	}
	return true
}

// Step 1: Value of the accumulator when it finds an infinite loop
func step1(inst []Instruction) int {
	c := NewComputer(inst)
	c.FindInfiniteLoop()
	return c.acc
}

// Step 2: How many bags musta shiny gold bag contain?
func step2(inst []Instruction) int {
	for idx, i := range inst {
		if i.op == "acc" {
			continue
		}
		newProgram := append([]Instruction{}, inst...)
		newProgram[idx].op = inverseOp(i.op)
		c := NewComputer(newProgram)
		looped := c.FindInfiniteLoop()
		if !looped {
			return c.acc
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
