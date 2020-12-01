package aoc

import (
	"bufio"
	"os"
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
