package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestStep1(t *testing.T) {
	assert.Equal(t, "CMZ", step1(testInput))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, "MCD", step2(testInput))
}
