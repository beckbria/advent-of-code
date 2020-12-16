package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"class: 1-3 or 5-7",
	"row: 6-11 or 33-44",
	"seat: 13-40 or 45-50",
	"",
	"your ticket:",
	"7,1,14",
	"",
	"nearby tickets:",
	"7,3,47",
	"40,4,50",
	"55,2,20",
	"38,6,12",
}

func TestStep1(t *testing.T) {
	rules, _, others := parseInput(input)
	answer, invalidIdx := step1(rules, others)
	assert.Equal(t, 71, answer)
	assert.Equal(t, []int{1, 2, 3}, invalidIdx)
}

func TestFindCandidates(t *testing.T) {
	input := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	rules, _, others := parseInput(input)
	_, invalidIdx := step1(rules, others)
	tickets := pruneInvalid(others, invalidIdx)
	candidates := findCandidates(rules, tickets)
	assert.Equal(t, 0, candidates["row"])
	assert.Equal(t, 1, candidates["class"])
	assert.Equal(t, 2, candidates["seat"])
}
