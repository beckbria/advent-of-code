package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	yard := ReadYard([]string{
		".#.#...|#.",
		".....#|##|",
		".|..|...#.",
		"..|#.....#",
		"#.#|||#|#|",
		"...#.||...",
		".|....|...",
		"||...#|.#|",
		"|.||||..|.",
		"...#.|..|."})
	AdvanceTime(&yard, 10)
	assert.Equal(t, 1147, ResourceValue(&yard))
}
