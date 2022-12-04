package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(2), step1(testInput))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(4), step2(testInput))
}
