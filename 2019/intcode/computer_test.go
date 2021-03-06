package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/2
func TestRunProgramBaseCase(t *testing.T) {
	program := Program{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, c.Memory)
}

func TestRunProgramExample1(t *testing.T) {
	program := Program{1, 0, 0, 0, 99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2, 0, 0, 0, 99}, c.Memory)
}

func TestRunProgramExample2(t *testing.T) {
	program := Program{2, 3, 0, 3, 99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2, 3, 0, 6, 99}, c.Memory)
}

func TestRunProgramExample3(t *testing.T) {
	program := Program{2, 4, 4, 5, 99, 0}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{2, 4, 4, 5, 99, 9801}, c.Memory)
}

func TestRunProgramExample4(t *testing.T) {
	program := Program{1, 1, 1, 4, 99, 5, 6, 0, 99}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{30, 1, 1, 4, 2, 5, 6, 0, 99}, c.Memory)
}

// Examples of programs from https://adventofcode.com/2019/day/5

func TestImmediateAddressing(t *testing.T) {
	program := Program{1002, 4, 3, 4, 33}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{1002, 4, 3, 4, 99}, c.Memory)
}

func TestNegativeValues(t *testing.T) {
	program := Program{1101, 100, -1, 4, 0}
	c := NewComputer(program)
	c.Run()
	assertMemory(t, Program{1101, 100, -1, 4, 99}, c.Memory)
}

func TestEqualPosition(t *testing.T) {
	// Outputs 1 iff input value is 8
	program := Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	io := NewConstantInputOutput(8)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1}, io.Outputs)

	c.Reset()
	io = NewConstantInputOutput(5)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{0}, io.Outputs)
}

func TestEqualImmediate(t *testing.T) {
	// Outputs 1 iff input value is 8
	program := Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	io := NewConstantInputOutput(8)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1}, io.Outputs)

	c.Reset()
	io = NewConstantInputOutput(5)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{0}, io.Outputs)
}

func TestLessThanPosition(t *testing.T) {
	// Outputs 1 iff input value is <8
	program := Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	io := NewConstantInputOutput(8)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{0}, io.Outputs)

	c.Reset()
	io = NewConstantInputOutput(5)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1}, io.Outputs)
}

func TestLessThanImmediate(t *testing.T) {
	// Outputs 1 iff input value is <8
	program := Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	io := NewConstantInputOutput(8)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{0}, io.Outputs)

	c.Reset()
	io = NewConstantInputOutput(5)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1}, io.Outputs)
}

// Examples of programs from https://adventofcode.com/2019/day/9
func TestRelativeMode1(t *testing.T) {
	program := Program{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	io := NewConstantInputOutput(0)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, program, io.Outputs)
}

func TestRelativeMode2(t *testing.T) {
	program := Program{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	io := NewConstantInputOutput(0)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1219070632396864}, io.Outputs)
}

func TestRelativeMode3(t *testing.T) {
	program := Program{104, 1125899906842624, 99}
	io := NewConstantInputOutput(0)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{1125899906842624}, io.Outputs)
}

// Other test cases
func TestConstantInputOutput(t *testing.T) {
	program := Program{3, 0, 4, 0, 99}
	io := NewConstantInputOutput(5)
	c := NewComputer(program)
	c.Io = io
	c.Run()
	assert.Equal(t, []int64{5}, io.Outputs)
}

// Generates a map representation of a program to compare memory to
func assertMemory(t *testing.T, expected Program, actual map[Address]Instruction) {
	eMem := make(map[Address]Instruction)
	for addr, val := range expected {
		eMem[int64(addr)] = val
	}
	assert.Equal(t, eMem, actual)
}
