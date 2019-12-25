package main

import (
	"container/list"
	"fmt"
	"log"

	"../aoc"
)

const debug = false

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	m := readMaze(input)
	fmt.Println(m.distance("AA", "ZZ"))
	// 658 is too low
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()

	fmt.Println(sw.Elapsed())
}

const (
	wall             = '#'
	hallway          = '.'
	invalid          = ' '
	infiniteDistance = int64(99999999)
)

type cell struct {
	adjacent map[*cell]bool
	distance int64
}

func newCell() *cell {
	c := cell{adjacent: make(map[*cell]bool)}
	return &c
}

type maze struct {
	grid         map[aoc.Point]*cell
	named        map[string][]aoc.Point
	reverseNamed map[aoc.Point]string
}

func (m *maze) distance(fromName, toName string) int64 {
	m.assertTerminus(fromName)
	m.assertTerminus(toName)
	m.resetDistance()
	from := m.named[fromName][0]
	to := m.named[toName][0]

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
	if len(m.named[name]) != 1 {
		log.Fatalf("Expected exactly one point for %s\n", name)
	}
}

func (m *maze) resetDistance() {
	for _, c := range m.grid {
		c.distance = infiniteDistance
	}
}

func newMaze() *maze {
	m := maze{grid: make(map[aoc.Point]*cell), named: make(map[string][]aoc.Point), reverseNamed: make(map[aoc.Point]string)}
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
				m.grid[aoc.Point{X: int64(x), Y: int64(y)}] = newCell()
			}
		}
	}

	// Read any named points
	for y, row := range grid {
		for x, p := range row {
			if aoc.IsUpper(byte(p)) {
				label := ""
				pt := aoc.Point{}

				// Is this a horizontal label?
				if (x < (len(row) - 1)) && aoc.IsUpper(byte(grid[y][x+1])) {
					label = string([]rune{grid[y][x], grid[y][x+1]})
					pt.Y = int64(y)
					if x > 0 && grid[y][x-1] == hallway { // Are we attached to the space to the left?
						pt.X = int64(x - 1)
					} else if x < (len(row)-2) && grid[y][x+2] == hallway { // How about the space to the right?
						pt.X = int64(x + 2)
					} else {
						log.Fatalf("Could not find hallway attached to horizontal label %s [%d,%d]\n", label, x, y)
					}
				}

				// Is this a vertical label?
				if y < (len(grid)-1) && aoc.IsUpper(byte(grid[y+1][x])) {
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
				}

				if len(label) > 0 {
					if debug {
						fmt.Printf("Found %s at [%d,%d]\n", label, pt.X, pt.Y)
					}
					if _, found := m.named[label]; !found {
						m.named[label] = []aoc.Point{}
					}
					m.named[label] = append(m.named[label], pt)
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
			for _, eq := range m.named[label] {
				if eq == pt {
					continue
				}
				//neighbors = append(neighbors, eq.Neighbors()...)
				neighbors = append(neighbors, eq)
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
