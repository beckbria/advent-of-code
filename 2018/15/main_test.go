package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutcome(t *testing.T) {
	assert.Equal(t, 36334, Outcome([]string{
		"#######",
		"#G..#E#",
		"#E#E.E#",
		"#G.##.#",
		"#...#E#",
		"#...E.#",
		"#######",
	}))

	assert.Equal(t, 39514, Outcome([]string{
		"#######",
		"#E..EG#",
		"#.#G.E#",
		"#E.##E#",
		"#G..#.#",
		"#..E#.#",
		"#######",
	}))

	assert.Equal(t, 27755, Outcome([]string{
		"#######",
		"#E.G#.#",
		"#.#G..#",
		"#G.#.G#",
		"#G..#.#",
		"#...E.#",
		"#######",
	}))

	assert.Equal(t, 28944, Outcome([]string{
		"#######",
		"#.E...#",
		"#.#..G#",
		"#.###.#",
		"#E#G#G#",
		"#...#G#",
		"#######",
	}))

	assert.Equal(t, 18740, Outcome([]string{
		"#########",
		"#G......#",
		"#.E.#...#",
		"#..##..G#",
		"#...##..#",
		"#...#...#",
		"#.G...G.#",
		"#.....G.#",
		"#########",
	}))
}
