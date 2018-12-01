package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int64{1, -2, 3, 1}), int64(3), "Invalid sum")
	assert.Equal(t, Sum([]int64{1, 1, 1}), int64(3), "Invalid sum")
	assert.Equal(t, Sum([]int64{1, 1, -2}), int64(0), "Invalid sum")
	assert.Equal(t, Sum([]int64{-1, -2, -3}), int64(-6), "Invalid sum")
}

func TestFirstSumRepeat(t *testing.T) {
	assert.Equal(t, FirstSumRepeat([]int64{1, -2, 3, 1}), int64(2), "Invalid repeat")
	assert.Equal(t, FirstSumRepeat([]int64{1, -1}), int64(0), "Invalid repeat")
	assert.Equal(t, FirstSumRepeat([]int64{3, 3, 4, -2, -4}), int64(10), "Invalid repeat")
	assert.Equal(t, FirstSumRepeat([]int64{-6, 3, 8, 5, -6}), int64(5), "Invalid repeat")
	assert.Equal(t, FirstSumRepeat([]int64{7, 7, -2, -7, -4}), int64(14), "Invalid repeat")
}
