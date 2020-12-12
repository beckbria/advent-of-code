package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// A ship driving around the ocean

func main() {
	moves := readMoves(aoc.ReadFileLines("input.txt"))
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(moves))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(moves))
	fmt.Println(sw.Elapsed())
}

var (
	inputRegex = regexp.MustCompile(`^([A-Z])([0-9]+)$`)
)

type move struct {
	dir       aoc.Direction
	dirString string
	dist      int64
	rotate    bool
}

func readMoves(lines []string) []move {
	m := []move{}
	for _, l := range lines {
		tokens := inputRegex.FindStringSubmatch(l)
		rotate, dir := parseDir(tokens[1])
		dist, _ := strconv.ParseInt(tokens[2], 10, 64)
		m = append(m, move{dir: dir, dist: dist, dirString: tokens[1], rotate: rotate})
	}
	return m
}

func parseDir(dir string) (bool, aoc.Direction) {
	switch dir {
	case "L":
		return true, aoc.Left
	case "R":
		return true, aoc.Right
	case "N":
		return false, aoc.North
	case "S":
		return false, aoc.South
	case "W":
		return false, aoc.West
	case "E":
		return false, aoc.East
	case "F":
		return false, aoc.Forward
	}
	log.Fatalln("Unknown direction")
	return false, aoc.Forward
}

func step1(moves []move) int64 {
	home := aoc.Point{X: 0, Y: 0}
	here := aoc.Point{X: 0, Y: 0}
	facing := aoc.East
	for _, m := range moves {
		if m.rotate {
			facing = rotate(facing, m)
		} else {
			xd, yd := distance(facing, m)
			here.X += xd
			here.Y += yd
		}
	}
	return here.ManhattanDistance(&home)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(facing aoc.Direction, m move) (int64, int64) {
	dir := m.dir
	if dir == aoc.Forward {
		dir = facing
	}
	return dir.DeltaX() * m.dist, dir.DeltaY() * m.dist
}

func rotate(facing aoc.Direction, m move) aoc.Direction {
	amount := m.dist
	if m.dir == aoc.Left {
		for amount > 0 {
			facing = facing.Ccw()
			amount -= 90
		}
	} else {
		for amount > 0 {
			facing = facing.Cw()
			amount -= 90
		}
	}
	return facing
}

func step2(moves []move) int64 {
	home := aoc.Point{X: 0, Y: 0}
	here := aoc.Point{X: 0, Y: 0}
	waypoint := aoc.Point{X: 10, Y: 1}
	for _, m := range moves {
		if m.rotate {
			rotateWaypoint(&waypoint, m)
		} else {
			if m.dir == aoc.Forward {
				here.X += m.dist * waypoint.X
				here.Y += m.dist * waypoint.Y
			} else {
				xd, yd := distance(aoc.North, m)
				waypoint.X += xd
				waypoint.Y -= yd // North = -1 for graphical coordinates
			}
		}
	}
	return here.ManhattanDistance(&home)
}

func rotateWaypoint(w *aoc.Point, m move) {
	amount := m.dist % 360

	if amount == 0 {
		return
	}

	if amount == 180 {
		w.X = -w.X
		w.Y = -w.Y
	}

	if (amount == 90 && m.dir == aoc.Left) || (amount == 270 && m.dir == aoc.Right) {
		x, y := w.X, w.Y
		w.X = -y
		w.Y = x
	} else if (amount == 90 && m.dir == aoc.Right) || (amount == 270 && m.dir == aoc.Left) {
		x, y := w.X, w.Y
		w.X = y
		w.Y = -x
	}

}
