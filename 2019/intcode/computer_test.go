package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/2
func TestRunProgramBaseCase(t *testing.T) {
	program := Program{1,9,10,3,2,3,11,0,99,30,40,50}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{3500,9,10,70,2,3,11,0,99,30,40,50}, c.Memory)
}

func TestRunProgramExample1(t *testing.T) {
	program := Program{1,0,0,0,99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2,0,0,0,99}, c.Memory)
}

func TestRunProgramExample2(t *testing.T) {
	program := Program{2,3,0,3,99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2,3,0,6,99}, c.Memory)
}

func TestRunProgramExample3(t *testing.T) {
	program := Program{2,4,4,5,99,0}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2,4,4,5,99,9801}, c.Memory)
}

func TestRunProgramExample4(t *testing.T) {
	program := Program{1,1,1,4,99,5,6,0,99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t,  Program{30,1,1,4,2,5,6,0,99}, c.Memory)
}

// Generates a map representation of a program to compare memory to
func assertMemory(t *testing.T, expected Program, actual map[Address]Instruction) {
	eMem := make(map[Address]Instruction)
	for addr, val := range expected {
		eMem[int64(addr)] = val
	}
	assert.Equal(t, eMem, actual)
}