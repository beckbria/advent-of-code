package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = []string{"A Y","B X","C Z"}

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(15), step1(rawInput))
}

func TestStep1AllCombos(t *testing.T) {
	rawAll := []string{"A X", "A Y", "A Z", "B X", "B Y", "B Z", "C X", "C Y", "C Z"}
	expectedScore := (
		1 * 3 +	// three rock throws
		2 * 3 +	// three paper throws
		3 * 3 +	// three scissors throws
		3 * 3 +	// 3 ties
		3 * 6)	// 3 victories
	
	assert.Equal(t, int64(expectedScore), step1(rawAll))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(12), step2(rawInput))
}