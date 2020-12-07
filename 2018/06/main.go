package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var (
	// Input format: "X, Y"
	pointRegEx = regexp.MustCompile(`^(\d+), (\d+)$`)
)

// Point describes a point on a coordinate plane
type Point struct {
	X int64
	Y int64
}

// GridFill represents what point initiated the fill of an area and the distance to the home point
type GridFill struct {
	home       Point
	distance   int  // This could technically be computed on the fly, but memory is cheap
	noMansLand bool // If this point is of equal distance to 2 or more points, it's in no man's land
	seen       bool
}

// ReadPoint converts an input line to a Point
func ReadPoint(input string) Point {
	tokens := pointRegEx.FindStringSubmatch(input)
	x, err := strconv.ParseInt(tokens[1], 10, 64)
	check(err)
	y, err := strconv.ParseInt(tokens[2], 10, 64)
	check(err)

	// The -1 is because the rectangles are zero-indexed, i.e. the pixel at x=left is the first in the width.
	return Point{X: x, Y: y}
}

// ReadPoints reads one Point per line
func ReadPoints(input []string) []Point {
	var pts []Point
	for _, i := range input {
		pts = append(pts, ReadPoint(i))
	}
	return pts
}

// MaximumArea solves part 1 of https://adventofcode.com/2018/day/6 - Given a set of points, find the one which has
// the largest non-infinite number of points which are closer to it than to any other point
func MaximumArea(input []Point) int64 {
	grid, infiniteFill := fill(input)
	counts := countPoints(grid)
	maxArea := int64(0)
	for home, count := range counts {
		if _, found := infiniteFill[home]; !found {
			maxArea = max(maxArea, count)
		}
	}
	return maxArea
}

// Returns a grid based on the home points a map of home points which have infinite fill
// (0,0) in the grid is at (minX, minY) for the provided points.  Each grid point will
// be filled in with its nearest home point.
func fill(homes []Point) ([][]GridFill, map[Point]bool) {
	grid, fillQueue := populateGrid(homes)
	infiniteFill := make(map[Point]bool)

	for len(fillQueue) > 0 {
		// Pop the top point from the queue
		pt := fillQueue[0]
		ptGridCell := grid[pt.X][pt.Y]
		fillQueue = fillQueue[1:]

		// Add each of the four points around this one to the queue
		distance := ptGridCell.distance + 1
		adjacent := []Point{{X: pt.X - 1, Y: pt.Y}, {X: pt.X + 1, Y: pt.Y}, {X: pt.X, Y: pt.Y - 1}, {X: pt.X, Y: pt.Y + 1}}
		for _, a := range adjacent {
			if isOutsideGrid(a, grid) {
				infiniteFill[ptGridCell.home] = true
			} else {
				if grid[a.X][a.Y].seen {
					if (grid[a.X][a.Y].distance == distance) && (grid[a.X][a.Y].home != ptGridCell.home) {
						grid[a.X][a.Y].noMansLand = true
					}
				} else {
					grid[a.X][a.Y].seen = true
					grid[a.X][a.Y].home = ptGridCell.home
					grid[a.X][a.Y].distance = distance
					fillQueue = append(fillQueue, a)
				}
			}
		}
	}

	return grid, infiniteFill
}

func isOutsideGrid(pt Point, grid [][]GridFill) bool {
	return (pt.X < 0) || (pt.Y < 0) || (pt.X >= int64(len(grid))) || (pt.Y >= int64(len(grid[0])))
}

// Populates an initial grid with the specified homes, with the grid clipped
// tightly to the initial home points.  Returns the grid and a list of the
// transformed locations of the initial points.
func populateGrid(homes []Point) ([][]GridFill, []Point) {
	minX, minY, maxX, maxY, offsetHomes := outerBounds(homes)
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]GridFill, width)
	for i := int64(0); i < width; i++ {
		grid[i] = make([]GridFill, height)
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			grid[x][y].seen = false
			grid[x][y].noMansLand = false
		}
	}

	// Populate the Initial Homes
	for _, h := range offsetHomes {
		grid[h.X][h.Y].home = h
		grid[h.X][h.Y].seen = true
		grid[h.X][h.Y].distance = 0
	}

	return grid, offsetHomes
}

// Counts the number of points with a certain home that aren't in no man's land
func countPoints(grid [][]GridFill) map[Point]int64 {
	counts := make(map[Point]int64)
	for _, col := range grid {
		for _, gf := range col {
			if gf.seen && !gf.noMansLand {
				counts[gf.home]++
			}
		}
	}
	return counts
}

// Returns the edges of a rectangle containing all these points.  Returns minX, minY, maxX, maxY
func outerBounds(input []Point) (int64, int64, int64, int64, []Point) {
	minX := input[0].X
	maxX := input[0].X
	minY := input[0].X
	maxY := input[0].X
	for _, pt := range input {
		minX = min(minX, pt.X)
		maxX = max(maxX, pt.X)
		minY = min(minY, pt.Y)
		maxY = max(maxY, pt.Y)
	}
	var offsetInput []Point
	for _, i := range input {
		offsetInput = append(offsetInput, Point{X: i.X - minX, Y: i.Y - minY})
	}
	return minX, minY, maxX, maxY, offsetInput
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaximumSafeArea solves part 2 of https://adventofcode.com/2018/day/6
func MaximumSafeArea(input []Point, threshold int64) int64 {
	// For each space you go outside the grid, you add len(input) points to the
	// total distance, so we should only need threshold / len(input) points of
	// padding on each side.  This reduces the padding to a reasonable enough
	// level that the O(n^2) runtime of checking the full solution space still
	// isn't bad for our data set
	padding := int64(threshold / int64(len(input)))
	_, _, maxX, maxY, offsetHomes := outerBounds(input)

	size := int64(0)
	for x := -padding; x <= maxX+padding; x++ {
		for y := -padding; y <= maxY+padding; y++ {
			distance := int64(0)
			for _, h := range offsetHomes {
				distance += abs(h.X - int64(x))
				distance += abs(h.Y - int64(y))
			}
			if distance < threshold {
				size++
			}
		}
	}

	return size
}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
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
	pts := ReadPoints(input)
	start := time.Now()
	fmt.Println(MaximumArea(pts))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(MaximumSafeArea(pts, 10000))
	fmt.Println(time.Since(start))
}
