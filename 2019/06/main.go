package main

import (
	"fmt"
	"log"
	"strings"

	"../../aoc"
)

const debug = false
const invalidOrbitCount = -1

// https://adventofcode.com/2019/day/6
// Given a set of links, build a graph.  Do some graph stuff on it

type body struct {
	name           string
	orbits         *body
	orbitedBy      map[*body]bool
	orbitedByCount int
	orbitsCount    int
	visited        bool
}

func newBody(name string) *body {
	b := body{
		orbits:         nil,
		orbitedBy:      make(map[*body]bool),
		name:           name,
		orbitedByCount: invalidOrbitCount,
		orbitsCount:    invalidOrbitCount,
		visited:        false}
	return &b
}

type bodySet map[string]*body

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	bodies := parseBodies(lines)
	// Part 1
	fmt.Println(totalOrbits(bodies))
	// Part 2
	you := bodies["YOU"]
	san := bodies["SAN"]
	fmt.Println(distance(you, san) - 2) // Subtract 2 to account for the you and san nodes
	fmt.Println(sw.Elapsed())
}

const notFound = -1

func distance(start, end *body) int {
	if start == nil {
		return notFound
	}

	if start == end {
		// found
		return 0
	}

	// Breadth-first search
	start.visited = true
	neighbors := []*body{start.orbits}
	for moon := range start.orbitedBy {
		neighbors = append(neighbors, moon)
	}
	for _, n := range neighbors {
		if n == nil || n.visited {
			continue
		}
		dist := distance(n, end)
		if dist != notFound {
			return 1 + dist
		}
	}

	return notFound
}

func parseBodies(lines []string) bodySet {
	bodies := make(bodySet)
	for _, l := range lines {
		b := strings.Split(l, ")")
		planetName := b[0]
		moonName := b[1]
		if _, found := bodies[planetName]; !found {
			bodies[planetName] = newBody(planetName)
		}
		if _, found := bodies[moonName]; !found {
			bodies[moonName] = newBody(moonName)
		}
		planet := bodies[planetName]
		moon := bodies[moonName]
		planet.orbitedBy[moon] = true
		if moon.orbits != nil {
			log.Fatalf("%s already orbits something", moonName)
		}
		moon.orbits = planet
	}
	return bodies
}

func totalOrbits(bodies bodySet) int {
	orbits := 0
	for _, b := range bodies {
		orbits += orbitsCount(b)
	}
	return orbits
}

func orbitsCount(b *body) int {
	if b == nil {
		return -1
	}

	if b.orbitsCount == invalidOrbitCount {
		b.orbitsCount = 1 + orbitsCount(b.orbits)
	}
	return b.orbitsCount
}
