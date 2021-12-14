package main

import (
	"fmt"
	"log"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

const debug = false

func main() {
	program := intcode.ReadIntCode("input.txt")

	sw := lib.NewStopwatch()
	// Part 1
	m := buildMap(program)
	//m.print()
	fmt.Println(m.distanceToOxygen(&home))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	m = buildMap(program)
	fmt.Println(m.timeToOxygenation())
	fmt.Println(sw.Elapsed())
}

const (
	// The directions as instructions understood by the drone
	north int64 = 1
	south int64 = 2
	west  int64 = 3
	east  int64 = 4

	// The statuses returned by the drone
	hitWall     int64 = 0
	moved       int64 = 1
	foundOxygen int64 = 2

	wall   = '#'
	oxygen = 'O'
	hall   = ' '

	infiniteDistance = int64(99999999)
)

var (
	invalidPoint = lib.Point{X: infiniteDistance, Y: infiniteDistance}
	home         = lib.Point{X: 0, Y: 0}
)

// dirToAocDir maps the drone direction to our direction class
func dirToAocDir(dir int64) lib.Direction {
	switch dir {
	case north:
		return lib.North
	case south:
		return lib.South
	case east:
		return lib.East
	case west:
		return lib.West
	default:
		log.Fatalf("Unexpected direction: %d\n", dir)
		return -1
	}
}

// aocDirToDir maps our direction class to the drone instruction
func aocDirToDir(dir lib.Direction) int64 {
	switch dir {
	case lib.North:
		return north
	case lib.South:
		return south
	case lib.East:
		return east
	case lib.West:
		return west
	default:
		log.Fatalf("Unexpected direction: %d\n", dir)
		return -1
	}
}

// cell represents a single room in the maze
type cell struct {
	contents  rune
	distance  int64
	preceding lib.Point
	visited   bool
}

func newCell(c rune) *cell {
	cl := cell{contents: c, distance: infiniteDistance, visited: false}
	return &cl
}

type cellMap map[lib.Point]*cell

// oxygenLocation finds a point in the maze containing oxygen
func (m cellMap) oxygenLocation() lib.Point {
	for pt, c := range m {
		if c.contents == oxygen {
			return pt
		}
	}
	return lib.Point{X: -1, Y: -1}
}

// distanceToOxygen finds the distance to the oxygen from a location
func (m cellMap) distanceToOxygen(start *lib.Point) int64 {
	pt := m.oxygenLocation()
	m.findAllShortestPaths(start, &pt)
	return m[pt].distance
}

// timeToOxygenation indicates how much time will pass before oxygen permeates the entire maze
func (m cellMap) timeToOxygenation() int64 {
	pt := m.oxygenLocation()
	m.findAllShortestPaths(&pt, nil)
	maxDistance := int64(0)
	for _, c := range m {
		if c.distance != infiniteDistance {
			maxDistance = lib.Max(maxDistance, c.distance)
		}
	}
	return maxDistance
}

// print draws the map to stdout
func (m cellMap) print() {
	minX := infiniteDistance
	minY := infiniteDistance
	maxX := -infiniteDistance
	maxY := -infiniteDistance
	for pt := range m {
		minX = lib.Min(minX, pt.X)
		maxX = lib.Max(maxX, pt.X)
		minY = lib.Min(minY, pt.Y)
		maxY = lib.Max(maxY, pt.Y)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pt := lib.Point{X: x, Y: y}
			c, found := m[pt]
			if pt == home {
				fmt.Print("D")
			} else if found {
				fmt.Print(string(c.contents))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

// shortestPath will return the shortest path between two points
func (m cellMap) shortestPath(start, end *lib.Point) []lib.Point {
	if debug {
		fmt.Print("Navigating from ")
		fmt.Print(start)
		fmt.Print(" to ")
		fmt.Println(end)
	}

	fakeEnd := false
	if !m.contains(end) {
		// We're navigating to an unexplored point; temporarily insert it into the map so we can navigate to it
		m[*end] = newCell(hall)
		fakeEnd = true
	}

	m.findAllShortestPaths(start, end)

	// Walk the path back to the start
	path := []lib.Point{}
	pt := *end
	for pt != invalidPoint {
		path = append(path, pt)
		pt = m[pt].preceding
	}
	if debug {
		fmt.Print("Path: ")
		fmt.Println(path)
	}
	if fakeEnd {
		delete(m, *end)
	}
	return path
}

// findAllShortestPaths will find the shortest path from a start location to all halls in the
// map.  If end is non-nil, it will return early once it has found the shortest path to the end
// location
func (m cellMap) findAllShortestPaths(start, end *lib.Point) {
	for _, c := range m {
		c.distance = infiniteDistance
		c.visited = false
	}

	m[*start].distance = 0
	m[*start].preceding = invalidPoint

	for true {
		// Find the closest unvisited point
		var bestPoint lib.Point
		var bestCell *cell
		for pt, c := range m {
			if !c.visited && c.contents != wall && (bestCell == nil || c.distance < bestCell.distance) {
				bestPoint = pt
				bestCell = c
			}
		}

		if bestCell == nil || (end != nil && bestPoint == *end) {
			// We've visited all points
			return
		}

		bestCell.visited = true
		// Add its neighbors
		d := bestCell.distance + 1
		for _, n := range bestPoint.Neighbors() {
			if c, found := m[n]; found && c.contents != wall {
				if d < c.distance {
					c.distance = d
					c.preceding = bestPoint
				}
			}
		}
	}
}

// contains returns true if the map contains the specified point
func (m cellMap) contains(pt *lib.Point) bool {
	_, found := m[*pt]
	return found
}

// mazeRunner represents a class that interacts with an intcode program to explore and map a maze
type mazeRunner struct {
	c         intcode.Computer
	io        *intcode.StreamInputOutput
	pos       lib.Point
	m         cellMap
	toExplore []lib.Point
}

// newMazeRunner creates a new MazeRunner object for mapping and searching a maze
func newMazeRunner(p intcode.Program) *mazeRunner {
	mr := mazeRunner{
		c:         intcode.NewComputer(p),
		io:        intcode.NewStreamInputOutput([]int64{}),
		pos:       lib.Point{X: 0, Y: 0},
		m:         make(cellMap),
		toExplore: []lib.Point{}}
	mr.io.Debug = debug
	mr.toExplore = append(mr.toExplore, mr.pos)
	mr.c.Io = mr.io
	mr.m[mr.pos] = newCell(hall)
	mr.c.RunToNextInput()

	return &mr
}

// moveTo moves the drone to the specified location in the maze
func (mr *mazeRunner) moveTo(pt *lib.Point) {
	if *pt != mr.pos {
		path := mr.m.shortestPath(&mr.pos, pt)
		for _, i := range pathToInstructions(path) {
			mr.io.AppendInput(i)
			// Take the input
			mr.c.Step()
			mr.c.RunToNextInput()
			if debug {
				// Verify our path moves us to the right place
				if mr.io.LastOutput() == hitWall {
					log.Fatalf("Hit wall when moving from [%d,%d]\n", mr.pos.X, mr.pos.Y)
				}
				switch i {
				case north:
					mr.pos.Y--
				case south:
					mr.pos.Y++
				case west:
					mr.pos.X--
				case east:
					mr.pos.X++
				}
			}
		}
		if debug && mr.pos != *pt {
			log.Fatalf("Expected to be at [%d,%d], actually at [%d,%d]\n", pt.X, pt.Y, mr.pos.X, mr.pos.Y)
		}
		mr.pos = *pt
	}
}

// probe attempts to move one space in a direction.  If it fails, it records the location of a wall.  If
// it succeeds, it notes the location for future exploration and backtracks
func (mr *mazeRunner) probe(dir int64) {
	aocD := dirToAocDir(dir)
	pt := lib.Point{X: mr.pos.X + aocD.DeltaX(), Y: mr.pos.Y + aocD.DeltaY()}
	if debug {
		fmt.Printf("At [%d,%d], probing [%d,%d]", mr.pos.X, mr.pos.Y, pt.X, pt.Y)
	}
	if !mr.m.contains(&pt) {
		mr.io.AppendInput(dir)
		mr.c.Step()
		mr.c.RunToNextInput()
		if mr.io.LastOutput() != hitWall {
			// We should look at this cell later
			mr.toExplore = append([]lib.Point{pt}, mr.toExplore...)
			if debug {
				fmt.Print(": Should Explore\n")
			}

			// Backtrack
			mr.io.AppendInput(aocDirToDir(aocD.Inverse()))
			mr.c.Step()
			mr.c.RunToNextInput()
		} else {
			if debug {
				fmt.Print(": Hit Wall\n")
			}
			mr.m[pt] = newCell(wall)
		}
	} else if debug {
		fmt.Print(": Already known\n")
	}
}

// explore goes to a location and adds its contents to the map
func (mr *mazeRunner) explore(target *lib.Point) {
	mr.moveTo(target)
	if mr.io.LastOutput() == foundOxygen {
		mr.m[*target] = newCell(oxygen)
	} else {
		mr.m[*target] = newCell(hall)
	}

	for _, dir := range []int64{north, south, east, west} {
		mr.probe(dir)
	}
}

// buildMap uses depth-first search to fill out the map
func buildMap(p intcode.Program) cellMap {
	mr := newMazeRunner(p)

	// We start in a hall, but the computer doesn't output anything to indicate that.
	// To avoid explore() crashing trying to read an output that doesn't exist, manually
	// probe the initial position
	for _, dir := range []int64{north, south, east, west} {
		mr.probe(dir)
	}

	for len(mr.toExplore) > 0 {
		// Take the first element in the list
		target := mr.toExplore[0]
		mr.toExplore = mr.toExplore[1:]
		// If we've already visited, ignore it
		if mr.m.contains(&target) {
			continue
		}

		// Go to the target and mark its type
		mr.explore(&target)
	}

	return mr.m
}

// pathToInstructions takes a path generated by shortestPath and converts it into
// the drone instructions necessary to move to the start of the path
func pathToInstructions(path []lib.Point) []int64 {
	inst := []int64{}
	// The path starts at the destination, so go in reverse order
	for i := len(path) - 2; i >= 0; i-- {
		from := path[i+1]
		to := path[i]
		if from.X == to.X-1 {
			inst = append(inst, east)
		} else if from.X == to.X+1 {
			inst = append(inst, west)
		} else if from.Y == to.Y-1 {
			inst = append(inst, south)
		} else if from.Y == to.Y+1 {
			inst = append(inst, north)
		} else {
			log.Fatalf("Cannot move from [%d,%d] to [%d,%d]", from.X, from.Y, to.X, to.Y)
		}
	}
	return inst
}
