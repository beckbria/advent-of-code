package main

import (
	"testing"

	"github.com/beckbria/advent-of-code/2020/lib"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/3
func Test2(t *testing.T) {
	_, a, b := lib.FindSum2([]int64{1721, 979, 366, 299, 675, 1456}, 2020)
	assert.Equal(t, a*b, int64(514579))
}

func Test3(t *testing.T) {
	a, b, c := findSum3([]int64{1721, 979, 366, 299, 675, 1456}, 2020)
	assert.Equal(t, a*b*c, int64(241861950))
}
