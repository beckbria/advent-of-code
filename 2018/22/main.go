package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

const (
	debug     = false
	debugPath = false
)

type point struct {
	x, y int
}

type path struct {
	cost int
	prev *rescuer
}

type cave struct {
	depth         int
	target        point
	geologicCache map[point]int
	shortestPath  map[rescuer]path
}

type region int

const (
	rocky  region = 0
	wet    region = 1
	narrow region = 2
)

func (t region) toString() string {
	switch t {
	case rocky:
		return "rocky"
	case wet:
		return "wet"
	case narrow:
		return "narrow"
	default:
		return fmt.Sprintf("Unknown region: %d", int(t))
	}
}

type tool int

const (
	gear    tool = 1
	torch   tool = 2
	neither tool = 3
)

func (t tool) toString() string {
	switch t {
	case gear:
		return "gear"
	case torch:
		return "torch"
	case neither:
		return "neither"
	default:
		return fmt.Sprintf("Unknown tool: %d", int(t))
	}
}

type rescuer struct {
	location point
	equipped tool
}

func (r *rescuer) toString(c *cave) string {
	return fmt.Sprintf("(%d,%d) [%s] holding %s",
		r.location.x, r.location.y,
		c.regionType(r.location).toString(),
		r.equipped.toString())
}

var rescuerStart = rescuer{equipped: torch, location: point{x: 0, y: 0}}

const (
	unreachable = math.MaxInt32
	switchTime  = 7
	moveTime    = 1
)

func makeCave(depth int, target point) cave {
	c := cave{depth: depth, target: target, geologicCache: make(map[point]int), shortestPath: make(map[rescuer]path)}
	c.shortestPath[rescuerStart] = path{cost: 0, prev: nil}
	if debug {
		fmt.Println(c.toString())
	}
	return c
}

func (c *cave) toString() string {
	var sb strings.Builder
	for y := 0; y <= c.target.y; y++ {
		for x := 0; x <= c.target.x; x++ {
			switch c.regionType(point{x: x, y: y}) {
			case wet:
				sb.WriteRune('=')
			case rocky:
				sb.WriteRune('.')
			case narrow:
				sb.WriteRune('|')
			}
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
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
		if curr.location == c.target && curr.equipped == torch {
			if debugPath {
				r := &curr
				for r != nil {
					p := c.shortestPath[*r]
					fmt.Println(r.toString(c))
					r = p.prev
				}
			}
			return c.shortestPath[curr].cost
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
			for _, ti := range allTools {
				t := tool(ti)
				if isToolOk(t, currRegion) && isToolOk(t, neighborRegion) {
					newRescuer := rescuer{location: neighbor, equipped: t}
					newCost := c.shortestPath[curr].cost + moveTime
					if t != curr.equipped {
						newCost += switchTime
					}
					if oldPath, present := c.shortestPath[newRescuer]; !present || (newCost < oldPath.cost) {
						var prev *rescuer
						if debugPath {
							prev = &curr
						} else {
							prev = nil
						}
						c.shortestPath[newRescuer] = path{cost: newCost, prev: prev}
						toProcess = append(toProcess, newRescuer)
						// TODO: A heap should be more effiient, but by my calculations we should only have to deal
						// with ~50k nodes, so sorting every step is painful but not the end of the world
						sort.Slice(toProcess, func(i, j int) bool {
							return c.shortestPath[toProcess[i]].cost < c.shortestPath[toProcess[j]].cost
						})
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
	fmt.Println(c.timeToTarget())
	fmt.Println(time.Since(start))
}
