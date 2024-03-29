package main

import (
	"container/list"
	"fmt"
	"log"

	"github.com/beckbria/advent-of-code/2019/lib"
)

const debug = false

func main() {
	input := lib.ReadFileLines("2019/20/input.txt")
	sw := lib.NewStopwatch()
	// Part 1
	m := readMaze(input)
	fmt.Println(m.distance("AA", "ZZ", false))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(m.distance("AA", "ZZ", true))
	fmt.Println(sw.Elapsed())
}

const (
	wall             = '#'
	hallway          = '.'
	invalid          = ' '
	infiniteDistance = int64(99999999)
)

var (
	unknownPoint  = lib.Point{X: infiniteDistance, Y: infiniteDistance}
	unknownPoint3 = lib.Point3{X: infiniteDistance, Y: infiniteDistance, Z: infiniteDistance}
)

type cell struct {
	location lib.Point3
	adjacent map[*cell]bool
	distance int64
}

func newCell(pt lib.Point3) *cell {
	c := cell{adjacent: make(map[*cell]bool), location: pt, distance: infiniteDistance}
	return &c
}

type warp struct {
	internal, external lib.Point
}

func newWarp() *warp {
	w := warp{internal: unknownPoint, external: unknownPoint}
	return &w
}

type maze struct {
	grid         map[lib.Point]*cell
	named        map[string]*warp
	reverseNamed map[lib.Point]string
	searchGrid   map[lib.Point3]*cell
}

func (m *maze) getCell(x, y, z int64) *cell {
	pt := lib.Point3{X: x, Y: y, Z: z}
	if _, found := m.searchGrid[pt]; found {
		pt0 := lib.Point{X: x, Y: y}
		if _, found := m.grid[pt0]; !found {
			log.Fatalf("Cannot create [%d,%d,%d]\n", x, y, z)
		}

		m.searchGrid[pt] = newCell(pt)
	}
	return m.searchGrid[pt]
}

func (m *maze) distance(fromName, toName string, threeD bool) int64 {
	m.assertTerminus(fromName)
	m.assertTerminus(toName)
	m.resetDistance()
	from := m.named[fromName].external
	to := m.named[toName].external

	// Breadth-first search
	m.grid[from].distance = 0

	toProcess := list.New()
	toProcess.PushBack(m.grid[from])
	for toProcess.Len() > 0 {
		current := toProcess.Front()
		nCost := current.Value.(*cell).distance + 1
		for n := range current.Value.(*cell).adjacent {
			if n.distance > nCost {
				n.distance = nCost
				toProcess.PushBack(n)
				if n.location.X == to.X && n.location.Y == to.Y {
					return m.grid[to].distance
				}
			}
		}
		toProcess.Remove(current)
	}

	return m.grid[to].distance
}

func (m *maze) assertTerminus(name string) {
	if _, found := m.named[name]; !found {
		log.Fatalf("Unknown point %s\n", name)
	}
	if m.named[name].internal != unknownPoint {
		log.Fatalf("Expected external-only point for %s\n", name)
	}
}

func (m *maze) resetDistance() {
	for _, c := range m.grid {
		c.distance = infiniteDistance
	}
}

func newMaze() *maze {
	m := maze{grid: make(map[lib.Point]*cell), named: make(map[string]*warp), reverseNamed: make(map[lib.Point]string)}
	return &m
}

func readMaze(input []string) *maze {
	m := newMaze()

	// Read into a 2D array for quick lookup
	grid := make([][]rune, len(input))
	for i, row := range input {
		grid[i] = []rune(row)
	}

	// Create cells for every hallway point
	for y, row := range grid {
		for x, p := range row {
			if p == hallway {
				pt := lib.Point{X: int64(x), Y: int64(y)}
				pt3 := lib.Point3{X: int64(x), Y: int64(y), Z: 0}
				m.grid[pt] = newCell(pt3)
			}
		}
	}

	// Read any named points
	for y, row := range grid {
		for x, p := range row {
			if lib.IsUpper(byte(p)) {
				label := ""
				pt := lib.Point{}
				internal := true
				internalThreshold := 5

				// Is this a horizontal label?
				if (x < (len(row) - 1)) && lib.IsUpper(byte(grid[y][x+1])) {
					label = string([]rune{grid[y][x], grid[y][x+1]})
					pt.Y = int64(y)
					if x > 0 && grid[y][x-1] == hallway { // Are we attached to the space to the left?
						pt.X = int64(x - 1)
					} else if x < (len(row)-2) && grid[y][x+2] == hallway { // How about the space to the right?
						pt.X = int64(x + 2)
					} else {
						log.Fatalf("Could not find hallway attached to horizontal label %s [%d,%d]\n", label, x, y)
					}
					internal = x > internalThreshold && x < (len(row)-internalThreshold)
				}

				// Is this a vertical label?
				if y < (len(grid)-1) && lib.IsUpper(byte(grid[y+1][x])) {
					if len(label) > 0 {
						log.Fatalf("Found Horizontal+Vertical label at [%d,%d]\n", x, y)
					}
					label = string([]rune{grid[y][x], grid[y+1][x]})
					pt.X = int64(x)
					if y > 0 && grid[y-1][x] == hallway { // Are we attached to the space above?
						pt.Y = int64(y - 1)
					} else if y < (len(grid)-2) && grid[y+2][x] == hallway { // What about the space below?
						pt.Y = int64(y + 2)
					} else {
						log.Fatalf("Could not find hallway attached to vertical label %s [%d,%d]\n", label, x, y)
					}
					internal = y > internalThreshold && y < (len(grid)-internalThreshold)
				}

				if len(label) > 0 {
					if debug {
						fmt.Printf("Found %s at [%d,%d]\n", label, pt.X, pt.Y)
					}
					if _, found := m.named[label]; !found {
						m.named[label] = newWarp()
					}
					if internal {
						m.named[label].internal = pt
					} else {
						m.named[label].external = pt
					}
					m.reverseNamed[pt] = label
				}
			}
		}
	}

	// Build the adjacency lists
	for pt, c := range m.grid {
		// Start with traditional neighbors
		neighbors := pt.Neighbors()

		// Check any other named points
		if label, found := m.reverseNamed[pt]; found {
			w := m.named[label]
			if w.internal != pt {
				neighbors = append(neighbors, w.internal)
			}
			if w.external != pt {
				neighbors = append(neighbors, w.external)
			}
		}

		// Check traditional neighbors
		for _, n := range neighbors {
			if nc, found := m.grid[n]; found {
				c.adjacent[nc] = true
			}
		}
	}

	return m
}
