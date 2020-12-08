package cpu

import (
	"strconv"
	"strings"

	"../../aoc"
)

// ReadProgram reads a program consisting of CPU instructions, one per line, from a file
func ReadProgram(fileName string) []Instruction {
	lines := aoc.ReadFileLines(fileName)
	return ParseProgram(lines)
}

// ParseProgram reads a series of instructions, one per line, into a CPU program
func ParseProgram(lines []string) []Instruction {
	inst := make([]Instruction, 0)
	for _, l := range lines {
		tokens := strings.Split(l, " ")
		i := Instruction{Op: opFromString(tokens[0])}
		val, err := strconv.Atoi(tokens[1])
		aoc.Check(err)
		i.Num = Data(val)
		inst = append(inst, i)
	}
	return inst
}
