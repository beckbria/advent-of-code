package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacent(t *testing.T) {
	c := ReadCave([]string{
		"#######",
		"#G..#E#",
		"#E#E.E#",
		"#G.##.#",
		"#...#E#",
		"#...E.#",
		"#######",
	})

	loc := MakeUnitLocationMap(&c)
	pt := Point{x: 1, y: 1}
	units, empty := FindAdjacent(&pt, loc, &c)
	assert.Equal(t, []Point{Point{x: 2, y: 1}}, empty)
	assert.Equal(t, 1, len(units))
}

func TestFindShortest(t *testing.T) {
	c := ReadCave([]string{
		"#######",
		"#G..#E#",
		"#E#E.E#",
		"#G.##.#",
		"#...#E#",
		"#...E.#",
		"#######",
	})

	loc := MakeUnitLocationMap(&c)
	paths := ShortestPaths(loc[1][3], loc, &c)
	assert.Equal(t, 1, paths[1][2].distance)
	assert.Equal(t, 3, paths[2][5].distance)
	assert.Equal(t, Unreachable, paths[0][3].distance)
	assert.Equal(t, Unreachable, paths[5][3].distance)
}

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

	/*assert.Equal(t, 39514, Outcome([]string{
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
	}))*/
}
