package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Player 1:",
	"9",
	"2",
	"6",
	"3",
	"1",
	"",
	"Player 2:",
	"5",
	"8",
	"4",
	"7",
	"10",
}

func TestStep1(t *testing.T) {
	d := parseDecks(input)
	assert.Equal(t, int64(306), step1(d))
}

func TestStep2(t *testing.T) {
	d := parseDecks(input)
	assert.Equal(t, int64(291), step2(d))
}
