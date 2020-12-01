package main

import (
	"testing"

	"../../aoc"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/16
func TestReadDigits(t *testing.T) {
	assert.Equal(t, []int64{1, 4, 2, 3}, readDigits("1423"))
}

func TestModulo(t *testing.T) {
	assert.Equal(t, 7, 37%10)
	assert.Equal(t, int64(7), aoc.Abs(-17%10))
}

func TestPattern(t *testing.T) {
	assert.Equal(t, []int64{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0, 0, -1}, buildPattern(2, 14))
}

func TestFft(t *testing.T) {
	assert.Equal(t, []int64{4, 8, 2, 2, 6, 1, 5, 8}, fft([]int64{1, 2, 3, 4, 5, 6, 7, 8}, 1))
	assert.Equal(t, []int64{3, 4, 0, 4, 0, 4, 3, 8}, fft([]int64{1, 2, 3, 4, 5, 6, 7, 8}, 2))
	assert.Equal(t, []int64{0, 3, 4, 1, 5, 5, 1, 8}, fft([]int64{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	assert.Equal(t, []int64{0, 1, 0, 2, 9, 4, 9, 8}, fft([]int64{1, 2, 3, 4, 5, 6, 7, 8}, 4))
}

func TestFftExtended(t *testing.T) {
	assert.Equal(t, []int64{2, 4, 1, 7, 6, 1, 7, 6}, fft(readDigits("80871224585914546619083218645595"), 100)[:8])
	assert.Equal(t, []int64{7, 3, 7, 4, 5, 4, 1, 8}, fft(readDigits("19617804207202209144916044189917"), 100)[:8])
	assert.Equal(t, []int64{5, 2, 4, 3, 2, 1, 3, 3}, fft(readDigits("69317163492948606335995924319873"), 100)[:8])
}

func TestReadDigitSlice(t *testing.T) {
	assert.Equal(t, 1234, readDigitSlice([]int64{1, 2, 3, 4}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 84462026, part2(readDigits("03036732577212944063491565474664")))
}
