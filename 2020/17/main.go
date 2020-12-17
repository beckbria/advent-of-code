package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/17
// Hey, look, it's another Game of Life.  I forget how this was even themed.

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

const (
	active   = '#'
	inactive = '.'
)

type point4 struct {
	W, X, Y, Z int64
}

type grid map[point4]bool

func (g *grid) countActive() int64 {
	a := int64(0)
	for _, v := range *g {
		if v {
			a++
		}
	}
	return a
}

// extremities returns minX, maxX, minY, maxY, minZ, maxZ, minW, maxW
func (g *grid) extremities() (int64, int64, int64, int64, int64, int64, int64, int64) {
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
	for k, v := range *g {
		if v {
			minX = aoc.Min(minX, k.X)
			maxX = aoc.Max(maxX, k.X)
			minY = aoc.Min(minY, k.Y)
			maxY = aoc.Max(maxY, k.Y)
			minZ = aoc.Min(minZ, k.Z)
			maxZ = aoc.Max(maxZ, k.Z)
			minW = aoc.Min(minW, k.W)
			maxW = aoc.Max(maxW, k.W)
		}
	}
	return minX, maxX, minY, maxY, minZ, maxZ, minW, maxW
}

func (g *grid) neighbors(pt *point4) int {
	neighbors := 0
	for x := pt.X - 1; x <= pt.X+1; x++ {
		for y := pt.Y - 1; y <= pt.Y+1; y++ {
			for z := pt.Z - 1; z <= pt.Z+1; z++ {
				for w := pt.W - 1; w <= pt.W+1; w++ {
					if (*g)[point4{X: x, Y: y, Z: z, W: w}] {
						neighbors++
					}
				}
			}
		}
	}
	if (*g)[*pt] {
		neighbors--
	}
	return neighbors
}

func initGrid(lines []string) grid {
	g := make(grid)
	for y, l := range lines {
		for x, c := range []rune(l) {
			g[point4{X: int64(x), Y: int64(y), Z: 0, W: 0}] = (c == active)
		}
	}
	return g
}

func advance3(old grid) grid {
	g := make(grid)
	minX, maxX, minY, maxY, minZ, maxZ, _, _ := old.extremities()

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				here := point4{X: x, Y: y, Z: z, W: 0}
				neighbors := old.neighbors(&here)
				if old[here] {
					g[here] = neighbors == 2 || neighbors == 3
				} else {
					g[here] = neighbors == 3
				}
			}
		}
	}
	return g
}

func advance4(old grid) grid {
	g := make(grid)
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := old.extremities()

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for w := minW - 1; w <= maxW+1; w++ {
					here := point4{X: x, Y: y, Z: z, W: w}
					neighbors := old.neighbors(&here)
					if old[here] {
						g[here] = neighbors == 2 || neighbors == 3
					} else {
						g[here] = neighbors == 3
					}
				}
			}
		}
	}
	return g
}

func step1(lines []string) int64 {
	g := initGrid(lines)
	for i := 0; i < 6; i++ {
		g = advance3(g)
	}
	return g.countActive()
}

func step2(lines []string) int64 {
	g := initGrid(lines)
	for i := 0; i < 6; i++ {
		g = advance4(g)
	}
	return g.countActive()
}
