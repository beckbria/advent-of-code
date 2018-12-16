package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidOpcodes(t *testing.T) {
	inst := ReadInstruction([]string{
		"Before: [3, 2, 1, 1]",
		"9 2 1 2",
		"After:  [3, 2, 2, 1]"})
	assert.Equal(t, []string{"addi", "mulr", "seti"}, ValidOpcodes(inst))
}

func TestMulR(t *testing.T) {
	inst := ReadInstruction([]string{
		"Before: [3, 2, 1, 1]",
		"9 2 1 2",
		"After:  [3, 2, 2, 1]"})
	assert.Equal(t, inst.regAfter, Mulr(inst))
}
