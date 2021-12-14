package lib

import (
	"bufio"
	"os"
	"strconv"
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

// ReadFileNumbers opens a file and reads each line as a number
func ReadFileNumbers(fileName string) []int64 {
	lines := ReadFileLines(fileName)
	nums := []int64{}
	for _, l := range lines {
		n, err := strconv.ParseInt(l, 10, 64)
		Check(err)
		nums = append(nums, n)
	}
	return nums
}
