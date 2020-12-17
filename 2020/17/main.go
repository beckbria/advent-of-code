package main

import (
	"fmt"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// TODO: Description

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

type Point4 struct {
	W, X, Y, Z int64
}

type grid map[Point4]bool

func (g *grid) countActive() int64 {
	a := int64(0)
	for _, v := range *g {
		if v {
			a++
		}
	}
	return a
}

// extremities returns minX, maxX, minY, maxY, minZ, maxZ
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
	if debug {
		fmt.Printf("minX, maxX, minY, maxY, minZ, maxZ, minW, maxW = %d,%d,%d,%d,%d,%d,%d,%d\n", minX, maxX, minY, maxY, minZ, maxZ, minW, maxW)
	}
	return minX, maxX, minY, maxY, minZ, maxZ, minW, maxW
}

func (g *grid) print3() {
	minX, maxX, minY, maxY, minZ, maxZ, _, _ := g.extremities()
	for z := minZ; z <= maxZ; z++ {
		fmt.Printf("z=%d\n", z)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if (*g)[Point4{X: x, Y: y, Z: z, W: 0}] {
					fmt.Print(string(active))
				} else {
					fmt.Print(string(inactive))
				}
				fmt.Print()
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func initGrid(lines []string) grid {
	g := make(grid)
	for y, l := range lines {
		for x, c := range []rune(l) {
			g[Point4{X: int64(x), Y: int64(y), Z: 0, W: 0}] = (c == active)
		}
	}
	return g
}

const debug = false

func advance3(old grid) grid {
	g := make(grid)
	minX, maxX, minY, maxY, minZ, maxZ, _, _ := old.extremities()

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				here := Point4{X: x, Y: y, Z: z, W: 0}
				neighbors := 0
				for nx := x - 1; nx <= x+1; nx++ {
					for ny := y - 1; ny <= y+1; ny++ {
						for nz := z - 1; nz <= z+1; nz++ {
							if old[Point4{X: nx, Y: ny, Z: nz, W: 0}] {
								neighbors++
							}
						}
					}
				}
				neighbors = g.neighbors(&here)
				if old[here] {
					//neighbors--
					g[here] = neighbors == 2 || neighbors == 3
				} else {
					g[here] = neighbors == 3
				}
				if debug {
					fmt.Printf("Point {%d,%d,%d} has %d neighbors\n", x, y, z, neighbors)
				}
			}
		}
	}
	return g
}

func (g *grid) neighbors(pt *Point4) int {
	neighbors := 0
	for nx := pt.X - 1; nx <= pt.X+1; nx++ {
		for ny := pt.Y - 1; ny <= pt.Y+1; ny++ {
			for nz := pt.Z - 1; nz <= pt.Z+1; nz++ {
				for nw := pt.W - 1; nw <= pt.W+1; nw++ {
					if (*g)[Point4{X: nx, Y: ny, Z: nz, W: nw}] {
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

func advance4(old grid) grid {
	g := make(grid)
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := old.extremities()

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for w := minW - 1; w <= maxW+1; w++ {
					here := Point4{X: x, Y: y, Z: z, W: w}
					neighbors := 0
					for nx := x - 1; nx <= x+1; nx++ {
						for ny := y - 1; ny <= y+1; ny++ {
							for nz := z - 1; nz <= z+1; nz++ {
								for nw := w - 1; nw <= w+1; nw++ {
									if old[Point4{X: nx, Y: ny, Z: nz, W: nw}] {
										neighbors++
									}
								}
							}
						}
					}
					if old[here] {
						neighbors--
						g[here] = neighbors == 2 || neighbors == 3
					} else {
						g[here] = neighbors == 3
					}
					if debug {
						fmt.Printf("Point {%d,%d,%d} has %d neighbors\n", x, y, z, neighbors)
					}
				}
			}
		}
	}
	return g
}

func step1(lines []string) int64 {
	g := initGrid(lines)
	if debug {
		g.print3()
	}
	for i := 0; i < 6; i++ {
		g = advance3(g)
		if debug {
			fmt.Printf("Turn %d\n", i+1)
			g.print3()
		}
	}
	return g.countActive()
}

func step2(lines []string) int64 {
	g := initGrid(lines)
	for i := 0; i < 6; i++ {
		g = advance4(g)
		if debug {
			fmt.Printf("Turn %d\n", i+1)
			//g.print()
		}
	}
	return g.countActive()
}
