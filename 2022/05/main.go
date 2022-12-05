package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/5

const (
	DEBUG_OFF = 0
	DEBUG_ERROR = 1
	DEBUG_FINAL_OUTPUT = 3
	DEBUG_PARSE = 5
	DEBUG_MOVE = 10

	debugLevel = DEBUG_OFF
)

var (
	moveRegex = regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`)
)

type moveInstruction struct {
	count, from, to int64
}

type columns [][]string

func (c columns) moveOne(from, to int64) {
	target := c[from - 1][0]
	if debugEnabled(DEBUG_MOVE) {
		fmt.Printf("Moving %s from %d to %d\n", target, from, to)
	}
	c[from - 1] = c[from - 1][1:]
	c[to - 1] = append([]string{target}, c[to - 1]...)
	if debugEnabled(DEBUG_MOVE) {
		fmt.Printf("%v\n", c)
	}
}

func (c columns) moveManyOneByOne(from, to, count int64) {
	if debugEnabled(DEBUG_ERROR) && int64(len(c[from - 1])) < count {
		fmt.Printf("Cannot move %d from column %d.\n%v\n", count, from, c)
		lib.Check(fmt.Errorf("Failed to move\n"))
	}

	if debugEnabled(DEBUG_MOVE) {
		fmt.Printf("Moving %d from %d to %d\n", count, from, to)
	}

	if count < 0 {
		lib.Check(fmt.Errorf("Invalid move count"))
	}

	for i := int64(0); i < count; i++ {
		c.moveOne(from, to)
	}
}

func (c columns) moveAllAtOnce(m *moveInstruction) {
	if debugEnabled(DEBUG_ERROR) && int64(len(c[m.from - 1])) < m.count {
		fmt.Printf("Cannot move %d from column %d.\n%v\n", m.count, m.from, c)
		lib.Check(fmt.Errorf("Failed to move\n"))
	}

	target := make([]string, m.count)
	copy(target, c[m.from - 1][:m.count])
	if debugEnabled(DEBUG_MOVE) {
		fmt.Printf("Moving %d (%v) from %d to %d\n", m.count, target, m.from, m.to)
	}
	c[m.from - 1] = c[m.from - 1][m.count:]
	c[m.to - 1] = append(target, c[m.to - 1]...)
	if debugEnabled(DEBUG_MOVE) {
		fmt.Printf("%v\n", c)
	}
}

func (c columns) moveOneByOne (m *moveInstruction) {
	c.moveManyOneByOne(m.from, m.to, m.count)
}

func debugEnabled(logLevel int) bool {
	return debugLevel >= logLevel
}

func parse(lines []string) (columns, []*moveInstruction) {
	i := 0
	cols := columns{}
	// Parse the initial column state
	for ; i < len(lines) && len(lines[i]) > 0; i++ {
		// Ensure the columns are sufficiently padded
		for ;len(cols) < (len(lines[i]) + 1) / 4; {
			cols = append(cols, []string{})
		}

		for j := 0; j < len(cols); j++ {

			pos := 1 + (j * 4)
			if pos > len(lines[i]) {
				if debugEnabled(DEBUG_ERROR) {
					fmt.Printf("Cannot read position %d from line '%s'\n", pos, lines[i])
					continue
				}
			}
			
			c := lines[i][pos:pos+1]
			r := []rune(c)[0]
			if r >= 'A' && r <= 'Z' {
				cols[j] = append(cols[j], c)
			}

			if debugEnabled(DEBUG_PARSE) {
				fmt.Printf("Read '%s' at position %d(%d) from '%s'\n", c, j, pos, lines[i])
			}
		}
	}

	i++	// Move past the blank line

	moves := []*moveInstruction{}
	for ; i < len(lines); i++ {
		// Parse the moves
		tokens := moveRegex.FindStringSubmatch(lines[i])

		m := moveInstruction{}
		m.count, _ = strconv.ParseInt(tokens[1], 10, 64)
		m.from, _ = strconv.ParseInt(tokens[2], 10, 64)
		m.to, _ = strconv.ParseInt(tokens[3], 10, 64)
		moves = append(moves, &m)

		if debugEnabled(DEBUG_PARSE) {
			fmt.Printf("Read Move: %d from %d to %d\n", m.count, m.from, m.to)
		}
	}

	return cols, moves
}

func main() {
	lines := lib.ReadFileLines("2022/05/input.txt")
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

func step1(lines []string) string {
	cols, moves := parse(lines)

	if debugEnabled(DEBUG_FINAL_OUTPUT) {
		fmt.Printf("Initial State:\n%v\n", cols)
	}

	for _, m := range moves {
		cols.moveOneByOne(m)
	}

	if debugEnabled(DEBUG_FINAL_OUTPUT) {
		fmt.Printf("%v\n", cols)
	}

	firstLetters := ""
	for _, c := range cols {
		if len(c) > 0 {
			firstLetters = firstLetters + c[0]
		} else {
			firstLetters = firstLetters + " "
		}
	}

	return firstLetters
}

func step2(lines []string) string {
	cols, moves := parse(lines)

	if debugEnabled(DEBUG_FINAL_OUTPUT) {
		fmt.Printf("Initial State:\n%v\n", cols)
	}

	for _, m := range moves {
		cols.moveAllAtOnce(m)
	}

	if debugEnabled(DEBUG_FINAL_OUTPUT) {
		fmt.Printf("%v\n", cols)
	}

	firstLetters := ""
	for _, c := range cols {
		if len(c) > 0 {
			firstLetters = firstLetters + c[0]
		} else {
			firstLetters = firstLetters + " "
		}
	}

	return firstLetters
}
