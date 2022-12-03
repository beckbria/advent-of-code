package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2022/lib"
)

// https://adventofcode.com/2022/day/2

const (
	Rock     int64 = 1
	Paper          = 2
	Scissors       = 3
)

const debug = false

func parseThrow(r rune) (int64, error) {
	switch r {
	case 'A':
		return Rock, nil
	case 'B':
		return Paper, nil
	case 'C':
		return Scissors, nil
	}
	return Rock, fmt.Errorf("Invalid throw: '%s'", string(r))
}

type Round struct {
	 opponent, player int64
}

func (r *Round) score() int64 {
	fromThrow := r.player
	fromOutcome := int64(0)
	if r.opponent == r.player {
		fromOutcome = 3
	}
	if r.player == (r.opponent + 1) || (r.player == Rock && r.opponent == Scissors) {
		fromOutcome = 6
	}
	total := fromThrow + fromOutcome
	if (debug) {
		fmt.Printf("Score Opp(%d) Player(%d) - Throw=%d, Total=%d\n", r.opponent, r.player, fromThrow, total)
	}
	return total
}

func makeRoundStep1(inp string) (*Round, error) {
	var err error
	r := Round{}
	if len(inp) != 3 {
		return nil, fmt.Errorf("Invalid input length: '%s'", inp)
	}
	input := []rune(inp)
	r.opponent, err = parseThrow(input[0])
	if err != nil {
		return nil, err
	}
	r.player, err = parseThrow(input[2] - 'X' + 'A')
	if (debug) {
		fmt.Printf("Read '%s' to Round {%d, %d}\n", inp, r.opponent, r.player)
	}
	return &r, nil
}

func makeRoundStep2(inp string) (*Round, error) {
	var err error
	r := Round{}
	if len(inp) != 3 {
		return nil, fmt.Errorf("Invalid input length: '%s'", inp)
	}
	input := []rune(inp)
	r.opponent, err = parseThrow(input[0])
	if err != nil {
		return nil, err
	}

	// Read the final glyph as if it were a throw
	r.player, err = parseThrow(input[2] - 'X' + 'A')
	if err != nil {
		return nil, err
	}

	/*
	  Use that throw as an offset.
	  Throw		Val	Outcome	Offset
	  Rock		1	Lose	-1
	  Paper		2	Tie		0
	  Scissors 	3	Win		1

	  Then:
	  1) Subtract 1 from the opponent to become 0-indexed
	  2) Add 3 because negative modulos are annoying
	  3) Subtract 2 for the difference between the enum and the desired offset
	  4) Notice that all of the above add up to 0 and do nothing
	  5) Add the player's throw
	  6) Modulo 3
	  7) Add the 1 back to become 1-indexed again
	*/
	r.player = (r.opponent + r.player) % 3 + 1

	if (debug) {
		fmt.Printf("Read '%s' to Round {%d, %d}\n", inp, r.opponent, r.player)
	}
	return &r, nil
}

func main() {
	lines := lib.ReadFileLines("2022/02/input.txt")
	
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

func parseLinesStep1(lines []string) []*Round {
	rounds := []*Round{}
	for _, line := range lines {
		r, err := makeRoundStep1(line)
		lib.Check(err)
		rounds = append(rounds, r)
	}
	return rounds
}

func parseLinesStep2(lines []string) []*Round {
	rounds := []*Round{}
	for _, line := range lines {
		r, err := makeRoundStep2(line)
		lib.Check(err)
		rounds = append(rounds, r)
	}
	return rounds
}

func step1(lines []string) int64 {
	rounds := parseLinesStep1(lines)
	total := int64(0)
	for _, r := range rounds {
		total += r.score()
	}
	return total
}

func step2(lines []string) int64 {
	rounds := parseLinesStep2(lines)
	total := int64(0)
	for _, r := range rounds {
		total += r.score()
	}
	return total
}
