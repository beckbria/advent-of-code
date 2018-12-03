package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Fabric represents a rectangular shape with an Id
type Fabric struct {
	ID     int64
	Left   int64
	Top    int64
	Right  int64 // inclusive
	Bottom int64 // inclusive
}

type point struct {
	X int64
	Y int64
}

// ReadFabric converts a line from the input file into a Fabric object
func ReadFabric(input string) Fabric {
	// Input format: "#id @ X,Y: WidthxHeight"
	tokens := strings.Split(input, " ")
	id, err := strconv.ParseInt(tokens[0][1:], 10, 64)
	check(err)
	xyTokens := strings.Split(tokens[2], ",")
	x, err := strconv.ParseInt(xyTokens[0], 10, 64)
	check(err)
	y, err := strconv.ParseInt(xyTokens[1][:len(xyTokens[1])-1], 10, 64)
	check(err)
	sizeTokens := strings.Split(tokens[3], "x")
	width, err := strconv.ParseInt(sizeTokens[0], 10, 64)
	check(err)
	height, err := strconv.ParseInt(sizeTokens[1], 10, 64)
	check(err)

	// The -1 is because the rectangles are zero-indexed, i.e. the pixel at x=left is the first in the width.
	return Fabric{ID: id, Left: x, Top: y, Right: x + width - 1, Bottom: y + height - 1}
}

// OverlappingArea takes a list of rectangles and returns the total area of all spaces covered by 2 or more rectangles.
func OverlappingArea(input []Fabric) int {
	hitCount := make(map[point]int64)
	for _, f := range input {
		for x := f.Left; x <= f.Right; x++ {
			for y := f.Top; y <= f.Bottom; y++ {
				p := point{X: x, Y: y}
				hitCount[p]++
			}
		}
	}

	area := 0
	for _, h := range hitCount {
		if h > 1 {
			area++
		}
	}
	return area
}

func intersects(a Fabric, b Fabric) bool {
	return !((a.Left > b.Right) || (a.Right < b.Left) || (a.Top > b.Bottom) || (a.Bottom < b.Top))
}

// DistinctFabric takes a list of rectangles and returns the ID of the rectangle which has no intersection with any other
func DistinctFabric(input []Fabric) (int64, error) {
	for i := 0; i < len(input); i++ {
		intersect := false
		for j := 0; !intersect && j < len(input); j++ {
			if j == i {
				continue
			}
			intersect = intersects(input[i], input[j])
		}
		if !intersect {
			return input[i].ID, nil
		}
	}
	return 0, fmt.Errorf("No non-intersecting fabric")
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []Fabric
	for scanner.Scan() {
		f := ReadFabric(scanner.Text())
		input = append(input, f)
	}
	check(scanner.Err())
	fmt.Printf("Overlapping Area: %d\n", OverlappingArea(input))
	distinct, err := DistinctFabric(input)
	check(err)
	fmt.Printf("Distinct Fabric: %d\n", distinct)
}
