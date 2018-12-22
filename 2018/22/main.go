package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

const (
	debug = true
)

type point struct {
	x, y int
}

type cave struct {
	depth         int
	target        point
	geologicCache map[point]int
	shortestPath  map[rescuer]int
}

type region int

const (
	rocky  region = 0
	wet    region = 1
	narrow region = 2
)

type tool int

const (
	gear    tool = iota + 1
	torch   tool = iota + 1
	neither tool = iota + 1
)

type rescuer struct {
	location point
	equipped tool
}

var rescuerStart = rescuer{equipped: torch, location: point{x: 0, y: 0}}

const (
	unreachable = math.MaxInt32
	switchTime  = 7
	moveTime    = 1
)

func makeCave(depth int, target point) cave {
	c := cave{depth: depth, target: target, geologicCache: make(map[point]int), shortestPath: make(map[rescuer]int)}
	c.shortestPath[rescuerStart] = 0
	return c
}

func (c *cave) geologicIndex(p point) int {
	if gi, present := c.geologicCache[p]; present {
		return gi
	}

	gi := 0
	if p == c.target {
		gi = 0
	} else if p.y == 0 {
		gi = p.x * 16807
	} else if p.x == 0 {
		gi = p.y * 48271
	} else {
		gi = c.erosionLevel(point{x: p.x - 1, y: p.y}) * c.erosionLevel(point{x: p.x, y: p.y - 1})
	}
	c.geologicCache[p] = gi
	return gi
}

func (c *cave) erosionLevel(p point) int {
	return (c.geologicIndex(p) + c.depth) % 20183
}

func (c *cave) regionType(p point) region {
	return region(c.erosionLevel(p) % 3)
}

func (c *cave) riskToTarget() int {
	sum := 0
	for x := 0; x <= c.target.x; x++ {
		for y := 0; y <= c.target.y; y++ {
			sum += int(c.regionType(point{x: x, y: y}))
		}
	}
	return sum
}

func isToolOk(t tool, r region) bool {
	switch r {
	case rocky:
		return t != neither
	case wet:
		return t != torch
	case narrow:
		return t != gear
	}
	log.Fatalf("Unknown tool: %d\n", t)
	return false
}

// Oh hey look, another graph search.  Oh hey look, it's time to implement
// Djikstra's algorithm again
func (c *cave) timeToTarget() int {
	toProcess := []rescuer{rescuerStart}
	neighborOffset := []point{point{x: 0, y: -1}, point{x: -1, y: 0}, point{x: 1, y: 0}, point{x: 0, y: 1}}
	allTools := []tool{torch, gear, neither}
	for len(toProcess) > 0 {
		// Pop the to entry
		curr := toProcess[0]
		if curr.location == c.target {
			return c.shortestPath[curr]
		}
		currRegion := c.regionType(curr.location)
		toProcess = toProcess[1:]

		// Try all combinations of neighbor
		for _, no := range neighborOffset {
			neighbor := point{x: curr.location.x + no.x, y: curr.location.y + no.y}
			if (neighbor.x < 0) || (neighbor.y < 0) {
				continue
			}
			neighborRegion := c.regionType(neighbor)
			for ti := range allTools {
				t := tool(ti)
				if isToolOk(t, currRegion) && isToolOk(t, neighborRegion) {
					newRescuer := rescuer{location: neighbor, equipped: t}
					newCost := c.shortestPath[curr] + moveTime
					if t != curr.equipped {
						newCost += switchTime
					}
					if oldCost, present := c.shortestPath[newRescuer]; !present || (newCost < oldCost) {
						c.shortestPath[newRescuer] = newCost
						// TODO: May have to make this a heap to always get the shortest point
						toProcess = append(toProcess, newRescuer)
					}
				}
			}
		}
	}
	// No path found
	return -1
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	start := time.Now()
	depth := 10914
	target := point{x: 9, y: 739}
	c := makeCave(depth, target)
	fmt.Println(c.riskToTarget())
	fmt.Println(time.Since(start))
}
