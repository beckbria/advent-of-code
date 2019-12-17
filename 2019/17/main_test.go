package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/16
func TestRobotProgram(t *testing.T) {
	assert.Equal(
		t,
		[]int64{
			65, 44, 66, 44, 67, 44, 66, 44, 65, 44, 67, 10,
			82, 44, 56, 44, 82, 44, 56, 10,
			82, 44, 52, 44, 82, 44, 52, 44, 82, 44, 56, 10,
			76, 44, 54, 44, 76, 44, 50, 10,
			114, 10},
		robotProgram([]string{
			"A,B,C,B,A,C", "R,8,R,8", "R,4,R,4,R,8", "L,6,L,2"}))
}
