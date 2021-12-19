package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const debug = false

// Resource is an enum of Ground/Trees/Lumberyard
type Resource rune

// Resource
const (
	Ground     = '.'
	Trees      = '|'
	Lumberyard = '#'
)

// Yard represents a map of the construction yard - map[x][y]Resource
type Yard struct {
	res    map[int]map[int]Resource
	width  int
	height int
}

type point struct {
	x, y int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadYard parses the initial contents of the yard
func ReadYard(input []string) Yard {
	yard := Yard{
		height: len(input),
		width:  len(input[0]),
		res:    make(map[int]map[int]Resource)}
	for y, line := range input {
		for x, c := range []rune(line) {
			if _, present := yard.res[x]; !present {
				yard.res[x] = make(map[int]Resource)
			}
			yard.res[x][y] = Resource(c)
		}
	}
	return yard
}

// AdvanceTime determines the state of the yard after a
// period of minutes
func AdvanceTime(yard *Yard, rounds int) {
	// Create a copy of the input yard so that we
	// can move between the two without allocating
	// a new yard for every iteration
	var yard2 Yard
	copyYard(yard, &yard2)

	for i := 0; i < rounds; i++ {
		advanceTimeWorker(yard, &yard2)
		i++
		if i == rounds {
			// Copy the data back to the in/out Yard
			copyYard(&yard2, yard)
			break
		}
		advanceTimeWorker(&yard2, yard)
	}
}

func advanceTimeWithCycle(yard *Yard, rounds int) {
	// Create a copy of the input yard so that we
	// can move between the two without allocating
	// a new yard for every iteration
	var yard2 Yard
	copyYard(yard, &yard2)
	cycleStart, cycleEnd := findCycle(&yard2)
	cycleLength := cycleEnd - cycleStart
	solutionRounds := ((rounds - cycleStart) % cycleLength) + cycleStart

	for i := 0; i < solutionRounds; i++ {
		advanceTimeWorker(yard, &yard2)
		i++
		if i == rounds {
			// Copy the data back to the in/out Yard
			copyYard(&yard2, yard)
			break
		}
		advanceTimeWorker(&yard2, yard)
	}
}

// Looks for a repeated pattern.  Returns (start, end)
func findCycle(yard *Yard) (int, int) {
	patternsSeen := make(map[string]int)
	for i := 0; ; i++ {
		ystr := YardString(yard)
		start, present := patternsSeen[ystr]
		if present {
			if debug {
				fmt.Printf("Cycle found from %d to %d\n", start, i)
			}
			return start, i
		}
		patternsSeen[ystr] = i
		AdvanceTime(yard, 1)
	}
}

func advanceTimeWorker(in, out *Yard) {
	for x := 0; x < in.width; x++ {
		for y := 0; y < in.height; y++ {
			_, trees, lumberyards := countSurrounding(in, x, y)
			switch in.res[x][y] {
			case Ground:
				if trees >= 3 {
					out.res[x][y] = Trees
				} else {
					out.res[x][y] = Ground
				}

			case Trees:
				if lumberyards >= 3 {
					out.res[x][y] = Lumberyard
				} else {
					out.res[x][y] = Trees
				}

			case Lumberyard:
				if (lumberyards >= 1) && (trees >= 1) {
					out.res[x][y] = Lumberyard
				} else {
					out.res[x][y] = Ground
				}
			}
		}
	}
}

// Counts the contents of the surrounding spaces.
// Returns open, trees, lumberyards
func countSurrounding(yard *Yard, x, y int) (int, int, int) {
	open, trees, lumberyards := 0, 0, 0
	pts := []point{
		point{x: x - 1, y: y - 1},
		point{x: x, y: y - 1},
		point{x: x + 1, y: y - 1},
		point{x: x - 1, y: y},
		point{x: x + 1, y: y},
		point{x: x - 1, y: y + 1},
		point{x: x, y: y + 1},
		point{x: x + 1, y: y + 1},
	}

	for _, pt := range pts {
		if (pt.x < 0) || (pt.y < 0) || (pt.x >= yard.width) || (pt.y >= yard.height) {
			continue
		}
		switch yard.res[pt.x][pt.y] {
		case Ground:
			open++
		case Trees:
			trees++
		case Lumberyard:
			lumberyards++
		}
	}

	return open, trees, lumberyards
}

func countAll(yard *Yard) (int, int, int) {
	open, trees, lumberyards := 0, 0, 0

	for x := 0; x < yard.width; x++ {
		for y := 0; y < yard.height; y++ {
			switch yard.res[x][y] {
			case Ground:
				open++
			case Trees:
				trees++
			case Lumberyard:
				lumberyards++
			}
		}
	}

	return open, trees, lumberyards
}

// ResourceValue returns # of lumberyards * # of trees
func ResourceValue(yard *Yard) int {
	_, trees, lumberyards := countAll(yard)
	return trees * lumberyards
}

func copyYard(from, to *Yard) {
	to.height = from.height
	to.width = from.width
	to.res = make(map[int]map[int]Resource)
	for x := 0; x < from.width; x++ {
		to.res[x] = make(map[int]Resource)
		for y := 0; y < from.height; y++ {
			to.res[x][y] = from.res[x][y]
		}
	}
}

// YardString returns a graphical representation of a yard
func YardString(yard *Yard) string {
	var sb strings.Builder
	for y := 0; y < yard.height; y++ {
		for x := 0; x < yard.width; x++ {
			sb.WriteRune(rune(yard.res[x][y]))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func main() {
	file, err := os.Open("2018/18/input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	yard := ReadYard(input)
	AdvanceTime(&yard, 10)
	fmt.Println(ResourceValue(&yard))
	fmt.Println(time.Since(start))
	start = time.Now()
	yard = ReadYard(input)
	advanceTimeWithCycle(&yard, 1000000000)
	fmt.Println(ResourceValue(&yard))
	fmt.Println(time.Since(start))
}
