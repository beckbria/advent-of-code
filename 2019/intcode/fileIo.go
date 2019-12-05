package intcode

import (
	"../aoc"
	"log"
	"strconv"
	"strings"
)

// ReadIntCode reads a program consisting of IntCode instructions - comma-separated integers
func ReadIntCode(fileName string) []int64 {
	lines := aoc.ReadFileLines(fileName)
	if len(lines) > 1 {
		log.Fatalf("ReadIntCode expects a single line of input")
	}
	nums := strings.Split(lines[0], ",")
	program := []int64{}
	for _, n := range nums {
		i, _ := strconv.ParseInt(n, 10, 64)
		program = append(program, i)
	}
	return program
}
