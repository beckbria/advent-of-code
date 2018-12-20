package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const (
	debug             = true
	debugPattern      = debug && false
	debugMap          = debug && true
	debugShortestPath = debug && false
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

func (m *RoomMap) set(pt Point, c Component) {
	if _, present := (*m)[pt.x]; !present {
		(*m)[pt.x] = make(map[int]Location)
	}
	(*m)[pt.x][pt.y] = makeLocation(c)
}

func (m *RoomMap) setIfAbsent(pt Point, c Component) {
	if _, present := (*m)[pt.x]; !present {
		(*m)[pt.x] = make(map[int]Location)
	}
	if _, present := (*m)[pt.x][pt.y]; !present {
		(*m)[pt.x][pt.y] = makeLocation(c)
	}
}

// Direction is an enum representing all of the options in the instruction pattern.
type Direction rune

// The directions
const (
	West         Direction = 'W'
	East                   = 'E'
	North                  = 'N'
	South                  = 'S'
	BeginGroup             = '('
	EndGroup               = ')'
	GroupOption            = '|'
	BeginPattern           = '^'
	EndPattern             = '$'
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

// Point in Cartesian space
type Point struct {
	x, y int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadPattern state machine conditions
type state int

const (
	idle       state = iota + 1
	readString       = iota + 1
	choice           = iota + 1
)

// ReadPattern turns the string RegEx into a pattern state machine
func ReadPattern(s string) Pattern {
	return readPattern(s, 0)
}

// debugOffset is added to any incides when returning errors
func readPattern(s string, debugOffset int) Pattern {
	if len(s) < 1 {
		return SinglePattern("")
	}

	// What patterns we have seen
	patterns := make([]Pattern, 0)

	status := idle     // What are we parsing right now?
	sectionStart := 0  // Where we started parsing the current pattern
	choiceDivider := 0 // Where the | divider is
	parenDepth := 0    // How many levels of lParen have we seen since we started looking at a group
	for i, c := range []rune(s) {
		switch Direction(c) {
		case BeginPattern:
			status = idle
		case West, East, North, South:
			if status == idle {
				status = readString
				sectionStart = i
			}

		case BeginGroup:
			if status == readString {
				// We've finished reading a string group.
				patterns = append(patterns, SinglePattern(s[sectionStart:i]))
				status = idle
			}

			if status == idle {
				status = choice
				choiceDivider = -1
				sectionStart = i + 1
				parenDepth = 1
			} else if status == choice {
				parenDepth++
			}

		case EndGroup:
			if status != choice {
				log.Fatalf("Found EndGroup when not parsing group at index %d\n", debugOffset+i)
			}
			parenDepth--
			if parenDepth == 0 {
				// We've reached the end of this group
				if choiceDivider < 0 {
					log.Fatalf("Found group without divider from %d-%d\n", debugOffset+sectionStart, debugOffset+i)
				}
				first := s[sectionStart:choiceDivider]
				second := s[choiceDivider+1 : i]
				if debugPattern {
					fmt.Printf("Choice between \"%s\" and \"%s\" from %d-%d\n",
						first, second, debugOffset+sectionStart, debugOffset+i)
				}
				patterns = append(patterns, ChoicePattern(
					readPattern(first, debugOffset+sectionStart),
					readPattern(second, debugOffset+choiceDivider+1)))
				status = idle
			}

		case GroupOption:
			if status != choice {
				log.Fatalf("Found EndGroup when not parsing group at index %d\n", debugOffset+i)
			}
			if parenDepth == 1 {
				if choiceDivider >= 0 {
					log.Fatalf("Found duplicate | at %d and %d\n", debugOffset+choiceDivider, debugOffset+i)
				}
				choiceDivider = i
			}
		case EndPattern:
			if status == readString {
				// We've finished reading a string group.
				patterns = append(patterns, SinglePattern(s[sectionStart:i]))
				status = idle
			}

			if status != idle {
				log.Fatalf("Unexpected end of pattern parsing \"%s\"\n", s)
			}
		}
	}

	// Terminate a final pattern
	if status == readString {
		// We've finished reading a string group.
		patterns = append(patterns, SinglePattern(s[sectionStart:]))
		status = idle
	}

	if status != idle {
		log.Fatalf("Unexpected end of pattern parsing \"%s\"\n", s)
	}

	if len(patterns) == 1 {
		return patterns[0]
	} else if len(patterns) > 1 {
		return ConcatenatePattern(patterns...)
	} else {
		log.Fatalf("Found no patterns: \"%s\"\n", s)
		return SinglePattern("")
	}
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

	origin := Point{x: 0, y: 0}
	m.addRoom(origin)
	m.addPatternToMap(p, origin)

	return m
}

func (m *RoomMap) addPatternToMap(p *Pattern, start Point) {
	if !p.isCompound() {
		m.addSimplePatternToMap(p, start)
	} else {
		for _, path := range p.paths {
			// Add each choice to the map separately
			m.addCompoundPatternToMap(path, start)
		}
	}
}

// Add a single directional path to the map.  Return its end point for any patterns
// which come after it.
func (m *RoomMap) addSimplePatternToMap(p *Pattern, start Point) Point {
	current := start
	for _, c := range []rune(p.pattern) {
		switch Direction(c) {
		case West:
			(*m)[current.x-1][current.y] = makeLocation(HDoor)
			m.addRoom(Point{x: current.x - 2, y: current.y})
			current.x -= 2
		case East:
			(*m)[current.x+1][current.y] = makeLocation(HDoor)
			m.addRoom(Point{x: current.x + 2, y: current.y})
			current.x += 2
		case North:
			(*m)[current.x][current.y-1] = makeLocation(VDoor)
			m.addRoom(Point{x: current.x, y: current.y - 2})
			current.y -= 2
		case South:
			(*m)[current.x][current.y+1] = makeLocation(VDoor)
			m.addRoom(Point{x: current.x, y: current.y + 2})
			current.y += 2
		default:
			log.Fatalf("Unexpected character in simple pattern: %c\n", c)
		}
	}
	return current
}

// Add what may be a concatenated series of patterns to the map
func (m *RoomMap) addCompoundPatternToMap(path []Pattern, start Point) {
	if len(path) < 1 {
		return
	} else if len(path) == 1 {
		m.addPatternToMap(&(path[0]), start)
	} else {
		current := path[0]
		rest := path[1:]
		if current.isCompound() {
			for _, path := range current.paths {
				// Add each choice to the map separately
				m.addCompoundPatternToMap(append(path, rest...), start)
			}
		} else {
			newStart := m.addSimplePatternToMap(&current, start)
			m.addCompoundPatternToMap(rest, newStart)
		}
	}
}

func (m *RoomMap) addRoom(pt Point) {
	m.set(Point{x: pt.x, y: pt.y}, Room)
	// Add walls
	for x := pt.x - 1; x <= pt.x+1; x++ {
		for y := pt.y - 1; y <= pt.y+1; y++ {
			m.setIfAbsent(Point{x: x, y: y}, Wall)
		}
	}
}

// MostDoors returns the most doors you must pass through to reach any room of the house along
// the shortest path
func MostDoors(s string) int {
	p := ReadPattern(s)
	m := BuildMap(&p)
	m.findShortestDistances(Point{x: 0, y: 0})

	maxDoors := 0
	for _, yr := range m {
		for _, c := range yr {
			if (c.c == Room) && (c.dist != Unreachable) {
				maxDoors = max(maxDoors, c.dist)
			}
		}
	}

	return maxDoors
}

func location(c Component, distance int) Location {
	return Location{c: c, dist: distance}
}

func isDoor(c Component) bool {
	return (c == HDoor) || (c == VDoor)
}

// findShortestDistances uses Djikstra's algorithm to find the shortest path to every room in the map
func (m *RoomMap) findShortestDistances(from Point) {
	m.resetDistance()
	(*m)[0][0] = location((*m)[0][0].c, 0)
	if debugShortestPath {
		fmt.Println("Setting origin distance to 0")
	}

	// Offset to N/W/E/S neighbors
	neighborOffset := []Point{Point{x: 0, y: -1}, Point{x: -1, y: 0}, Point{x: 1, y: 0}, Point{x: 0, y: 1}}

	// The queue of unvisited nodes
	toProcess := []Point{Point{x: 0, y: 0}}
	for len(toProcess) > 0 {
		// Pop the first element
		pt := toProcess[0]
		toProcess = toProcess[1:]
		newDistance := (*m)[pt.x][pt.y].dist + 1

		// Check for neighbors
		for _, no := range neighborOffset {
			x := pt.x + no.x
			y := pt.y + no.y
			if debugShortestPath {
				fmt.Printf("Checking for door at %d,%d\n", x, y)
			}
			if isDoor((*m)[x][y].c) {
				// Double the offset to get through the door to the room
				x += no.x
				y += no.y
				if debugShortestPath {
					fmt.Printf("Found door.  Checking room at %d,%d [Current dist %d, new %d]\n", x, y, (*m)[x][y].dist, newDistance)
				}
				if newDistance < (*m)[x][y].dist {
					if debugShortestPath {
						fmt.Printf("Updating distance to %d\n", newDistance)
					}
					(*m)[x][y] = location((*m)[x][y].c, newDistance)
					toProcess = append(toProcess, Point{x: x, y: y})
				}
			}
		}
	}
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
