package main

import (
	"fmt"
	"math"
	"sort"

	"../aoc"
)

// https://adventofcode.com/2019/day/10

const (
	debug        = false
	empty        = '.'
	asteroid     = '#'
	notDestroyed = -1
)

type starMap [][]rune

// asteroidSet is a hash set of asteroid locations.  The key is the location.
// The value can be used to store the destruction order
type asteroidSet map[aoc.Point]int

func (a asteroidSet) put(x, y, order int) {
	a[aoc.Point{X: int64(x), Y: int64(y)}] = order
}

func (a asteroidSet) canSee(from, to *aoc.Point) bool {
	if from.Equals(to) {
		return false
	}

	// We need to find exact integer factors of the delta to find blocking patterns
	slope := from.SlopeTo(to)

	for target := range a {
		if target.Equals(to) || target.Equals(from) {
			continue
		}
		s := from.SlopeTo(&target)
		if aoc.SameSlope(&s, &slope) && from.ManhattanDistance(&target) < from.ManhattanDistance(to) {
			return false
		}
	}

	return true
}

func main() {
	input := aoc.ReadFileLines("input.txt")
	m := newMap(input)
	sw := aoc.NewStopwatch()
	pt, count, _ := bestMonitoringStation(m)
	// Part 1
	fmt.Println(count)
	fmt.Println(sw.Elapsed())
	sw.Reset()
	// Part 2
	a := findByDestructionOrder(m, 200, &pt)
	fmt.Println(a.X*100 + a.Y)
	fmt.Println(sw.Elapsed())
}

func newMap(input []string) starMap {
	m := make(starMap, 0)
	for _, l := range input {
		m = append(m, []rune(l))
	}
	return m
}

// bestMonitoringStation finds the best position for a monitoring
// station.  It returns its x coordinate, y coordinate, and the number
// of asteroids which can be seen
func bestMonitoringStation(m starMap) (aoc.Point, int, map[aoc.Point]int) {
	bestCount := 0
	bestLocation := aoc.Point{X: -1, Y: -1}
	a := findAsteroids(m)
	counts := make(map[aoc.Point]int)

	for monitor := range a {
		count := 0
		for target := range a {
			canSee := a.canSee(&monitor, &target)
			if canSee {
				count++
			}
		}
		counts[monitor] = count
		if count > bestCount {
			bestCount = count
			bestLocation = monitor
		}
	}
	return bestLocation, bestCount, counts
}

func findAsteroids(m starMap) asteroidSet {
	a := make(asteroidSet)
	for y, row := range m {
		for x, c := range row {
			if c == asteroid {
				a.put(x, y, notDestroyed)
			}
		}
	}
	return a
}

func findByDestructionOrder(m starMap, desiredOrder int, from *aoc.Point) aoc.Point {
	// Compute the angle in degrees of each star and go around and around, marking them as destroyed
	do := destructionOrder(m, from)
	for a, order := range do {
		if order == desiredOrder {
			return a
		}
	}
	return aoc.Point{X: -1, Y: -1}
}

func destructionOrder(m starMap, from *aoc.Point) asteroidSet {
	order := findAsteroids(m)

	angles := []float64{}
	// Maps from angle to a slice of points.  These slices will later be sorted by distance from the laser
	pointsByAngle := make(map[float64][]aoc.Point, 0)

	for a := range order {
		if from.Equals(&a) {
			continue
		}
		slope := from.SlopeTo(&a)
		// negative y is actually up
		slope.Numerator = -slope.Numerator
		angle := aoc.SlopeToRadians(&slope)
		// We want to start at the vertical and go clockwise, so play with the numbers a bit
		for angle <= aoc.PiOver2 {
			angle += 2 * math.Pi
		}
		angle = -angle

		if debug {
			fmt.Printf("[%d,%d]->[%d,%d] = %d/%d = %f\n", from.X, from.Y, a.X, a.Y, slope.Numerator, slope.Denominator, angle)
		}

		if _, found := pointsByAngle[angle]; !found {
			pointsByAngle[angle] = make([]aoc.Point, 0)
			angles = append(angles, angle)
		}
		pointsByAngle[angle] = append(pointsByAngle[angle], a)
	}

	// Sort the angles
	sort.Float64s(angles)
	if debug {
		fmt.Println(angles)
	}
	// Sort the points for each angle
	for _, pts := range pointsByAngle {
		sort.Slice(pts, func(i, j int) bool {
			return from.ManhattanDistance(&(pts[i])) < from.ManhattanDistance(&(pts[j]))
		})
	}

	next := 1
	for shotSomething := true; shotSomething; {
		shotSomething = false
		for _, ang := range angles {
			// Go through the asteroids at this angle and shoot the next one
			for _, pt := range pointsByAngle[ang] {
				if order[pt] != notDestroyed {
					continue
				}
				order[pt] = next
				next++
				shotSomething = true
				break
			}
		}
	}

	return order
}
