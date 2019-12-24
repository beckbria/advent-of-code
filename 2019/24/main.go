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
	fmt.Println(sw.Elapsed())
}

type bugs map[aoc.Point]bool

func (b bugs) score() int64 {
	score := int64(0)
	for pt, bug := range b {
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

func (b bugs) next() bugs {
	newBugs := make(bugs)
	for pt := range b.adjacent() {
		count := 0
		neighbors := []aoc.Point{
			aoc.Point{X: pt.X - 1, Y: pt.Y},
			aoc.Point{X: pt.X + 1, Y: pt.Y},
			aoc.Point{X: pt.X, Y: pt.Y - 1},
			aoc.Point{X: pt.X, Y: pt.Y + 1}}
		for _, n := range neighbors {
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

// Returns a set of the points adjacent to bugs
func (b bugs) adjacent() bugs {
	adj := make(bugs)
	for pt := range b {
		for x := aoc.Max(0, pt.X-1); x <= aoc.Min(4, pt.X+1); x++ {
			for y := aoc.Max(0, pt.Y-1); y <= aoc.Min(4, pt.Y+1); y++ {
				adj[aoc.Point{X: x, Y: y}] = true
			}
		}
	}
	return adj
}

func (b bugs) toString() string {
	var ret strings.Builder
	for y := int64(0); y < 5; y++ {
		for x := int64(0); x < 5; x++ {
			if b[aoc.Point{X: x, Y: y}] {
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
				b[aoc.Point{X: int64(x), Y: int64(y)}] = true
			}
		}
	}
	return b
}

func firstDuplicate(b bugs) bugs {
	iteration := 0
	seen := make(map[string]int)
	for current := b; true; current = current.next() {
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
