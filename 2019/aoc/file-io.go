package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFileLines opens a file and reads each line as a string
func ReadFileLines(fileName string) []string {
	file, err := os.Open(fileName)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	Check(scanner.Err())
	return input
}

// ReadIntCode reads a program consisting of IntCode instructions - comma-separated integers
func ReadIntCode(fileName string) []int64 {
	lines := ReadFileLines(fileName)
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

