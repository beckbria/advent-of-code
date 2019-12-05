package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/3
func TestExample1(t *testing.T) {
	manhattan, cost := bestValues("R8,U5,L5,D3", "U7,R6,D4,L4")
	assert.Equal(t, int64(6), manhattan)
	assert.Equal(t, 30, cost)
}

func TestExample2(t *testing.T) {
	manhattan, cost := bestValues("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83")
	assert.Equal(t, int64(159), manhattan)
	assert.Equal(t, 610, cost)
}

func TestExample3(t *testing.T) {
	manhattan, cost := bestValues("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	assert.Equal(t, int64(135), manhattan)
	assert.Equal(t, 410, cost)
}
