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
const debug = false

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	d := parseDecks(lines)
	fmt.Println(step1(d))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	d = parseDecks(lines)
	fmt.Println(step2(d))
	fmt.Println(sw.Elapsed())
}

func step1(d decks) int64 {
	winner := -1
	for !d.finished() {
		drawn := d.draw()
		winner = drawn.winner()
		d.add(drawn, winner)
	}
	return d[winner].score()
}

func step2(d decks) int64 {
	winner := recursiveCombat(d)
	return d[winner].score()
}

// deck represents a deck of cards
type deck struct {
	cards []int
}

// newDeck creates an empty deck
func newDeck() *deck {
	d := deck{}
	d.cards = make([]int, 0)
	return &d
}

// size returns the number of cards remaining in the deck
func (d *deck) size() int {
	return len(d.cards)
}

// isEmpty indicates if the decks has no remaining cards
func (d *deck) isEmpty() bool {
	return len(d.cards) == 0
}

// draw removes and returns the top card from the deck
func (d *deck) draw() int {
	c := d.cards[0]
	if len(d.cards) > 1 {
		d.cards = d.cards[1:]
	} else {
		d.cards = []int{}
	}
	return c
}

// adds a card to the bottom of the deck.  Returns self for chaining
func (d *deck) add(c int) *deck {
	d.cards = append(d.cards, c)
	return d
}

// score calculates the game value of a deck
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

// String generates a visual representation of a deck's contents
func (d *deck) String() string {
	var b strings.Builder
	d.stringBuild(&b)
	return b.String()
}

func (d *deck) stringBuild(b *strings.Builder) {
	b.Grow(len(d.cards) * 4)
	for i, n := range d.cards {
		fmt.Fprintf(b, "%d", n)
		if i < (len(d.cards) - 1) {
			b.WriteString(", ")
		}
	}
}

// drawResult reprsents the cards drawn by each player in a round.  Includes an unused value at
// index 0 so that the draws can be accessed by player ID
type drawResult [3]int

// winner returns the winning player in a draw
func (r drawResult) winner() int {
	if r[1] > r[2] {
		return 1
	}
	return 2
}

// draw returns the cards drawn by each player in a round.
func (d decks) draw() drawResult {
	return [3]int{-1, d[1].draw(), d[2].draw()}
}

// decks represents the game state with both Player 1 and 2's decks
type decks map[int]*deck

// recurse generates a new deck for a recursive game
func (d decks) recurse(drawn drawResult) decks {
	nd := make(map[int]*deck)
	nd[1] = newDeck()
	newone := d[1].cards[:drawn[1]]
	nd[1].cards = make([]int, drawn[1])
	copy(nd[1].cards, newone)
	nd[2] = newDeck()
	newtwo := d[2].cards[:drawn[2]]
	nd[2].cards = make([]int, drawn[2])
	copy(nd[2].cards, newtwo)
	return nd
}

// finished returns true if at least one player's deck is empty
func (d decks) finished() bool {
	return d[1].isEmpty() || d[2].isEmpty()
}

func (d decks) String() string {
	var b strings.Builder
	b.Grow((len(d[1].cards) + len(d[2].cards)) * 4)
	d[1].stringBuild(&b)
	b.WriteRune('/')
	d[2].stringBuild(&b)
	return b.String()
}

// add adds the two cards drawn in a round to the winner's deck
func (d decks) add(drawn drawResult, winner int) {
	d[winner].add(drawn[winner])
	d[winner].add(drawn[opposite(winner)])
}

// parseDecks reads the decks for players 1 and 2 from input
func parseDecks(lines []string) decks {
	d := make(decks)
	readPlayer := true
	player := -1
	for _, l := range lines {
		if readPlayer {
			if !strings.HasPrefix(l, "Player") {
				log.Fatalf("Invalid player line: %s", l)
			}
			p, _ := strconv.Atoi(l[7:8])
			player = p
			d[player] = newDeck()
			readPlayer = false
		} else if len(l) == 0 {
			// Next line is a new player
			readPlayer = true
		} else {
			n, _ := strconv.Atoi(l)
			d[player].add(n)
		}
	}
	return d
}

// opposite returns the ID of the other player
func opposite(player int) int {
	return (player % 2) + 1
}

// An iterator to track how many games (including recursive games) have been played
var nextGame = 1

// recursiveCombat simulates a game and returns the winner
func recursiveCombat(d decks) int {
	game := nextGame
	nextGame++
	if debug {
		fmt.Printf("\n=== Game %d ===\n", game)
	}

	seenGames := make(map[string]bool)
	round := 1
	winner := -1
	for !d.finished() {
		if debug {
			fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
			fmt.Printf("Player 1's deck: %s\n", d[1])
			fmt.Printf("Player 2's deck: %s\n", d[2])
		}
		h := d.String()
		if _, found := seenGames[h]; found {
			// Infinite loop prevention rule
			if debug {
				fmt.Printf("Game state seen before.  Player 1 wins\n")
			}
			return 1
		}
		seenGames[h] = true

		winner = -1
		drawn := d.draw()
		if debug {
			fmt.Printf("Player 1 plays: %d\n", drawn[1])
			fmt.Printf("Player 2 plays: %d\n", drawn[2])
		}
		if d[1].size() >= drawn[1] && d[2].size() >= drawn[2] {
			// Recurse.  Create a copy of the cards for the new game.
			nd := d.recurse(drawn)

			if debug {
				fmt.Println("Playing a sub-game to determine the winner...")
			}
			winner = recursiveCombat(nd)
			if debug {
				fmt.Printf("\n...anyway, back to game %d.\n", game)
			}
		} else {
			winner = drawn.winner()
		}
		d.add(drawn, winner)

		if debug {
			fmt.Printf("Player %d wins round %d of game %d!\n", winner, round, game)
		}
		round++
	}

	if debug {
		if game == 1 {
			fmt.Printf("\n== Post-game results ==\n")
			fmt.Printf("Player 1's deck: %s\n", d[1])
			fmt.Printf("Player 2's deck: %s\n", d[2])
		}
		fmt.Printf("The winner of game %d is player %d!\n", game, winner)
	}
	return winner
}
