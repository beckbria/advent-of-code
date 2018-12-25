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

const (
	debug = false

	noConstellationID = -1
	neighborThreshold = 3
)

var (
	// Input format: "# #"
	regEx = regexp.MustCompile("^(-?\\d+),(-?\\d+),(-?\\d+),(-?\\d+)$")
)

type spacetime struct {
	x, y, z, t    int
	constellation int
}

// Manhattan Distance
func (v1 *spacetime) distance(v2 *spacetime) int {
	return abs(v1.x-v2.x) + abs(v1.y-v2.y) + abs(v1.z-v2.z) + abs(v1.t-v2.t)
}

// Helpers
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Input Parsing
func readSpacetime(input string) spacetime {
	tokens := regEx.FindStringSubmatch(input)
	x, err := strconv.Atoi(tokens[1])
	check(err)
	y, err := strconv.Atoi(tokens[2])
	check(err)
	z, err := strconv.Atoi(tokens[3])
	check(err)
	t, err := strconv.Atoi(tokens[4])
	check(err)
	return spacetime{x: x, y: y, z: z, t: t, constellation: noConstellationID}
}

func readVectors(input []string) []*spacetime {
	v := make([]*spacetime, 0)
	for _, s := range input {
		st := readSpacetime(s)
		v = append(v, &st)
	}
	return v
}

// Find connected sub-graphs.  O(n^2), but for N=1000 who cares, it runs in 8ms
func findConstellations(points []*spacetime) int {
	constellationID := 0
	for _, st := range points {
		if st.constellation != noConstellationID {
			continue
		}

		if debug {
			fmt.Printf("Starting constellation %d at [%d,%d,%d,%d]\n", constellationID, st.x, st.y, st.z, st.t)
		}
		st.constellation = constellationID
		constellationID++
		findNeighbors(st, points)
	}
	return constellationID
}

func findNeighbors(st *spacetime, points []*spacetime) {
	for _, neighbor := range points {
		if neighbor.constellation != noConstellationID {
			continue
		}

		if st.distance(neighbor) <= neighborThreshold {
			if debug {
				fmt.Printf("[%d,%d,%d,%d] is near [%d,%d,%d,%d], adding to constellation %d\n",
					neighbor.x, neighbor.y, neighbor.z, neighbor.t,
					st.x, st.y, st.z, st.t, st.constellation)
			}
			neighbor.constellation = st.constellation
			findNeighbors(neighbor, points)
		}
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
	st := readVectors(input)
	fmt.Println(findConstellations(st))
	fmt.Println(time.Since(start))
}
