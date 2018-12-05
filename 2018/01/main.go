package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Sum returns the sum of a list of integers
func Sum(input []int64) int64 {
	sum := int64(0)
	for _, i := range input {
		sum += i
	}
	return sum
}

// FirstSumRepeat loops over a list of integers, keeping a rolling sum, and returns the first default sum seen.
func FirstSumRepeat(input []int64) int64 {
	sum := int64(0)
	seen := make(map[int64]bool)
	seen[0] = true

	for {
		for _, i := range input {
			sum += i
			_, exists := seen[sum]
			if exists {
				return sum
			} else {
				seen[sum] = true
			}
		}
	}
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
	start := time.Now()
	fmt.Printf("Sum: %d\n", Sum(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("First Repeat: %d\n", FirstSumRepeat(input))
	fmt.Println(time.Since(start))
}
