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

var (
	// Input format: "# #"
	nanobotRegEx = regexp.MustCompile("^pos=<(-?\\d+),(-?\\d+),(-?\\d+)>, r=(\\d+)$")
)

type nanobot struct {
	// Position and activation radius
	x, y, z, r int
}

// Manhattan distance to another nanobot
func (b1 *nanobot) distance(b2 *nanobot) int {
	return abs(b1.x-b2.x) + abs(b1.y-b2.y) + abs(b1.z-b2.z)
}

type fleet []nanobot

func (f *fleet) inRangeOfStrongest() int {
	// Find the strongest nanobot
	strongest := (*f)[0]
	for _, n := range *f {
		if n.r > strongest.r {
			strongest = n
		}
	}
	// Count the bots in range of the strongest
	count := 0
	for _, n := range *f {
		if strongest.distance(&n) <= strongest.r {
			count++
		}
	}
	return count
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

// ReadLine parses a line from the input file
func readNanobot(input string) nanobot {
	tokens := nanobotRegEx.FindStringSubmatch(input)
	x, err := strconv.Atoi(tokens[1])
	check(err)
	y, err := strconv.Atoi(tokens[2])
	check(err)
	z, err := strconv.Atoi(tokens[3])
	check(err)
	r, err := strconv.Atoi(tokens[4])
	check(err)
	return nanobot{x: x, y: y, z: z, r: r}
}

func readFleet(input []string) fleet {
	f := make(fleet, 0)
	for _, s := range input {
		f = append(f, readNanobot(s))
	}
	return f
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
	f := readFleet(input)
	// Part 1
	fmt.Println(f.inRangeOfStrongest())
	fmt.Println(time.Since(start))
}
