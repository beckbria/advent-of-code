package main

import (
	"fmt"
	"strconv"
	"strings"

	"../aoc"
)

const (
	right = 'R'
	down = 'D'
	up = 'U'
	left = 'L'
)

type path struct {
	direction rune
	distance int
}

type wire map[aoc.Point]int

func main() {
	lines := aoc.ReadFileLines("input.txt")
	firstPath := readWire(lines[0])
	secondPath := readWire(lines[1])

	sw := aoc.NewStopwatch()	
	firstPoints := findPoints(firstPath)
	secondPoints := findPoints(secondPath)
	crosses := intersection(firstPoints, secondPoints)
	// Find the shortest Manhattan Distance
	home := aoc.Point{X:0, Y:0}
	bestPart1 := aoc.Point{X: 999999, Y: 999999}
	bestPart2 := 99999999
	for pt, distance := range crosses {
		if home.ManhattanDistance(pt) < home.ManhattanDistance(bestPart1) {
			bestPart1 = pt
		}
		bestPart2 = aoc.MinInt(bestPart2, distance)
	}
	fmt.Println(home.ManhattanDistance(bestPart1))
	fmt.Println(bestPart2)
	fmt.Println(sw.Elapsed())
}

func readWire(line string) []path {
	input := strings.Split(line, ",")
	program := []path{}
	for _, i := range input {
		runes := []rune(i)
		distance, _ := strconv.ParseInt(string(runes[1:]), 10, 32) 
		p := path{
			direction: runes[0],
			distance: int(distance)}
		program = append(program, p)
	}
	return program
}

func findPoints(trace []path) wire{
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
			key := aoc.Point{X:int64(x), Y:int64(y)}
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