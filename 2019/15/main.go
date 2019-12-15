package main

import (
	"fmt"
	"log"

	"../aoc"
	"../intcode"
)

const debug = false

func main() {
	program := intcode.ReadIntCode("input.txt")

	sw := aoc.NewStopwatch()
	// Part 1
	m := buildMap2(program)
	fmt.Println(m.distanceToOxygen(&home))
	fmt.Println(sw.Elapsed())
	m.print()

	// Part 2
	//sw.Reset()
	//fmt.Println(finalGameScore(program))
	//fmt.Println(sw.Elapsed())
}

const (
	// The directions as instructions understood by the drone
	north int64 = 1
	south int64 = 2
	west  int64 = 3
	east  int64 = 4

	hitWall     int64 = 0
	moved       int64 = 1
	foundOxygen int64 = 2

	wall   = '#'
	oxygen = 'O'
	hall   = ' '

	infiniteDistance = int64(99999999)
)

var (
	invalidPoint = aoc.Point{X: infiniteDistance, Y: infiniteDistance}
	home         = aoc.Point{X: 0, Y: 0}
)

// dirToAocDir maps the drone direction to our direction class
func dirToAocDir(dir int64) aoc.Direction {
	switch dir {
	case north:
		return aoc.North
	case south:
		return aoc.South
	case east:
		return aoc.East
	case west:
		return aoc.West
	default:
		log.Fatalf("Unexpected direction: %d\n", dir)
		return -1
	}
}

// aocDirToDir maps our direction class to the drone instruction
func aocDirToDir(dir aoc.Direction) int64 {
	switch dir {
	case aoc.North:
		return north
	case aoc.South:
		return south
	case aoc.East:
		return east
	case aoc.West:
		return west
	default:
		log.Fatalf("Unexpected direction: %d\n", dir)
		return -1
	}
}

type cell struct {
	contents  rune
	distance  int64
	preceding aoc.Point
	visited   bool
}

func newCell(c rune) *cell {
	cl := cell{contents: c, distance: infiniteDistance, visited: false}
	return &cl
}

type cellMap map[aoc.Point]*cell

func (m cellMap) cellsWithUnexploredNeighbors() []aoc.Point {
	u := []aoc.Point{}
	for pt, c := range m {
		if c.contents == wall {
			continue
		}
		unexplored := false
		for _, n := range neighbors(&pt) {
			if !m.contains(&n) {
				// We haven't mapped this neighbor yet
				unexplored = true
				break
			}
		}
		if unexplored {
			u = append(u, pt)
		}
	}

	return u
}

func neighbors(pt *aoc.Point) []aoc.Point {
	return []aoc.Point{
		aoc.Point{X: pt.X - 1, Y: pt.Y},
		aoc.Point{X: pt.X + 1, Y: pt.Y},
		aoc.Point{X: pt.X, Y: pt.Y - 1},
		aoc.Point{X: pt.X, Y: pt.Y + 1},
	}
}

func (m cellMap) distanceToOxygen(start *aoc.Point) int {
	for pt, c := range m {
		if c.contents == oxygen {
			return len(m.shortestPath(start, &pt))
		}
	}
	return -1
}

