package aoc

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
	for i := 0; i < 1000000; i++ {}
	elapsed1 := int64(sw.Elapsed())
	for i := 0; i < 1000000; i++ {}
	elapsed2 := int64(sw.Elapsed())
	assert.Greater(t, elapsed1, int64(0))
	assert.Greater(t, elapsed2, elapsed1)
}

func TestStopwatchReset(t *testing.T) {
	sw1 := NewStopwatch()
	sw2 := NewStopwatch()
	for i := 0; i < 1000000; i++ {}
	sw1.Reset()
	for i := 0; i < 1000000; i++ {}
	assert.Greater(t, int64(sw2.Elapsed()), int64(sw1.Elapsed()))
}