package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = [][]int64{
	{1000, 2000, 3000},
	{4000},
	{5000, 6000},
	{7000, 8000, 9000},
	{10000},
}

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(24000), step1(testInput))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(45000), step2(testInput))
}
