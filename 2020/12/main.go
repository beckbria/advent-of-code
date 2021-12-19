package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/#
// A ship driving around the ocean

func main() {
	moves := readMoves(lib.ReadFileLines("2020/12/input.txt"))
	sw := lib.NewStopwatch()
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
	dir       lib.Direction
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

func parseDir(dir string) (bool, lib.Direction) {
	switch dir {
	case "L":
		return true, lib.Left
	case "R":
		return true, lib.Right
	case "N":
		return false, lib.North
	case "S":
		return false, lib.South
	case "W":
		return false, lib.West
	case "E":
		return false, lib.East
	case "F":
		return false, lib.Forward
	}
	log.Fatalln("Unknown direction")
	return false, lib.Forward
}

func step1(moves []move) int64 {
	home := lib.Point{X: 0, Y: 0}
	here := lib.Point{X: 0, Y: 0}
	facing := lib.East
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

func distance(facing lib.Direction, m move) (int64, int64) {
	dir := m.dir
	if dir == lib.Forward {
		dir = facing
	}
	return dir.DeltaX() * m.dist, dir.DeltaY() * m.dist
}

func rotate(facing lib.Direction, m move) lib.Direction {
	amount := m.dist
	if m.dir == lib.Left {
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
	home := lib.Point{X: 0, Y: 0}
	here := lib.Point{X: 0, Y: 0}
	waypoint := lib.Point{X: 10, Y: 1}
	for _, m := range moves {
		if m.rotate {
			rotateWaypoint(&waypoint, m)
		} else {
			if m.dir == lib.Forward {
				here.X += m.dist * waypoint.X
				here.Y += m.dist * waypoint.Y
			} else {
				xd, yd := distance(lib.North, m)
				waypoint.X += xd
				waypoint.Y -= yd // North = -1 for graphical coordinates
			}
		}
	}
	return here.ManhattanDistance(&home)
}

func rotateWaypoint(w *lib.Point, m move) {
	amount := m.dist % 360

	if amount == 0 {
		return
	}

	if amount == 180 {
		w.X = -w.X
		w.Y = -w.Y
	}

	if (amount == 90 && m.dir == lib.Left) || (amount == 270 && m.dir == lib.Right) {
		x, y := w.X, w.Y
		w.X = -y
		w.Y = x
	} else if (amount == 90 && m.dir == lib.Right) || (amount == 270 && m.dir == lib.Left) {
		x, y := w.X, w.Y
		w.X = y
		w.Y = -x
	}

}
