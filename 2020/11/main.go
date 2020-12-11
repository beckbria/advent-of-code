package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// The game of life simulator we all expected

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	g := newGame(lines)
	fmt.Println(step1(g))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	g = newGame(lines)
	fmt.Println(step2(g))
	fmt.Println(sw.Elapsed())
}

const (
	empty    = rune('L')
	occupied = rune('#')
	floor    = rune('.')
)

type game struct {
	grid [][]rune
}

// Advance and return the layout
func (g *game) advance1() *game {
	prev := g.clone()

	for ridx, r := range prev.grid {
		for idx, was := range r {
			o := prev.adjacent(ridx, idx)
			if was == empty && o == 0 {
				g.grid[ridx][idx] = occupied
			} else if was == occupied && o > 3 {
				g.grid[ridx][idx] = empty
			}
		}
	}

	return prev
}

// Advance and return the layout
func (g *game) advance2() *game {
	prev := g.clone()

	for ridx, r := range prev.grid {
		for idx, was := range r {
			o := prev.sightlines(ridx, idx)
			if was == empty && o == 0 {
				g.grid[ridx][idx] = occupied
			} else if was == occupied && o > 4 {
				g.grid[ridx][idx] = empty
			}
		}
	}

	return prev
}

// adjacent returns the number of adjacent occupied seats
func (g *game) adjacent(row, col int) int {
	o := 0
	for r := row - 1; r <= row+1; r++ {
		if r < 0 || r >= len(g.grid) {
			continue
		}
		for c := col - 1; c <= col+1; c++ {
			if c < 0 || c >= len(g.grid[0]) || (r == row && c == col) {
				continue
			}
			if g.grid[r][c] == occupied {
				o++
			}
		}
	}
	return o
}

var directions = []aoc.Point{
	aoc.Point{X: -1, Y: -1},
	aoc.Point{X: -1, Y: 0},
	aoc.Point{X: -1, Y: 1},
	aoc.Point{X: 0, Y: -1},
	aoc.Point{X: 0, Y: 1},
	aoc.Point{X: 1, Y: -1},
	aoc.Point{X: 1, Y: 0},
	aoc.Point{X: 1, Y: 1},
}

func (g *game) sightlines(row, col int) int {
	o := 0
OUTER:
	for _, d := range directions {
		dx, dy := int(d.X), int(d.Y)
		for x, y := col+dx, row+dy; y >= 0 && y < len(g.grid) && x >= 0 && x < len(g.grid[0]); x, y = x+dx, y+dy {
			seen := g.grid[y][x]
			if seen == empty {
				continue OUTER
			} else if seen == occupied {
				o++
				continue OUTER
			}
		}
	}
	return o
}

func (g *game) equals(g2 *game) bool {
	if len(g.grid) != len(g2.grid) {
		return false
	}
	for ridx := range g.grid {
		if len(g.grid[ridx]) != len(g2.grid[ridx]) {
			return false
		}
		for idx, c := range g.grid[ridx] {
			if c != g2.grid[ridx][idx] {
				return false
			}
		}
	}
	return true
}

func (g *game) occupied() int {
	o := 0
	for _, r := range g.grid {
		for _, s := range r {
			if s == occupied {
				o++
			}
		}
	}
	return o
}

func (g *game) clone() *game {
	g2 := game{}
	g2.grid = make([][]rune, len(g.grid))
	for ridx, r := range g.grid {
		g2.grid[ridx] = make([]rune, len(r))
		copy(g2.grid[ridx], r)
	}
	return &g2
}

func (g *game) print() {
	for _, r := range g.grid {
		for _, c := range r {
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
}

func newGame(rows []string) *game {
	g := game{}
	g.grid = make([][]rune, len(rows))
	for ridx, r := range rows {
		g.grid[ridx] = make([]rune, len(r))
		for idx, c := range []rune(r) {
			g.grid[ridx][idx] = c
		}
	}
	return &g
}

func step1(g *game) int {
	for {
		prev := g.advance1()
		if prev.equals(g) {
			return g.occupied()
		}
	}
}

func step2(g *game) int {
	for {
		prev := g.advance2()
		if prev.equals(g) {
			return g.occupied()
		}
	}
}
