package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{"F10", "N3", "F7", "R90", "F11"}

func TestStep1(t *testing.T) {
	moves := readMoves(input)
	assert.Equal(t, int64(25), step1(moves))
}

func TestStep2(t *testing.T) {
	moves := readMoves(input)
	assert.Equal(t, int64(286), step2(moves))
}
