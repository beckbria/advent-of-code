package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(157), step1(testInput))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(70), step2(testInput))
}
