package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/13
// Bus schedules

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	start, departures := parseSchedule(lines)
	fmt.Println(step1(start, departures))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2Naive(departures, int64(100000000000000)))
	fmt.Println(sw.Elapsed())
}

const skip = -1

func parseSchedule(lines []string) (int64, []int64) {
	start, _ := strconv.ParseInt(lines[0], 10, 64)
	departures := make([]int64, 0)
	for _, d := range strings.Split(lines[1], ",") {
		if d == "x" {
			departures = append(departures, skip)
		} else {
			bus, _ := strconv.ParseInt(d, 10, 64)
			departures = append(departures, bus)
		}
	}
	return start, departures
}

// Part 1: When will the next bus depart?
func step1(start int64, departures []int64) int64 {
	nearestTime := start + 10000
	nearestBus := int64(0)
	for _, bus := range departures {
		if bus == skip {
			continue
		}
		nextStop := ((start + bus) / bus) * bus
		if nextStop < nearestTime {
			nearestTime = nextStop
			nearestBus = bus
		}
	}

	return nearestBus * (nearestTime - start)
}

func step2Naive(departures []int64, searchStart int64) int64 {
	offsets, maxMultiple, maxOffset := calcOffset(departures)
	start := ((searchStart + maxMultiple) / maxMultiple) * maxMultiple

SEARCH:
	for i := start; ; i += maxMultiple {
		for bus, off := range offsets {
			if (i+off)%bus != 0 {
				continue SEARCH
			}
		}
		return i - maxOffset
	}
}

func calcOffset(departures []int64) (map[int64]int64, int64, int64) {
	offset := make(map[int64]int64)
	maxMultiple := int64(-1)
	maxMultipleIdx := int64(-1)
	for i, bus := range departures {
		if bus == skip {
			continue
		}
		if bus > maxMultiple {
			maxMultiple = bus
			maxMultipleIdx = int64(i)
		}

		// Print out the modulo result and the divisor to plug into a tool such as https://www.dcode.fr/chinese-remainder
		fmt.Printf("%d\t%d\n", (bus*5-int64(i))%bus, bus)
	}
	for i, bus := range departures {
		if bus == skip {
			continue
		}
		offset[bus] = int64(i) - maxMultipleIdx
	}

	return offset, maxMultiple, maxMultipleIdx
}
