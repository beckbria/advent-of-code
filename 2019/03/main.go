package main

import (
	"fmt"
	"strconv"
	"strings"

	"../aoc"
)

// https://adventofcode.com/2019/day/3
// Given the path of two wires (that is, how far in a direction they go before turning)
// find the intersection closest to the initial port (not counting the initial port)
// Part 1: By Manhattan Distance
// Part 2: By the total length along the wire you must travel to reach the point

const (
	right = 'R'
	down  = 'D'
	up    = 'U'
	left  = 'L'
)

type path struct {
	direction rune
	distance  int
}

type wire map[aoc.Point]int

func main() {
	lines := aoc.ReadFileLines("input.txt")

	sw := aoc.NewStopwatch()
	optimalManhattan, optimalCost := bestValues(lines[0], lines[1])
	fmt.Println(optimalManhattan)
	fmt.Println(optimalCost)
	fmt.Println(sw.Elapsed())
}

// Returns the optimal manhattan distance and the optimal total cost
func bestValues(first, second string) (int64, int) {
	// Generate the set of points touched by each wire
	firstPath := readPath(first)
	secondPath := readPath(second)
	firstPoints := pathToWire(firstPath)
	secondPoints := pathToWire(secondPath)

	crosses := intersection(firstPoints, secondPoints)
	// Find the shortest Manhattan Distance
	home := aoc.Point{X: 0, Y: 0}
	optimalManhattan := int64(99999999)
	optimalCost := 99999999
	for pt, distance := range crosses {
		optimalManhattan = aoc.Min(optimalManhattan, home.ManhattanDistance(pt))
		optimalCost = aoc.MinInt(optimalCost, distance)
	}
	return optimalManhattan, optimalCost
}

func readPath(line string) []path {
	input := strings.Split(line, ",")
	program := []path{}
	for _, i := range input {
		runes := []rune(i)
		distance, _ := strconv.ParseInt(string(runes[1:]), 10, 32)
		p := path{
			direction: runes[0],
			distance:  int(distance)}
		program = append(program, p)
	}
	return program
}

func pathToWire(trace []path) wire {
	points := make(wire)

	x := 0
	y := 0
	totalDistance := 0
	for _, l := range trace {
		xDelta := 0
		yDelta := 0
		switch l.direction {
		case left:
			xDelta = -1
		case right:
			xDelta = 1
		case up:
			yDelta = -1
		case down:
			yDelta = 1
		}
		for i := 0; i < l.distance; i++ {
			totalDistance++
			x += xDelta
			y += yDelta
			key := aoc.Point{X: int64(x), Y: int64(y)}
			if oldVal, found := points[key]; !found || totalDistance < oldVal {
				points[key] = totalDistance
			}
		}
	}

	return points
}

func intersection(a, b wire) wire {
	i := make(wire)
	for pt, distanceA := range a {
		if distanceB, found := b[pt]; found {
			i[pt] = distanceA + distanceB
		}
	}
	return i
}
