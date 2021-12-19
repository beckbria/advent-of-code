package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

type color int

const (
	black = 0
	white = 1
)

// https://adventofcode.com/2019/day/11
// Run an intcode program that moves a robot around.  Keep track
// of where it goes and what it does

func main() {
	program := intcode.ReadIntCode("2019/11/input.txt")
	sw := lib.NewStopwatch()

	// Part 1
	painted := paint(program, black)
	fmt.Println(len(painted))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	painted = paint(program, white)
	draw(painted)
	fmt.Println(sw.Elapsed())
}

func paint(program intcode.Program, startColor color) map[lib.Point]color {
	c := intcode.NewComputer(program)
	io := intcode.NewStreamInputOutput([]int64{})
	c.Io = io
	grid := make(map[lib.Point]color, 0)
	grid[lib.Point{X: 0, Y: 0}] = startColor

	x := int64(0)
	y := int64(0)
	dir := lib.Up

	for c.IsRunning() {
		// Input the current panel.  All panels start black, and the default map
		// read is the 0 value, which conveniently corresponds to black
		current := grid[lib.Point{X: x, Y: y}]
		io.AppendInput(int64(current))
		previousOutputCount := len(io.Outputs)
		// Run until two values are output
		for c.IsRunning() && (len(io.Outputs) < (previousOutputCount + 2)) {
			c.Step()
		}

		if !c.IsRunning() {
			break
		}

		newColor := color(io.Outputs[len(io.Outputs)-2])
		turnRight := io.LastOutput()
		grid[lib.Point{X: x, Y: y}] = newColor
		if turnRight == 1 {
			dir = dir.Cw()
		} else {
			dir = dir.Ccw()
		}
		x += dir.DeltaX()
		y += dir.DeltaY()
	}

	return grid
}

func draw(grid map[lib.Point]color) {
	minX := int64(0)
	minY := int64(0)
	maxX := int64(0)
	maxY := int64(0)

	for pt := range grid {
		minX = lib.Min(minX, pt.X)
		maxX = lib.Max(maxX, pt.X)
		minY = lib.Min(minY, pt.Y)
		maxY = lib.Max(maxY, pt.Y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c := grid[lib.Point{X: x, Y: y}]
			if c == black {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
