package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initial = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

var round1 = []string{
	"#.##.##.##",
	"#######.##",
	"#.#.#..#..",
	"####.##.##",
	"#.##.##.##",
	"#.#####.##",
	"..#.#.....",
	"##########",
	"#.######.#",
	"#.#####.##",
}

var round2 = []string{
	"#.LL.L#.##",
	"#LLLLLL.L#",
	"L.L.L..L..",
	"#LLL.LL.L#",
	"#.LL.LL.LL",
	"#.LLLL#.##",
	"..L.L.....",
	"#LLLLLLLL#",
	"#.LLLLLL.L",
	"#.#LLLL.##",
}

var round3 = []string{
	"#.##.L#.##",
	"#L###LL.L#",
	"L.#.#..#..",
	"#L##.##.L#",
	"#.##.LL.LL",
	"#.###L#.##",
	"..#.#.....",
	"#L######L#",
	"#.LL###L.L",
	"#.#L###.##",
}

var round4 = []string{
	"#.#L.L#.##",
	"#LLL#LL.L#",
	"L.L.L..#..",
	"#LLL.##.L#",
	"#.LL.LL.LL",
	"#.LL#L#.##",
	"..L.L.....",
	"#L#LLLL#L#",
	"#.LLLLLL.L",
	"#.#L#L#.##",
}

var round5 = []string{
	"#.#L.L#.##",
	"#LLL#LL.L#",
	"L.#.L..#..",
	"#L##.##.L#",
	"#.#L.LL.LL",
	"#.#L#L#.##",
	"..L.L.....",
	"#L#L##L#L#",
	"#.LLLLLL.L",
	"#.#L#L#.##",
}

var round6 = []string{
	"#.#L.L#.##",
	"#LLL#LL.L#",
	"L.#.L..#..",
	"#L##.##.L#",
	"#.#L.LL.LL",
	"#.#L#L#.##",
	"..L.L.....",
	"#L#L##L#L#",
	"#.LLLLLL.L",
	"#.#L#L#.##",
}

func TestAdvance(t *testing.T) {
	g := newGame(initial)
	g.advance1()
	assert.True(t, g.equals(newGame(round1)))
	g.advance1()
	assert.True(t, g.equals(newGame(round2)))
	g.advance1()
	assert.True(t, g.equals(newGame(round3)))
	g.advance1()
	assert.True(t, g.equals(newGame(round4)))
	g.advance1()
	assert.True(t, g.equals(newGame(round5)))
	g.advance1()
	assert.True(t, g.equals(newGame(round6)))
}

func TestStep1(t *testing.T) {
	g := newGame(initial)
	assert.Equal(t, 37, step1(g))
}

func TestStep2(t *testing.T) {
	g := newGame(initial)
	assert.Equal(t, 26, step2(g))
}
