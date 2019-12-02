package main

import (
	"fmt"

	"../aoc"
)

const add = int64(1)
const multiply = int64(2)
const terminate = int64(99)
const debug = false

func main() {
	program := aoc.ReadIntCode("input.txt")
	sw := aoc.NewStopwatch()

	// Corrections from instruction
	program[1] = int64(12)
	program[2] = int64(2)

	step1 := make([]int64, len(program))
	copy(step1, program)
	RunProgram(step1)
	fmt.Println(step1[0])
	fmt.Println(sw.Elapsed())

	a, b := findTarget(program, int64(19690720))
	fmt.Printf("%d %d -> %d\n", a, b, (100*a)+b)
	fmt.Println(sw.Elapsed())
}

func findTarget(program []int64, target int64) (int64, int64) {
	searchStart := int64(0)
	searchEnd := int64(1000)
	for a := searchStart; a < searchEnd; a++ {
		for b := searchStart; b < searchEnd; b++ {
			p := make([]int64, len(program))
			copy(p, program)
			p[1] = a
			p[2] = b
			succeeded := RunProgram(p)
			if succeeded && p[0] == target {
				return a, b
			}
		}
	}
	return -1, -1
}

// RunProgram simulates an Intcode computer.  If the program attempts to access memory out of bounds, returns false
func RunProgram(program []int64) bool {
	// PC tracks the current instruction
	done := false
	if (debug) {
		fmt.Print("Initial State: ")
		fmt.Println(program)
	}
	plen := int64(len(program))
	for ip := int64(0); ip < plen; ip += 4 {
		if (debug) {
			fmt.Printf("IP %d: ", ip)
		}
		targetSource := ip + 3
		source1 := ip + 1
		source2 := ip + 2
		switch program[ip] {
		case add:
			target := program[targetSource]
			s1 := program[source1]
			s2 := program[source2]
			if target < 0 || target >= plen || s1 < 0 || s1 >= plen || s2 < 0 || s2 >= plen {
				return false
			}
			if (debug) {
				fmt.Printf("Setting i%d to i%d(%d) + i%d(%d)\n", target, s1, program[s1], s2, program[s2])
			}
			program[target] = program[s1] + program[s2]
		case multiply:
			target := program[targetSource]
			s1 := program[source1]
			s2 := program[source2]
			if target < 0 || target >= plen || s1 < 0 || s1 >= plen || s2 < 0 || s2 >= plen {
				return false
			}
			if (debug) {
				fmt.Printf("Setting i%d to i%d(%d) * i%d(%d)\n", target, s1, program[s1], s2, program[s2])
			}
			program[target] = program[s1] * program[s2]
		case terminate:
			if (debug) {
				fmt.Println("TERMINATE")
			}
			done = true
		default:
			fmt.Printf("Unexpected instruction: %d\n", program[ip])
		}
		if (debug) {
			fmt.Println(program)
		}
		if done {
			break
		}
	}
	return true
}