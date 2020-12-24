package main

import (
	"fmt"
	"log"

	"../../aoc"
)

// https://adventofcode.com/2020/day/24
// Hex-grid pathfinding that becomes the game of life again, plus parsing without delimiters

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	inst := parseInstructions(lines)
	s1, grid := step1(inst)
	fmt.Println(s1)
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(grid))
	fmt.Println(sw.Elapsed())
}

func step1(inst [][]string) (int64, hexGrid) {
	grid := make(hexGrid)
	for _, i := range inst {
		end := followPath(i)
		grid[end] = !grid[end]
	}

	return grid.black(), grid
}

func step2(grid hexGrid) int64 {
	for day := 0; day < 100; day++ {
		grid = grid.advance()
	}

	return grid.black()
}

func parseInstructions(lines []string) [][]string {
	var inst [][]string
	for _, l := range lines {
		inst = append(inst, parseInst(l))
	}
	return inst
}

func parseInst(s string) []string {
	var inst []string
	for len(s) > 0 {
		switch s[0] {
		case 'n', 's':
			inst = append(inst, s[:2])
			s = s[2:]
		case 'e', 'w':
			inst = append(inst, s[:1])
			s = s[1:]
		}
	}
	return inst
}

type hex aoc.Point3

const (
	e  = "e"
	se = "se"
	sw = "sw"
	w  = "w"
	nw = "nw"
	ne = "ne"
)

func (h *hex) shift(s string) {
	// cube coordinates: https://www.redblobgames.com/grids/hexagons/
	switch s {
	case e:
		h.X++
		h.Y--
	case se:
		h.Z++
		h.Y--
	case ne:
		h.X++
		h.Z--
	case w:
		h.X--
		h.Y++
	case nw:
		h.Z--
		h.Y++
	case sw:
		h.X--
		h.Z++
	default:
		log.Fatalf("Unexpected token: %s", s)
	}
}

type hexGrid map[hex]bool

func (g hexGrid) black() int64 {
	count := int64(0)
	for _, v := range g {
		if v {
			count++
		}
	}
	return count
}

func (g hexGrid) pruneEmpty() {
	empty := make(hexGrid)
	for k, v := range g {
		if !v {
			empty[k] = true
		}
	}
	for k := range empty {
		delete(g, k)
	}
}

func (g hexGrid) advance() hexGrid {
	g.pruneEmpty()
	candidates := make(hexGrid)
	for k := range g {
		candidates[k] = true
		// Add neighbors
		for _, dir := range []string{e, ne, nw, w, sw, se} {
			h := k
			h.shift(dir)
			candidates[h] = true
		}
	}

	becoming := make(hexGrid)
	for here := range candidates {
		wasBlack := g[here]
		neighbors := 0
		for _, dir := range []string{e, ne, nw, w, sw, se} {
			h := here
			h.shift(dir)
			if g[h] {
				neighbors++
			}
		}
		becoming[here] = g[here]
		if wasBlack && (neighbors == 0 || neighbors > 2) {
			becoming[here] = false
		} else if !wasBlack && neighbors == 2 {
			becoming[here] = true
		}
	}

	return becoming
}

func followPath(inst []string) hex {
	h := hex{X: 0, Y: 0, Z: 0}
	for _, i := range inst {
		h.shift(i)
	}
	return h
}
