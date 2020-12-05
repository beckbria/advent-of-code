package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	min, max := split(0, 127, true)
	assert.Equal(t, min, 0)
	assert.Equal(t, max, 63)
	min, max = split(0, 63, false)
	assert.Equal(t, min, 32)
	assert.Equal(t, max, 63)
	min, max = split(32, 63, true)
	assert.Equal(t, min, 32)
	assert.Equal(t, max, 47)
	min, max = split(32, 47, false)
	assert.Equal(t, min, 40)
	assert.Equal(t, max, 47)
	min, max = split(40, 47, false)
	assert.Equal(t, min, 44)
	assert.Equal(t, max, 47)
	min, max = split(44, 47, true)
	assert.Equal(t, min, 44)
	assert.Equal(t, max, 45)
	min, max = split(44, 45, true)
	assert.Equal(t, min, 44)
	assert.Equal(t, max, 44)

	min, max = split(0, 7, false)
	assert.Equal(t, min, 4)
	assert.Equal(t, max, 7)
	min, max = split(4, 7, true)
	assert.Equal(t, min, 4)
	assert.Equal(t, max, 5)
	min, max = split(4, 5, false)
	assert.Equal(t, min, 5)
	assert.Equal(t, max, 5)
}
