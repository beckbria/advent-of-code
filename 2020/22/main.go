package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/22
// Card Games

// Turn this on to have the state of each round printed
const debug = true

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	decks := parseDecks(lines)
	fmt.Println(step1(decks))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	decks = parseDecks(lines)
	fmt.Println(step2(decks))
	fmt.Println(sw.Elapsed())
}

type deck struct {
	cards []int
}

func newDeck() *deck {
	d := deck{}
	d.cards = make([]int, 0)
	return &d
}

func (d *deck) size() int {
	return len(d.cards)
}

func (d *deck) isEmpty() bool {
	return len(d.cards) == 0
}

func (d *deck) draw() int {
	c := d.cards[0]
	if len(d.cards) > 1 {
		d.cards = d.cards[1:]
	} else {
		d.cards = []int{}
	}
	return c
}

func (d *deck) add(c int) {
	d.cards = append(d.cards, c)
}

func (d *deck) score() int64 {
	score := int64(0)
	lc := len(d.cards)
	for i := range d.cards {
		multiplier := int64(i + 1)
		card := int64(d.cards[lc-i-1])
		score += card * multiplier
	}
	return score
}

func (d *deck) toString() string {
	var b strings.Builder
	b.Grow(len(d.cards) * 3)
	for i, n := range d.cards {
		fmt.Fprintf(&b, "%d", n)
		if i < (len(d.cards) - 1) {
			b.WriteString(", ")
		}
	}
	return b.String()
}

func parseDecks(lines []string) map[int]*deck {
	decks := make(map[int]*deck)
	readPlayer := true
	player := -1
	for _, l := range lines {
		if readPlayer {
			if !strings.HasPrefix(l, "Player") {
				log.Fatalf("Invalid player line: %s", l)
			}
			p, _ := strconv.Atoi(l[7:8])
			player = p
			decks[player] = newDeck()
			readPlayer = false
		} else if len(l) == 0 {
			// Next line is a new player
			readPlayer = true
		} else {
			n, _ := strconv.Atoi(l)
			decks[player].add(n)
		}
	}
	return decks
}

func step1(decks map[int]*deck) int64 {
	for !decks[1].isEmpty() && !decks[2].isEmpty() {
		one := decks[1].draw()
		two := decks[2].draw()
		if one > two {
			decks[1].add(one)
			decks[1].add(two)
		} else {
			decks[2].add(two)
			decks[2].add(one)
		}
	}
	if decks[1].isEmpty() {
		return decks[2].score()
	}
	return decks[1].score()
}

func step2(decks map[int]*deck) int64 {
	winner := recursiveCombat(decks)
	return decks[winner].score()
}

// An iterator to track how many games (including recursive games) have been played
var nextGame = 1

// recursiveCombat simulates a game and returns the winner
func recursiveCombat(decks map[int]*deck) int {
	game := nextGame
	nextGame++

	if debug {
		fmt.Printf("\n=== Game %d ===\n", game)
	}

	seenGames := make(map[string]bool)
	round := 1
	winner := -1
	for !decks[1].isEmpty() && !decks[2].isEmpty() {
		if debug {
			fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
			fmt.Printf("Player 1's deck: %s\n", decks[1].toString())
			fmt.Printf("Player 2's deck: %s\n", decks[2].toString())
		}
		h := stateHash(decks)
		if _, found := seenGames[h]; found {
			// Infinite loop prevention rule
			if debug {
				fmt.Printf("Game state seen before.  Player 1 wins\n")
			}
			return 1
		}
		seenGames[h] = true

		one := decks[1].draw()
		two := decks[2].draw()
		if debug {
			fmt.Printf("Player 1 plays: %d\n", one)
			fmt.Printf("Player 2 plays: %d\n", two)
		}
		if decks[1].size() >= one && decks[2].size() >= two {
			// Recurse
			nd := make(map[int]*deck)
			nd[1] = newDeck()
			newone := decks[1].cards[:one]
			nd[1].cards = make([]int, one)
			copy(nd[1].cards, newone)
			nd[2] = newDeck()
			newtwo := decks[2].cards[:two]
			nd[2].cards = make([]int, two)
			copy(nd[2].cards, newtwo)

			if debug {
				fmt.Println("Playing a sub-game to determine the winner...")
			}
			winner = recursiveCombat(nd)
			if debug {
				fmt.Printf("\n...anyway, back to game %d.\n", game)
			}
		} else if one > two {
			winner = 1
		} else {
			winner = 2
		}

		switch winner {
		case 1:
			decks[1].add(one)
			decks[1].add(two)
		case 2:
			decks[2].add(two)
			decks[2].add(one)
		default:
			log.Fatalf("Unexpected winner: %d", winner)
		}
		if debug {
			fmt.Printf("Player %d wins round %d of game %d!\n", winner, round, game)
		}
		round++
	}

	if debug {
		if game == 1 {
			fmt.Printf("\n== Post-game results ==\n")
			fmt.Printf("Player 1's deck: %s\n", decks[1].toString())
			fmt.Printf("Player 2's deck: %s\n", decks[2].toString())
		}
		fmt.Printf("The winner of game %d is player %d!\n", game, winner)
	}
	return winner
}

func stateHash(decks map[int]*deck) string {
	return decks[1].toString() + "/" + decks[2].toString()
}