func (m cellMap) print() {
	minX := infiniteDistance
	minY := infiniteDistance
	maxX := -infiniteDistance
	maxY := -infiniteDistance
	for pt := range m {
		minX = aoc.Min(minX, pt.X)
		maxX = aoc.Max(maxX, pt.X)
		minY = aoc.Min(minY, pt.Y)
		maxY = aoc.Max(maxY, pt.Y)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c, found := m[aoc.Point{X: x, Y: y}]
			if found {
				fmt.Print(string(c.contents))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (m cellMap) shortestPath(start, end *aoc.Point) []aoc.Point {
	if debug {
		fmt.Print("Navigating from ")
		fmt.Print(start)
		fmt.Print(" to ")
		fmt.Println(end)
	}

	for _, c := range m {
		c.distance = infiniteDistance
		c.visited = false
	}

	m[*start].distance = 0
	m[*start].preceding = invalidPoint

	fakeEnd := false
	if !m.contains(end) {
		// We're navigating to an unexplored point; temporarily insert it into the map so we can navigate to it
		m[*end] = newCell(hall)
		fakeEnd = true
	}

	for true {
		// Find the closest unvisited point
		var bestPoint aoc.Point
		var bestCell *cell
		for pt, c := range m {
			if !c.visited && (bestCell == nil || c.distance < bestCell.distance) {
				bestPoint = pt
				bestCell = c
			}
		}

		bestCell.visited = true
		if bestPoint == *end {
			// Walk the path back to the start
			path := []aoc.Point{}
			for bestPoint != invalidPoint {
				path = append(path, bestPoint)
				bestPoint = m[bestPoint].preceding
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

		// Otherwise, add its neighbors
		d := bestCell.distance + 1
		for _, n := range neighbors(&bestPoint) {
			if c, found := m[n]; found {
				if d < c.distance {
					c.distance = d
					c.preceding = bestPoint
				}
			}
		}
	}

	return []aoc.Point{}
}

// returns true if the action caused the unit to move
func (m cellMap) populate(pos aoc.Point, response int64) bool {

	if debug {
		fmt.Printf("Tried to move to [%d,%d], got %d\n", pos.X, pos.Y, response)
	}

	switch response {
	case hitWall:
		m[pos] = newCell(wall)
		return false
	case foundOxygen:
		m[pos] = newCell(oxygen)
		return true
	case moved:
		m[pos] = newCell(hall)
		return true
	}
	log.Fatalf("Unexpected response: %d", response)
	return false
}

func (m cellMap) contains(pt *aoc.Point) bool {
	_, found := m[*pt]
	return found
}

func buildMap(p intcode.Program) cellMap {
	return buildMap2(p)
}

type mazeRunner struct {
	c         intcode.Computer
	io        *intcode.StreamInputOutput
	pos       aoc.Point
	m         cellMap
	toExplore []aoc.Point
}

func newMazeRunner(p intcode.Program) *mazeRunner {
	mr := mazeRunner{
		c:         intcode.NewComputer(p),
		io:        intcode.NewStreamInputOutput([]int64{}),
		pos:       aoc.Point{X: 0, Y: 0},
		m:         make(cellMap),
		toExplore: []aoc.Point{}}
	mr.toExplore = append(mr.toExplore, mr.pos)
	mr.c.Io = mr.io
	mr.m[mr.pos] = newCell(hall)
	mr.c.RunToNextInput()

	return &mr
}

func (mr *mazeRunner) moveTo(pt *aoc.Point) {
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

func (mr *mazeRunner) probe(dir int64) {
	aocD := dirToAocDir(dir)
	pt := aoc.Point{X: mr.pos.X + aocD.DeltaX(), Y: mr.pos.Y + aocD.DeltaY()}
	if debug {
		fmt.Printf("At [%d,%d], probing [%d,%d]", mr.pos.X, mr.pos.Y, pt.X, pt.Y)
	}
	if !mr.m.contains(&pt) {
		mr.io.AppendInput(dir)
		mr.c.Step()
		mr.c.RunToNextInput()
		if mr.io.LastOutput() != hitWall {
			// We should look at this cell later
			mr.toExplore = append(mr.toExplore, pt)
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

func (mr *mazeRunner) explore(target *aoc.Point) {
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

// Use breadth-first search to fill out the map
func buildMap2(p intcode.Program) cellMap {
	mr := newMazeRunner(p)

	// We start in a hall, but the computer doesn't output anything to indicate that
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

func buildMap1(p intcode.Program) cellMap {
	pos := aoc.Point{X: 0, Y: 0}
	m := make(cellMap)
	m[pos] = newCell(hall)
	c := intcode.NewComputer(p)
	io := intcode.NewStreamInputOutput([]int64{})
	c.Io = io

	for pts := m.cellsWithUnexploredNeighbors(); len(pts) > 0; pts = m.cellsWithUnexploredNeighbors() {
		for _, pt := range pts {
			// Move to the point
			if pt != pos {
				path := m.shortestPath(&pos, &pt)
				inst := pathToInstructions(path)
				if debug {
					fmt.Print("Instructions: ")
					fmt.Println(inst)
				}
				for _, i := range inst {
					io.AppendInput(i)
					if i == east {
						pos.X++
					} else if i == west {
						pos.X--
					} else if i == north {
						pos.Y--
					} else if i == south {
						pos.Y++
					}
					// Take the input
					c.Step()
					c.RunToNextInput()
					if debug {
						fmt.Print("Pos: ")
						fmt.Println(pos)
					}
				}
				if pos != pt {
					log.Fatalf("Expected to be at [%d,%d], actually at [%d,%d]", pt.X, pt.Y, pos.X, pos.Y)
				}
			}

			// Explore its neighbors
			io.AppendInput(west)
			c.Step()
			c.RunToNextInput()
			if m.populate(aoc.Point{X: pos.X - 1, Y: pos.Y}, io.LastOutput()) {
				// Backtrack
				io.AppendInput(east)
				c.Step()
				c.RunToNextInput()
			}

			io.AppendInput(east)
			c.Step()
			c.RunToNextInput()
			if m.populate(aoc.Point{X: pos.X + 1, Y: pos.Y}, io.LastOutput()) {
				// Backtrack
				io.AppendInput(west)
				c.Step()
				c.RunToNextInput()
			}

			io.AppendInput(north)
			c.Step()
			c.RunToNextInput()
			if m.populate(aoc.Point{X: pos.X, Y: pos.Y - 1}, io.LastOutput()) {
				// Backtrack
				io.AppendInput(south)
				c.Step()
				c.RunToNextInput()
			}

			io.AppendInput(south)
			c.Step()
			c.RunToNextInput()
			if m.populate(aoc.Point{X: pos.X, Y: pos.Y + 1}, io.LastOutput()) {
				// Backtrack
				io.AppendInput(north)
				c.Step()
				c.RunToNextInput()
			}
		}
	}

	return m
}

func pathToInstructions(path []aoc.Point) []int64 {
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
