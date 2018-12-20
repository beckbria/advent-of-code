package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

// Component represents a component of the base.  It's an enum
type Component rune

// The components
const (
	Wall    = '#'
	Room    = '.'
	HDoor   = '|'
	VDoor   = '-'
	Unknown = '?' // A door or wall - but we don't know which yet
)

// Location represents an item on the map along with the distance to it from the origin
type Location struct {
	c    Component
	dist int // Distance from origin by shortest path
}

// Unreachable is a placeholder for an unreachable location
const Unreachable = math.MaxInt32

func makeLocation(c Component) Location {
	return Location{c: c, dist: Unreachable}
}

// RoomMap represents a map of the rooms in the base.  map[X][Y]Component
// The origin is always at (0,0), so X and Y can (and probably will) go negative.RoomMap
// Rooms are at even coordinates; doors or walls are at odd coordinates.
type RoomMap map[int]map[int]Location

func (m *RoomMap) resetDistance() {
	for x, yr := range *m {
		for y := range yr {
			(*m)[x][y] = Location{c: (*m)[x][y].c, dist: Unreachable}
		}
	}
}

func (m *RoomMap) set(x, y int, c Component) {
	if _, present := (*m)[x]; !present {
		(*m)[x] = make(map[int]Location)
	}
	(*m)[x][y] = makeLocation(c)
}

// Direction is an enum representing all of the options in the instruction pattern.
type Direction rune

// The directions
const (
	West         = 'W'
	East         = 'E'
	North        = 'N'
	South        = 'S'
	BeginGroup   = '('
	EndGroup     = ')'
	GroupOption  = '|'
	BeginPattern = '^'
	EndPattern   = '$'
)

// Pattern represents a set of paths.  If there are multiple options, then each will
// appear in the outer slice.  If a pattern consists of multiple pieces sequentially,
// each piece will appear in the inner slice.
//
// Thus, (WN|ES) becomes
// Pattern {
// 	[0][0] = "WN"
// 	[1][0] = "ES"
// }
//
// And WW(NN) becomes
// Pattern {
//  [0][0] = "WW"
//  [0][1] = "NN"
// }
type Pattern struct {
	paths   [][]Pattern
	pattern string
}

func (p *Pattern) isCompound() bool {
	return len(p.paths) > 0
}

type point struct {
	x, y int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadPattern turns the string RegEx into a pattern state machine
func ReadPattern(s string) Pattern {
	var p Pattern
	return p
}

// SinglePattern returns a pattern which matches a single string
func SinglePattern(s string) Pattern {
	return Pattern{pattern: s}
}

// ChoicePattern returns a pattern which matches exactly one of a series of patterns
func ChoicePattern(p ...Pattern) Pattern {
	paths := make([][]Pattern, 0)
	for _, path := range p {
		paths = append(paths, []Pattern{path})
	}
	return Pattern{paths: paths}
}

// ConcatenatePattern returns a pattern which matches a series of patterns sequentially
func ConcatenatePattern(p ...Pattern) Pattern {
	return Pattern{paths: [][]Pattern{p}}
}

// BuildMap generates a map from a pattern
func BuildMap(p *Pattern) RoomMap {
	m := make(RoomMap)
	return m
}

// MostDoors returns the most doors you must pass through to reach any room of the house along
// the shortest path
func MostDoors(s string) int {
	p := ReadPattern(s)
	m := BuildMap(&p)
	m.findShortestDistances(point{x: 0, y: 0})

	maxDoors := 0
	for _, yr := range m {
		for _, c := range yr {
			if c.c == Room {
				maxDoors = max(maxDoors, c.dist)
			}
		}
	}

	return maxDoors
}

// findShortestDistances uses Djikstra's algorithm to find the shortest path to every room in the map
func (m *RoomMap) findShortestDistances(from point) int {
	m.resetDistance()
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	fmt.Println(MostDoors(input[0]))
	fmt.Println(time.Since(start))
}
