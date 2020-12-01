package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	digits := readDigits(input[0])
	fmt.Println(readDigitSlice(fft(digits, 100)[:8]))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(part2(digits))
	fmt.Println(sw.Elapsed())
}

func readDigits(input string) []int64 {
	digits := []int64{}
	for _, r := range []rune(input) {
		digits = append(digits, int64(r-'0'))
	}
	return digits
}

func fft(digits []int64, rounds int) []int64 {
	result := make([]int64, len(digits))
	copy(result, digits)
	for round := 0; round < rounds; round++ {
		result = fftRound(result)
	}
	return result
}

func fftRound(digits []int64) []int64 {
	result := make([]int64, len(digits))

	for i := 0; i < len(digits); i++ {
		p := buildPattern(i+1, len(digits))
		sum := int64(0)
		for j := 0; j < len(digits); j++ {
			sum += digits[j] * p[j]
		}
		result[i] = aoc.Abs(sum % 10)
	}

	return result
}

var basePattern = []int64{0, 1, 0, -1}

func buildPattern(pos, length int) []int64 {
	pattern := make([]int64, length+1) // We need to drop the first character, so add one more on the end

	for i := 0; i < len(pattern); i++ {
		pattern[i] = basePattern[(i/pos)%len(basePattern)]
	}

	return pattern[1:]
}

func part2(baseDigits []int64) int {
	// Build the full sequence
	digits := []int64{}
	for i := 0; i < 10000; i++ {
		digits = append(digits, baseDigits...)
	}
	// Keep only the bits starting at our offset
	answerIndex := readDigitSlice(digits[:7])
	digits = digits[answerIndex:]

	// Looking at the last few digits of each iteration, for the later digits of the sequence, the
	// formula each iteration is:
	// d[i] = sum(d[i:]) % 10
	for round := 0; round < 100; round++ {
		sum := int64(0)
		for i := len(digits) - 1; i >= 0; i-- {
			sum += digits[i]
			digits[i] = aoc.Abs(sum) % 10
		}
	}

	return readDigitSlice(digits[:8])
}

func readDigitSlice(digits []int64) int {
	answer := 0
	for _, d := range digits {
		answer *= 10
		answer += int(d)
	}
	return answer
}
