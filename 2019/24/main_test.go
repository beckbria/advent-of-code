package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/22

func TestScore(t *testing.T) {
	b := readBugs([]string{
		".....",
		".....",
		".....",
		"#....",
		".#..."})
	assert.Equal(t, int64(2129920), b.score())
}

func TestCountAfter(t *testing.T) {
	b := readBugs([]string{
		"....#",
		"#..#.",
		"#..##",
		"..#..",
		"#...."})
	assert.Equal(t, int64(99), countAfter(b, 10))
}
