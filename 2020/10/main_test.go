package main

import (
	"testing"

	"../../aoc"
	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, 35, step1(aoc.Int64Slice{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}))
	assert.Equal(t, 220, step1(aoc.Int64Slice{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(8), step2(aoc.Int64Slice{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}))
	assert.Equal(t, int64(19208), step2(aoc.Int64Slice{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}))
}
