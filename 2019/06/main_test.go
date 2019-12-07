package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/6
func TestOrbits(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L"}
	bodies := parseBodies(lines)
	assert.Equal(t, 42, totalOrbits(bodies))
}

func TestDistance(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN"}
	bodies := parseBodies(lines)
	you := bodies["YOU"]
	san := bodies["SAN"]
	assert.Equal(t, 4, distance(you, san)-2)
}
