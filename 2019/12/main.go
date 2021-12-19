package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/beckbria/advent-of-code/2019/lib"
)

var origin = lib.Point3{X: 0, Y: 0, Z: 0}

type moon struct {
	p lib.Point3 // Position
	v lib.Point3 // Velocity
}

func (m *moon) energy() int64 {
	// Potential energy is equal to the absolute value of each coordinate of the position
	p := m.p.ManhattanDistance(&origin)
	// Kinetic energy is equal to the absolute value of each coordinate of the velocity
	k := m.v.ManhattanDistance(&origin)
	return p * k
}

func (m *moon) gravitate(n *moon) {
	dm, dn := gravitate1D(m.p.X, n.p.X)
	m.v.X += dm
	n.v.X += dn
	dm, dn = gravitate1D(m.p.Y, n.p.Y)
	m.v.Y += dm
	n.v.Y += dn
	dm, dn = gravitate1D(m.p.Z, n.p.Z)
	m.v.Z += dm
	n.v.Z += dn
}

// Return a pair of the X position and velocity for use as a map key
func (m *moon) x() lib.Point {
	return lib.Point{X: m.p.X, Y: m.v.X}
}

// Return a pair of the Y position and velocity for use as a map key
func (m *moon) y() lib.Point {
	return lib.Point{X: m.p.Y, Y: m.v.Y}
}

// Return a pair of the Z position and velocity for use as a map key
func (m *moon) z() lib.Point {
	return lib.Point{X: m.p.Z, Y: m.v.Z}
}

// Very simplistic gravity model.  If two bodies are in the same plane they don't move,
// otherwise they move one unit towards each other
func gravitate1D(i, j int64) (int64, int64) {
	if i < j {
		return 1, -1
	} else if i > j {
		return -1, 1
	} else {
		return 0, 0
	}
}

func (m *moon) move() {
	m.p.X += m.v.X
	m.p.Y += m.v.Y
	m.p.Z += m.v.Z
}

type moons []moon

// Total energy of system
func (ms moons) energy() int64 {
	e := int64(0)
	for _, m := range ms {
		e += m.energy()
	}
	return e
}

// Return a pair of the X position and velocity for use as a map key
func (ms moons) x() [4]lib.Point {
	pts := [4]lib.Point{}
	for i := range ms {
		pts[i] = ms[i].x()
	}
	return pts
}

// Return a pair of the Y position and velocity for use as a map key
func (ms moons) y() [4]lib.Point {
	pts := [4]lib.Point{}
	for i := range ms {
		pts[i] = ms[i].y()
	}
	return pts
}

// Return a pair of the Z position and velocity for use as a map key
func (ms moons) z() [4]lib.Point {
	pts := [4]lib.Point{}
	for i := range ms {
		pts[i] = ms[i].z()
	}
	return pts
}

// Advances one time unit
func (ms moons) step() {
	for i := range ms {
		for j := i + 1; j < len(ms); j++ {
			ms[i].gravitate(&ms[j])
		}
	}

	for i := range ms {
		ms[i].move()
	}
}

func main() {
	input := lib.ReadFileLines("2019/12/input.txt")
	sw := lib.NewStopwatch()
	// Part 1
	m := readMoons(input)
	for i := 0; i < 1000; i++ {
		m.step()
	}
	fmt.Println(m.energy())
	fmt.Println(sw.Elapsed())
	sw.Reset()

	// Part 2
	m = readMoons(input)
	fmt.Println(firstCycle(m))
	fmt.Println(sw.Elapsed())
}

var (
	inputRegEx = regexp.MustCompile(`^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$`)
)

func readMoons(input []string) moons {
	m := moons{}
	for _, s := range input {
		tokens := inputRegEx.FindStringSubmatch(s)
		x, err := strconv.Atoi(tokens[1])
		lib.Check(err)
		y, err := strconv.Atoi(tokens[2])
		lib.Check(err)
		z, err := strconv.Atoi(tokens[3])
		m = append(m, moon{p: lib.Point3{X: int64(x), Y: int64(y), Z: int64(z)}, v: origin})
	}
	return m
}

const (
	xAxis   = 0
	yAxis   = 1
	zAxis   = 2
	numAxes = 3
)

func firstCycle(ms moons) int64 {
	// Find the period for each moon individually, then figure out when they'll all line up
	seen := make([]map[[4]lib.Point]int64, numAxes)
	start := make([]int64, numAxes)
	period := make([]int64, numAxes)

	for i := 0; i < numAxes; i++ {
		seen[i] = make(map[[4]lib.Point]int64)
	}
	seen[xAxis][ms.x()] = 0
	seen[yAxis][ms.y()] = 0
	seen[zAxis][ms.z()] = 0

	for t := int64(1); period[xAxis] < 1 || period[yAxis] < 1 || period[zAxis] < 1; t++ {
		ms.step()
		if period[xAxis] < 1 {
			x := ms.x()
			if first, found := seen[xAxis][x]; found {
				start[xAxis] = first
				period[xAxis] = t - first
			} else {
				seen[xAxis][x] = t
			}
		}

		if period[yAxis] < 1 {
			y := ms.y()
			if first, found := seen[yAxis][y]; found {
				start[yAxis] = first
				period[yAxis] = t - first
			} else {
				seen[yAxis][y] = t
			}
		}

		if period[zAxis] < 1 {
			z := ms.z()
			if first, found := seen[zAxis][z]; found {
				start[zAxis] = first
				period[zAxis] = t - first
			} else {
				seen[zAxis][z] = t
			}
		}
	}

	//fmt.Printf("https://www.wolframalpha.com/input/?i=least+common+multiple+of+%d%%2C+%d%%2C+and+%d\n", period[0], period[1], period[2])
	return lcm3(period[0], period[1], period[2])
}

// Returns the least common multiple of 3 numbers
func lcm3(x, y, z int64) int64 {
	return lib.Lcm(x, lib.Lcm(y, z))
}
