package main

import (
	"fmt"
	"math"
	"sort"

	"../../aoc"
)

// https://adventofcode.com/2020/day/5
// Parse seats on an airplane

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	seats := parseSeats(lines)
	fmt.Println(highestID(seats))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(missingID(seats))
	fmt.Println(sw.Elapsed())
}

type seat struct {
	row, col int
}

func (s *seat) ID() int {
	return 8*s.row + s.col
}

func parseSeats(lines []string) []seat {
	var seats []seat
	for _, l := range lines {
		var s seat
		min, max := 0, 127
		for i := 0; i < 7; i++ {
			min, max = split(min, max, l[i] == 'F')
		}
		s.row = min
		min, max = 0, 7
		for i := 7; i < 10; i++ {
			min, max = split(min, max, l[i] == 'L')
		}
		s.col = min
		seats = append(seats, s)
	}
	return seats
}

func split(min, max int, lower bool) (int, int) {
	halfway := (float64(min) + float64(max)) / 2.0
	if lower {
		max = int(math.Floor(halfway))
	} else {
		min = int(math.Ceil(halfway))
	}
	return min, max
}

func highestID(seats []seat) int {
	max := -1
	for _, s := range seats {
		id := s.ID()
		if id > max {
			max = id
		}
	}
	return max
}

func missingID(seats []seat) int {
	var ids []int
	for _, s := range seats {
		ids = append(ids, s.ID())
	}
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			return ids[i-1] + 1
		}
	}
	return -1
}
