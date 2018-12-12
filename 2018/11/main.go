package main

import (
	"fmt"
	"log"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func makeEmptyGrid() [][]int {
	grid := make([][]int, 301)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 301)
	}
	return grid
}

// MakeGrid constructs a grid according to the rules
func MakeGrid(seed int) [][]int {
	grid := makeEmptyGrid()
	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[0]); y++ {
			rid := x + 10
			power := ((rid * y) + seed) * rid
			power = ((power / 100) % 10) - 5
			grid[x][y] = power
		}
	}
	return grid
}

// LargestPowerLocation finds the largest section of the specified width and height in the grid
func LargestPowerLocation(seed, width, height int) (int, int) {
	grid := MakeGrid(seed)
	sums := makeEmptyGrid()
	for x := 1; x < (len(grid) - width); x++ {
		for y := 1; y < (len(grid[0]) - height); y++ {
			// This can be optimised, but for 3x3 grids it's not that wasteful
			s := 0
			for i := x; i < x+width; i++ {
				for j := y; j < y+height; j++ {
					s += grid[i][j]
				}
			}
			sums[x][y] = s
		}
	}

	maxX, maxY, maxSum := 0, 0, sums[0][0]
	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[0]); y++ {
			if sums[x][y] > maxSum {
				maxX, maxY, maxSum = x, y, sums[x][y]
			}
		}
	}
	return maxX, maxY
}

// LargestPowerSquare finds the largest square section (any size) of the grid
func LargestPowerSquare(seed int) (int, int, int) {
	// To any readers: This brute force solution is algorithmically awful, but for
	// an input of size 300x300 it simply doesn't matter.  My apologies.
	grid := MakeGrid(seed)
	sums := makeEmptyGrid()
	sizes := makeEmptyGrid()
	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[0]); y++ {
			s := grid[x][y]
			sums[x][y] = s
			sizes[x][y] = 1
			maxSquareSize := min(len(grid)-x, len(grid[0])-y)
			for i := 1; i < maxSquareSize; i++ {
				// Add the new right and bottom edge to get the new sum
				for j := 0; j < i; j++ {
					s += grid[x+i][y+j]
					s += grid[x+j][y+i]
				}
				s += grid[x+i][x+i] // Do this outside the loop to avoid double-counting
				if s > sums[x][y] {
					sums[x][y], sizes[x][y] = s, i+1
				}
			}
		}
	}

	maxX, maxY, maxSum, maxSize := 0, 0, sums[0][0], sizes[0][0]
	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[0]); y++ {
			if sums[x][y] > maxSum {
				maxX, maxY, maxSum, maxSize = x, y, sums[x][y], sizes[x][y]
			}
		}
	}
	return maxX, maxY, maxSize
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	start := time.Now()
	fmt.Println(LargestPowerLocation(1718, 3, 3))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(LargestPowerSquare(1718))
	fmt.Println(time.Since(start))
}
