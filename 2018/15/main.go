package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const debug = false

// Track is an "enum" of Wall/Cavern/Goblin/Elf
// Yes, this language has no concept of an enum.  I feel like I'm writing QBASIC again.
type Spot rune

// I already told you what these are but golint wants a comment here
const (
	Wall   = '#'
	Cavern = '.'
	Goblin = 'G'
	Elf    = 'E'
)

// CaveMap represents the grid of tracks
type CaveMap map[int]map[int]Spot // Maps X->Y->Spot

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadCaveMap parses the input into a cave network
func ReadCaveMap(input []string) CaveMap {
	cave := make(RailMap)
	carts := make(CartMap)
	for y, s := range input {
		for x, c := range []rune(s) {
			if _, exists := rails[x]; !exists {
				rails[x] = make(map[int]Track)
			}

			switch c {
			case Wall, Cavern, Goblin, Elf:
				cave[x][y] = Spot(c)

			case Up, Left, Right, Down:
				if _, present := carts[x]; !present {
					carts[x] = make(map[int][]RailCart)
				}
				carts[x][y] = []RailCart{RailCart{
					currentX:  x,
					currentY:  y,
					previousX: -1,
					previousY: -1,
					alive:     true,
					dir:       Direction(c),
					id:        len(carts)}}
				if (c == Up) || (c == Down) {
					rails[x][y] = Vertical
				} else { // Left or Right
					rails[x][y] = Horizontal
				}

				if debug {
					fmt.Printf("Created Cart %d at (%d,%d) facing %c\n", carts[x][y][0].id, x, y, c)
				}

			default:
				log.Fatalf("Unknown character in input: %c", c)
			}
		}
	}
	return cave
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
	cave := ReadCaveMap(input)
	fmt.Println(input)
	fmt.Println(time.Since(start))
}
