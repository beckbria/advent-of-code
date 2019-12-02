package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunProgramBaseCase(t *testing.T) {
	program := []int64{1,9,10,3,2,3,11,0,99,30,40,50}
	p := RunProgram(program)
	assert.Equal(t, []int64{3500,9,10,70,2,3,11,0,99,30,40,50}, p)
}

func TestRunProgramExample1(t *testing.T) {
	program := []int64{1,0,0,0,99}
	p := RunProgram(program)
	assert.Equal(t, []int64{2,0,0,0,99}, p)
}

func TestRunProgramExample2(t *testing.T) {
	program := []int64{2,3,0,3,99}
	p := RunProgram(program)
	assert.Equal(t, []int64{2,3,0,6,99}, p)
}

func TestRunProgramExample3(t *testing.T) {
	program := []int64{2,4,4,5,99,0}
	p := RunProgram(program)
	assert.Equal(t, []int64{2,4,4,5,99,9801}, p)
}

func TestRunProgramExample4(t *testing.T) {
	program := []int64{1,1,1,4,99,5,6,0,99}
	p := RunProgram(program)
	assert.Equal(t, p, []int64{30,1,1,4,2,5,6,0,99})
}