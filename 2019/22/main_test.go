package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/22

func TestNewDeck(t *testing.T) {
	d := newDeck(10)
	assert.Equal(t, deck{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, d)
}

func TestCut(t *testing.T) {
	d := newDeck(10)
	d.cut(3)
	assert.Equal(t, deck{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, d)
}

func TestCutNegative(t *testing.T) {
	d := newDeck(10)
	d.cut(-4)
	assert.Equal(t, deck{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, d)
}

func TestDealIncrement(t *testing.T) {
	d := newDeck(10)
	d.increment(3)
	assert.Equal(t, deck{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, d)
}

func TestReverse(t *testing.T) {
	d := newDeck(10)
	d.reverse()
	assert.Equal(t, deck{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, d)
}

func TestShuffle1(t *testing.T) {
	inst := parseInstructions([]string{
		"deal with increment 7",
		"deal into new stack",
		"deal into new stack"})
	d := newDeck(10)
	d.shuffle(inst)
	assert.Equal(t, deck{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}, d)
}

func TestShuffle2(t *testing.T) {
	inst := parseInstructions([]string{
		"cut 6",
		"deal with increment 7",
		"deal into new stack"})
	d := newDeck(10)
	d.shuffle(inst)
	assert.Equal(t, deck{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}, d)
}

func TestShuffle3(t *testing.T) {
	inst := parseInstructions([]string{
		"deal with increment 7",
		"deal with increment 9",
		"cut -2"})
	d := newDeck(10)
	d.shuffle(inst)
	assert.Equal(t, deck{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}, d)
}

func TestShuffle4(t *testing.T) {
	inst := parseInstructions([]string{
		"deal into new stack",
		"cut -2",
		"deal with increment 7",
		"cut 8",
		"cut -4",
		"deal with increment 7",
		"cut 3",
		"deal with increment 9",
		"deal with increment 3",
		"cut -1"})
	d := newDeck(10)
	d.shuffle(inst)
	assert.Equal(t, deck{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}, d)
}
