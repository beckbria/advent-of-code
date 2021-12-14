package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileLines(t *testing.T) {
	lines := ReadFileLines("test-input.txt")
	assert.Equal(t, []string{"First Line", "Second Line", "Third Line"}, lines)
}

func TestStopwatchIncreasing(t *testing.T) {
	sw := NewStopwatch()
	for i := 0; i < 1000000; i++ {
	}
	elapsed1 := int64(sw.Elapsed())
	for i := 0; i < 1000000; i++ {
	}
	elapsed2 := int64(sw.Elapsed())
	assert.Greater(t, elapsed1, int64(0))
	assert.Greater(t, elapsed2, elapsed1)
}

func TestStopwatchReset(t *testing.T) {
	sw1 := NewStopwatch()
	sw2 := NewStopwatch()
	for i := 0; i < 1000000; i++ {
	}
	sw1.Reset()
	for i := 0; i < 1000000; i++ {
	}
	assert.Greater(t, int64(sw2.Elapsed()), int64(sw1.Elapsed()))
}

func TestPermutation(t *testing.T) {
	arr := []int64{1, 2, 3}
	assert.True(t, NextPermutation(arr))
	assert.Equal(t, []int64{1, 3, 2}, arr)
	assert.True(t, NextPermutation(arr))
	assert.Equal(t, []int64{2, 1, 3}, arr)
	assert.True(t, NextPermutation(arr))
	assert.Equal(t, []int64{2, 3, 1}, arr)
	assert.True(t, NextPermutation(arr))
	assert.Equal(t, []int64{3, 1, 2}, arr)
	assert.True(t, NextPermutation(arr))
	assert.Equal(t, []int64{3, 2, 1}, arr)
	assert.False(t, NextPermutation(arr))
}
