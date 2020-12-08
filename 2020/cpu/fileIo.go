package cpu

import (
	"strconv"
	"strings"

	"../../aoc"
)

// ReadIntCode reads a program consisting of IntCode instructions - comma-separated integers
func ReadProgram(fileName string) []Instruction {
	lines := aoc.ReadFileLines(fileName)
	return ParseProgram(lines)
}

func ParseProgram(lines []string) []Instruction {
	inst := make([]Instruction, 0)
	for _, l := range lines {
		tokens := strings.Split(l, " ")
		i := Instruction{Op: tokens[0]}
		val, err := strconv.Atoi(tokens[1])
		aoc.Check(err)
		i.Num = val
		inst = append(inst, i)
	}
	return inst
}
