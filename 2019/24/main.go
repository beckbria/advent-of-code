package main

import (
	"fmt"
	"strings"

	"../aoc"
)

const (
	debug         = false
	debugNeighbor = debug && false
	debugScore    = debug && false
)

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	b := readBugs(input)
	dup := firstDuplicate(b)
	fmt.Println(dup.score())
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(countAfter(b, 200))
	fmt.Println(sw.Elapsed())
}

// bugs represents a set of points occupied bugs in a system reminiscent of Conway's Game of Life
type bugs map[aoc.Point3]bool

// score calculates the score of a system of bugs.  In level 0 of the grid, each bug scores points
// depending on 2 to the power of its location in the 5x5 grid
func (b bugs) score() int64 {
	score := int64(0)
	for pt, bug := range b {
		if pt.Z != 0 {
			continue
		}
		if bug {
			pow := 5*pt.Y + pt.X
			if debugScore {
				fmt.Printf("[%d,%d] == 2**%d == %d\n", pt.X, pt.Y, pow, aoc.Pow(2, pow))
			}
			score += aoc.Pow(2, pow)
		}
	}
	return score
}

// next applies the life rules to determine which locations contain bugs after the next day
func (b bugs) next(multiLevel bool) bugs {
	newBugs := make(bugs)
	for pt := range b.adjacent(multiLevel) {
		count := 0
		for n := range pointAdjacent(&pt, multiLevel) {
			if b[n] {
				count++
			}
		}
		hadBug := b[pt]
		newBugs[pt] = (hadBug && count == 1) || (!hadBug && (count == 1 || count == 2))
		if debugNeighbor {
			fmt.Printf("   [%d,%d] (%s) has %d neighbors == %s\n", pt.X, pt.Y, bugStr(hadBug), count, bugStr(newBugs[pt]))
		}
	}
	return newBugs
}

func bugStr(isBug bool) string {
	if !isBug {
		return "NOT BUG"
	}
	return "BUG"
}

// pointAdjacent returns a set of all points adjacent to a point in the bug grid
func pointAdjacent(pt *aoc.Point3, multiLevel bool) bugs {
	adj := make(bugs)

	// Add the left neighbors
	switch pt.X {
	case 0:
		if multiLevel {
			adj[aoc.Point3{X: 1, Y: 2, Z: pt.Z - 1}] = true
		}
	case 3:
		if multiLevel && pt.Y == 2 {
			for y := int64(0); y < 5; y++ {
				adj[aoc.Point3{X: 4, Y: y, Z: pt.Z + 1}] = true
			}
			break
		}
		fallthrough
	default:
		adj[aoc.Point3{X: pt.X - 1, Y: pt.Y, Z: pt.Z}] = true
	}

	// Add the right neighbors
	switch pt.X {
	case 4:
		if multiLevel {
			adj[aoc.Point3{X: 3, Y: 2, Z: pt.Z - 1}] = true
		}
	case 1:
		if multiLevel && pt.Y == 2 {
			for y := int64(0); y < 5; y++ {
				adj[aoc.Point3{X: 0, Y: y, Z: pt.Z + 1}] = true
			}
			break
		}
		fallthrough
	default:
		adj[aoc.Point3{X: pt.X + 1, Y: pt.Y, Z: pt.Z}] = true
	}

	// Add the top neighbors
	switch pt.Y {
	case 0:
		if multiLevel {
			adj[aoc.Point3{X: 2, Y: 1, Z: pt.Z - 1}] = true
		}
	case 3:
		if multiLevel && pt.X == 2 {
			for x := int64(0); x < 5; x++ {
				adj[aoc.Point3{X: x, Y: 4, Z: pt.Z + 1}] = true
			}
			break
		}
		fallthrough
	default:
		adj[aoc.Point3{X: pt.X, Y: pt.Y - 1, Z: pt.Z}] = true
	}

	// Add the bottom neighbors
	switch pt.Y {
	case 4:
		if multiLevel {
			adj[aoc.Point3{X: 2, Y: 3, Z: pt.Z - 1}] = true
		}
	case 1:
		if multiLevel && pt.X == 2 {
			for x := int64(0); x < 5; x++ {
				adj[aoc.Point3{X: x, Y: 0, Z: pt.Z + 1}] = true
			}
			break
		}
		fallthrough
	default:
		adj[aoc.Point3{X: pt.X, Y: pt.Y + 1, Z: pt.Z}] = true
	}

	return adj
}

// adjacent returns a set of the points adjacent to bugs
func (b bugs) adjacent(multiLevel bool) bugs {
	adj := make(bugs)
	for pt, isBug := range b {
		if isBug {
			for p := range pointAdjacent(&pt, multiLevel) {
				adj[p] = true
			}
		}
	}
	return adj
}

func (b bugs) toString() string {
	var ret strings.Builder
	for y := int64(0); y < 5; y++ {
		for x := int64(0); x < 5; x++ {
			if b[aoc.Point3{X: x, Y: y, Z: 0}] {
				ret.WriteRune('#')
			} else {
				ret.WriteRune('.')
			}
		}
		ret.WriteRune('\n')
	}
	return ret.String()
}

func readBugs(input []string) bugs {
	b := make(bugs)
	for y, row := range input {
		for x, c := range []rune(row) {
			if c == '#' {
				b[aoc.Point3{X: int64(x), Y: int64(y), Z: 0}] = true
			}
		}
	}
	return b
}

// firstDuplicate finds the first grid of bugs which is identical to a previously-known state.
// It uses the single-level rules.
func firstDuplicate(b bugs) bugs {
	iteration := 0
	seen := make(map[string]int)
	for current := b; true; current = current.next(false) {
		if debug {
			fmt.Printf("Iteration %d:\n", iteration)
			fmt.Println(current.toString())
		}
		s := fmt.Sprint(current)
		if prev, found := seen[s]; found {
			if debug {
				fmt.Printf("Duplicate found at iteration %d (previously iteration %d)\n", iteration, prev)
			}
			return current
		}
		seen[s] = iteration
		iteration++
	}
	return make(bugs)
}

// countAfter returns the number of bugs present across all levels after a number of iterations
func countAfter(b bugs, iterations int) int64 {
	current := b
	for i := 0; i < iterations; i++ {
		current = current.next(true)
	}
	count := int64(0)
	for _, isBug := range current {
		if isBug {
			count++
		}
	}
	return count
}
