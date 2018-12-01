package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sum(input []int64) int64 {
	sum := int64(0)
	for _, i := range input {
		sum += i
	}
	return sum
}

func firstSumRepeat(input []int64) int64 {
	sum := int64(0)
	repeat := int64(0)
	seen := make(map[int64]bool)
	seen[0] = true

InfiniteLoop:
	for {
		for _, i := range input {
			sum += i
			_, exists := seen[sum]
			if exists {
				repeat = sum
				break InfiniteLoop
			} else {
				seen[sum] = true
			}
		}
	}

	return repeat
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
		input = append(input, i)
	}
	check(scanner.Err())
	fmt.Printf("Sum: %d\nFirst Repeat: %d\n", sum(input), firstSumRepeat(input))
}
