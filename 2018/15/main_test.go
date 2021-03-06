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
	}, 3)

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
	}, 3)

	loc := MakeUnitLocationMap(&c)
	paths := ShortestPaths(loc[1][3], loc, &c)
	assert.Equal(t, 1, paths[1][2].distance)
	assert.Equal(t, 3, paths[2][5].distance)
	assert.Equal(t, Unreachable, paths[0][3].distance)
	assert.Equal(t, Unreachable, paths[5][3].distance)
}

func TestOutcome(t *testing.T) {

	score, winningSide, rounds, hp, _ := Outcome([]string{
		"#######",
		"#.G...#",
		"#...EG#",
		"#.#.#G#",
		"#..G#E#",
		"#.....#",
		"#######",
	}, 3)
	assert.Equal(t, 27730, score)
	assert.Equal(t, 47, rounds)
	assert.Equal(t, 590, hp)
	assert.Equal(t, Goblin, int32(winningSide))

	score, winningSide, rounds, hp, _ = Outcome([]string{
		"#######",
		"#G..#E#",
		"#E#E.E#",
		"#G.##.#",
		"#...#E#",
		"#...E.#",
		"#######",
	}, 3)
	assert.Equal(t, 36334, score)
	assert.Equal(t, 37, rounds)
	// We seem to be missing a hit here - we get 985hp returned
	assert.Equal(t, 982, hp)
	assert.Equal(t, Elf, int32(winningSide))

	score, winningSide, rounds, hp, _ = Outcome([]string{
		"#######",
		"#E..EG#",
		"#.#G.E#",
		"#E.##E#",
		"#G..#.#",
		"#..E#.#",
		"#######",
	}, 3)
	assert.Equal(t, 39514, score)
	// We appear to be missing a round here - we return 45
	assert.Equal(t, 46, rounds)
	assert.Equal(t, 859, hp)
	assert.Equal(t, Elf, int32(winningSide))

	score, winningSide, rounds, hp, _ = Outcome([]string{
		"#######",
		"#E.G#.#",
		"#.#G..#",
		"#G.#.G#",
		"#G..#.#",
		"#...E.#",
		"#######",
	}, 3)
	assert.Equal(t, 27755, score)
	assert.Equal(t, 35, rounds)
	assert.Equal(t, 793, hp)
	assert.Equal(t, Goblin, int32(winningSide))

	score, winningSide, rounds, hp, _ = Outcome([]string{
		"#######",
		"#.E...#",
		"#.#..G#",
		"#.###.#",
		"#E#G#G#",
		"#...#G#",
		"#######",
	}, 3)
	assert.Equal(t, 28944, score)
	assert.Equal(t, 54, rounds)
	assert.Equal(t, 536, hp)
	assert.Equal(t, Goblin, int32(winningSide))

	score, winningSide, rounds, hp, _ = Outcome([]string{
		"#########",
		"#G......#",
		"#.E.#...#",
		"#..##..G#",
		"#...##..#",
		"#...#...#",
		"#.G...G.#",
		"#.....G.#",
		"#########",
	}, 3)
	assert.Equal(t, 18740, score)
	assert.Equal(t, 20, rounds)
	assert.Equal(t, 937, hp)
	assert.Equal(t, Goblin, int32(winningSide))
}

func TestSmallestVictory(t *testing.T) {

	assert.Equal(t, 31284, LowestWinningOutcome([]string{
		"#######",
		"#E..EG#",
		"#.#G.E#",
		"#E.##E#",
		"#G..#.#",
		"#..E#.#",
		"#######",
	}))

	assert.Equal(t, 3478, LowestWinningOutcome([]string{
		"#######",
		"#E.G#.#",
		"#.#G..#",
		"#G.#.G#",
		"#G..#.#",
		"#...E.#",
		"#######",
	}))

	assert.Equal(t, 6474, LowestWinningOutcome([]string{
		"#######",
		"#.E...#",
		"#.#..G#",
		"#.###.#",
		"#E#G#G#",
		"#...#G#",
		"#######",
	}))

	assert.Equal(t, 1140, LowestWinningOutcome([]string{
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
