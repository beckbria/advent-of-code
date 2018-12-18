package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	debug               = false
	debugInput          = false
	debugPrintReservoir = true
)

// Point in Cartesian space
type Point struct {
	x, y int
}

// Material is an enum of Clay/Water
type Material rune

// You alredy know what these are
const (
	Clay         = '#'
	Water        = '~'
	RunningWater = '|'
)

// MinMax represents a paired minimum/maximum value
type MinMax struct {
	min, max int
}

// Reservoir represents the layout of the reservoir
type Reservoir struct {
	mat  map[Point]Material
	x, y MinMax
}

var (
	// Input format: "# #"
	inputRegEx = regexp.MustCompile("^([xy])=(\\d+), ([xy])=(\\d+)..(\\d+)$")
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeReservoir() Reservoir {
	var r Reservoir
	r.mat = make(map[Point]Material)
	// Set the min/max to the max/min so that the first bit of clay added becomes the min/max
	r.x.min = math.MaxInt32
	r.y.min = math.MaxInt32
	r.x.max = math.MinInt32
	r.y.max = math.MinInt32
	return r
}

func addClay(r *Reservoir, pt Point) {
	r.mat[pt] = Clay
	r.x.min = min(r.x.min, pt.x)
	r.x.max = max(r.x.max, pt.x)
	r.y.min = min(r.y.min, pt.y)
	r.y.max = max(r.y.max, pt.y)
}

// ReadReservoir parses the input file
func ReadReservoir(input []string) Reservoir {
	r := makeReservoir()
	for _, s := range input {
		tokens := inputRegEx.FindStringSubmatch(s)
		constant, err := strconv.Atoi(tokens[2])
		check(err)
		rangeStart, err := strconv.Atoi(tokens[4])
		check(err)
		rangeEnd, err := strconv.Atoi(tokens[5])
		check(err)
		if tokens[1] == "x" {
			x := constant
			for y := rangeStart; y <= rangeEnd; y++ {
				if debugInput {
					fmt.Printf("Clay: %d,%d\n", x, y)
				}
				addClay(&r, Point{x: x, y: y})
			}
		} else { // tokens[1] == "y"
			y := constant
			for x := rangeStart; x <= rangeEnd; x++ {
				if debugInput {
					fmt.Printf("Clay: %d,%d\n", x, y)
				}
				addClay(&r, Point{x: x, y: y})
			}
		}
	}
	return r
}

// Fill starts filling the reservoir from the faucet
func Fill(r *Reservoir, faucet Point) {
	if debug {
		fmt.Printf("Starting fill x %d->%d y %d->%d\n", r.x.min, r.x.max, r.y.min, r.y.max)
	}
	y := max(faucet.y, r.y.min)
	start := Point{x: faucet.x, y: y}
	fillWorker(r, start)
}

// Depth-first search.  Returns true if this space should expand further
func fillWorker(r *Reservoir, current Point) bool {
	if debug {
		fmt.Printf("Filling %d,%d\n", current.x, current.y)
	}
	r.mat[current] = Water
	sideFlow := false
	if current.y < r.y.max {
		_, present := r.mat[Point{x: current.x, y: current.y + 1}]
		if present {
			if debug {
				fmt.Printf("    Substance underneath %d,%d\n", current.x, current.y)
			}
			sideFlow = true
		} else if !present {
			if debug {
				fmt.Printf("    Nothing underneath %d,%d\n", current.x, current.y)
			}
			sideFlow = fillWorker(r, Point{x: current.x, y: current.y + 1})
		}
	}

	backflow := false
	if sideFlow {
		if debug {
			fmt.Printf("    Starting side flow for %d,%d\n", current.x, current.y)
		}
		sides := []Point{Point{x: current.x - 1, y: current.y}, Point{x: current.x + 1, y: current.y}}
		backflow = true
		for _, pt := range sides {
			// If there's any value (clay or water), we don't want to search a square that's already been seen
			_, present := r.mat[pt]
			if !present {
				result := fillWorker(r, pt)
				backflow = backflow && result
			}
		}
		if debug {
			fmt.Printf("    %t -- backflow for %d,%d\n", backflow, current.x, current.y)
		}
	}

	if !backflow {
		r.mat[current] = RunningWater
		for x := current.x - 1; r.mat[Point{x: x, y: current.y}] == Water; x-- {
			r.mat[Point{x: x, y: current.y}] = RunningWater
		}
		for x := current.x + 1; r.mat[Point{x: x, y: current.y}] == Water; x++ {
			r.mat[Point{x: x, y: current.y}] = RunningWater
		}
	}

	return backflow
}

// FloodCount solves part 1 of the problem - fills the reservoir and counts the water
func FloodCount(input []string) (int, int) {
	r := ReadReservoir(input)
	Fill(&r, Point{x: 500, y: 0})
	if debugPrintReservoir {
		PrintReservoir(&r)
	}
	standing := 0
	running := 0
	for _, v := range r.mat {
		if v == Water {
			standing++
		} else if v == RunningWater {
			running++
		}
	}
	return running, standing
}

// PrintReservoir prints a graphical representation
func PrintReservoir(r *Reservoir) {
	for y := r.y.min; y <= r.y.max; y++ {
		for x := r.x.min; x <= r.x.max; x++ {
			v, present := r.mat[Point{x: x, y: y}]
			if present {
				fmt.Printf("%c", rune(v))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
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
	start := time.Now()
	running, standing := FloodCount(input)
	fmt.Println(running + standing)
	fmt.Println(standing)
	fmt.Println(time.Since(start))
}
