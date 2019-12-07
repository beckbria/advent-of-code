package intcode

import (
	"log"
	"strconv"
	"strings"

	"../aoc"
)

// ReadIntCode reads a program consisting of IntCode instructions - comma-separated integers
func ReadIntCode(fileName string) []Instruction {
	lines := aoc.ReadFileLines(fileName)
	if len(lines) > 1 {
		log.Fatalf("ReadIntCode expects a single line of input")
	}
	return ParseProgram(lines[0])
}

// ReadIntCodePrograms reads a series of IntCode programs, one per line
func ReadIntCodePrograms(fileName string) []Program {
	programs := make([][]Instruction, 0)
	for _, line := range aoc.ReadFileLines(fileName) {
		programs = append(programs, ParseProgram(line))
	}
	return programs
}

// ParseProgram reads an intcode program from a string containing comma-separated instructions
func ParseProgram(line string) Program {
	nums := strings.Split(line, ",")
	program := []int64{}
	for _, n := range nums {
		i, _ := strconv.ParseInt(n, 10, 64)
		program = append(program, i)
	}
	return program
}
