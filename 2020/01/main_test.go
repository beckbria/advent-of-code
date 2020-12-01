package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/3
func Test2(t *testing.T) {
	a, b := findSum2([]int{1721, 979, 366, 299, 675, 1456}, 2020)
	assert.Equal(t, a*b, 514579)
}

func Test3(t *testing.T) {
	a, b, c := findSum3([]int{1721, 979, 366, 299, 675, 1456}, 2020)
	assert.Equal(t, a*b*c, 241861950)
}
