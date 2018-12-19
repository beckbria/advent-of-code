package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	comp := ReadComputer([]string{
		"#ip 0",
		"seti 5 0 1",
		"seti 6 0 2",
		"addi 0 1 0",
		"addr 1 2 3",
		"setr 1 0 0",
		"seti 8 0 4",
		"seti 9 0 5"})
	comp.run()
	assert.Equal(t, Registers{6, 5, 6, 0, 0, 9}, comp.reg)
}
