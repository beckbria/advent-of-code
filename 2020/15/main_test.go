package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryGame(t *testing.T) {
	input := []int64{0, 3, 6}
	assert.Equal(t, int64(4), memoryGame(input, 9))
	assert.Equal(t, int64(0), memoryGame(input, 10))
	assert.Equal(t, int64(175594), memoryGame(input, 30000000))
}
