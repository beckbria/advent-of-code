package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

const debug = false

func main() {
	p := intcode.ReadIntCode("input.txt")

	sw := lib.NewStopwatch()
	// Part 1
	d := newTractorBeamDrone(p)
	fmt.Println(d.part1())
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(d.part2())
	fmt.Println(sw.Elapsed())
}

type drone struct {
	grid map[lib.Point]bool
	c    intcode.Computer
	io   *intcode.StreamInputOutput
}

func newTractorBeamDrone(p intcode.Program) *drone {
	d := drone{
		grid: make(map[lib.Point]bool),
		c:    intcode.NewComputer(p),
		io:   intcode.NewStreamInputOutput([]int64{})}
	d.c.Io = d.io
	return &d
}

const (
	stationary = 0
	moving     = 1
)

func (d *drone) probe(pt *lib.Point) bool {
	if v, found := d.grid[*pt]; found {
		return v
	}
	d.c.Reset()
	d.io.AppendInput(pt.X)
	d.io.AppendInput(pt.Y)
	d.c.RunToNextInput()
	d.c.Step()
	d.c.RunToNextInput()
	d.c.Step()
	d.c.RunToNextInput()
	f := (d.io.LastOutput() == moving)
	d.grid[*pt] = f
	return f
}

func (d *drone) part1() int64 {
	count := int64(0)
	threshold := int64(49)
	for x := int64(0); x <= threshold; x++ {
		for y := int64(0); y <= threshold; y++ {
			pt := lib.Point{X: x, Y: y}
			if d.probe(&pt) {
				count++
			}
		}
	}
	return count
}

type rangee struct {
	start, end int64
}

func r(s, e int64) *rangee {
	r := rangee{start: s, end: e}
	return &r
}

func (d *drone) part2() int64 {
	// Notes: On line 300, the beam is 42 units wide and starts at 219
	// The X shift is 0, 1, 1 usually, which lines up - the slope is somewhat steeper than 2/3
	// As such, we need a width of at least 133, so we should start probing somewhere around line 960

	yStart := int64(950)
	yEnd := yStart + 750
	xStart := int64(665)
	xEnd := xStart + 900

YLOOP:
	for y := yStart; y <= yEnd; y++ {
		foundStart := false
		for x := xStart; x <= xEnd; x++ {
			pt := lib.Point{X: x, Y: y}
			d.probe(&pt)
			r := d.grid[pt]
			if foundStart && !r {
				// We've reached the end of the beam
				continue YLOOP
			} else {
				foundStart = r
			}
		}
		if foundStart {
			fmt.Printf("WARN: Ran of of X space on line %d\n", y)
		}
	}

	home := lib.Point{X: 0, Y: 0}
	best := lib.Point{X: 999999, Y: 999999}

	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			pt := lib.Point{X: x, Y: y}
			if d.square(x, y, int64(100)) && pt.ManhattanDistance(&home) < best.ManhattanDistance(&home) {
				fmt.Println(pt)
				best = pt
			}
		}
	}

	return 10000*best.X + best.Y
}

func (d *drone) square(x, y, size int64) bool {
	for X := x; X < x+size; X++ {
		for Y := y; Y < y+size; Y++ {
			if !d.grid[lib.Point{X: X, Y: Y}] {
				return false
			}
		}
	}
	return true
}

func (d *drone) print() {
	minX := int64(9999999)
	maxX := int64(-9999999)
	minY := int64(9999999)
	maxY := int64(-9999999)
	for pt := range d.grid {
		minX = lib.Min(minX, pt.X)
		maxX = lib.Max(maxX, pt.X)
		minY = lib.Min(minY, pt.Y)
		maxY = lib.Max(maxY, pt.Y)
	}
	fmt.Printf("mX:%d, MX:%d, mY:%d, MY: %d\n", minX, maxX, minY, maxY)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if d.grid[lib.Point{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
