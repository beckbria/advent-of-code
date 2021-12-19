package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	debug          = false
	debugPrintTree = false
)

var (
	// Input format: "10 players; last marble is worth 1618 points"
	fabricRegEx = regexp.MustCompile(`^(\d+) players; last marble is worth (\d+) points$`)
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadInput parses the input format
func ReadInput(input string) (int, int64) {
	tokens := fabricRegEx.FindStringSubmatch(input)
	players, err := strconv.Atoi(tokens[1])
	check(err)
	points, err := strconv.ParseInt(tokens[2], 10, 64)
	check(err)
	return players, points
}

type gameState struct {
	scores        []int64       // The players' scores
	marbles       *list.List    // The marbles which make up the circle.  Ascending index is clockwise
	currentMarble *list.Element // The index of the "current" marble
	nextPlayer    int           // The index of the next player to place a marble
	nextMarble    int64         // The value of the next marble to be placed
}

func createGame(players int) gameState {
	game := gameState{currentMarble: nil, nextPlayer: 1, nextMarble: int64(1), marbles: list.New(), scores: make([]int64, players)}
	game.currentMarble = game.marbles.PushFront(int64(0))
	return game
}

func printGame(game *gameState) {
	digits := 0
	n := game.nextMarble
	for n > 0 {
		digits++
		n /= 10
	}

	fmt.Printf("[%d]", game.nextPlayer)
	for e := game.marbles.Front(); e != nil; e = e.Next() {
		if game.currentMarble == e {
			fmt.Printf("(%d)", e.Value.(int64))
		} else {
			fmt.Printf(" %d ", e.Value.(int64))
		}
	}
	fmt.Print("\n")
}

func moveCurrent(game *gameState, delta int64) {
	for i := int64(0); i != delta; {
		if delta < 0 {
			// Move backwards
			game.currentMarble = game.currentMarble.Prev()
			if game.currentMarble == nil {
				game.currentMarble = game.marbles.Back()
			}
			i--
		} else {
			game.currentMarble = game.currentMarble.Next()
			if game.currentMarble == nil {
				game.currentMarble = game.marbles.Front()
			}
			i++
		}
	}
}

func removeCurrent(game *gameState) {
	next := game.currentMarble.Next()
	game.marbles.Remove(game.currentMarble)
	game.currentMarble = next
}

func insertRightOfCurrent(game *gameState, val int64) {
	game.marbles.InsertAfter(val, game.currentMarble)
}

func advanceTurn(game *gameState) {
	if game.nextMarble%23 == 0 {
		// This is an interesting turn.  First, score this marble
		game.scores[game.nextPlayer] += game.nextMarble
		// Second, remove and score the marble 7 CCW from the current marble.  The marble CW of it (6 CCW from the current)
		// becomes the new current marble
		moveCurrent(game, -7)
		if debug {
			fmt.Printf("Scoring Marble index %p (%d) for player %d\n", game.currentMarble, game.currentMarble.Value.(int64), game.nextPlayer)
			fmt.Println(game.marbles)
		}
		game.scores[game.nextPlayer] += game.currentMarble.Value.(int64)
		removeCurrent(game)
	} else {
		// This is the standard turn.  Place the next marble two clockwise from the current marble
		moveCurrent(game, 1)
		insertRightOfCurrent(game, game.nextMarble)
		// Make the newly-inserted marble the current
		moveCurrent(game, 1)
	}

	// Advance to the next turn
	game.nextPlayer++
	game.nextPlayer = game.nextPlayer % len(game.scores)
	game.nextMarble++
}

// HighScore runs the game to the desired final point score and returns the highest player score
func HighScore(players int, points int64) int64 {
	game := createGame(players)
	for game.nextMarble <= points {
		advanceTurn(&game)
		if debugPrintTree {
			printGame(&game)
		}
	}
	return bestScore(game.scores)
}

func bestScore(scores []int64) int64 {
	best := scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] > best {
			best = scores[i]
		}
	}
	return best
}

func main() {
	file, err := os.Open("2018/09/input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	players, points := ReadInput(input[0])
	fmt.Println(HighScore(players, points))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(HighScore(players, points*100))
	fmt.Println(time.Since(start))
}
