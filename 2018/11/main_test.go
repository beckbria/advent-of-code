package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPowerLocation(t *testing.T) {
	x, y := LargestPowerLocation(18, 3, 3)
	assert.Equal(t, 33, x)
	assert.Equal(t, 45, y)
	x, y = LargestPowerLocation(42, 3, 3)
	assert.Equal(t, 21, x)
	assert.Equal(t, 61, y)
}

func TestMakeGrid(t *testing.T) {
	assert.Equal(t, -5, MakeGrid(57)[122][79])
	assert.Equal(t, 0, MakeGrid(39)[217][196])
	assert.Equal(t, 4, MakeGrid(71)[101][153])
}

func TestLargestPowerSquare(t *testing.T) {
	// TODO: Determine why this test case doesn't match the supplied input.
	// The other test case and the answer for seed 1718 both come out correctly
	x, y, size := LargestPowerSquare(18)
	assert.Equal(t, 90, x)
	assert.Equal(t, 269, y)
	assert.Equal(t, 16, size)
	x, y, size = LargestPowerSquare(42)
	assert.Equal(t, 232, x)
	assert.Equal(t, 251, y)
	assert.Equal(t, 12, size)
}
