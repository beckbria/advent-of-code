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

// Cave represents the units in a cave network
type Cave struct {
	layout  caveLayout
	elves   unitMap
	goblins unitMap
	width   int
	height  int
}

type caveLayout map[int]map[int]Spot // Maps X->Y->Spot
type unitMap map[int]map[int]bool

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadCave parses the input into a cave network
func ReadCave(input []string) Cave {
	layout := make(caveLayout)
	elves := make(unitMap)
	goblins := make(unitMap)
	for y, s := range input {
		for x, c := range []rune(s) {
			if _, exists := layout[x]; !exists {
				layout[x] = make(map[int]Spot)
			}
			if _, exists := elves[x]; !exists {
				elves[x] = make(map[int]bool)
			}
			if _, exists := goblins[x]; !exists {
				goblins[x] = make(map[int]bool)
			}

			switch c {
			case Wall, Cavern:
				layout[x][y] = Spot(c)

			case Elf:
				layout[x][y] = Cavern
				elves[x][y] = true

			case Goblin:
				layout[x][y] = Cavern
				goblins[x][y] = true

			default:
				log.Fatalf("Unknown character in input: %c", c)
			}
		}
	}
	return Cave{
		layout:  layout,
		elves:   elves,
		goblins: goblins,
		width:   len(input[0]),
		height:  len(input)}
}

func printCave(c Cave) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if _, exists := c.goblins[x][y]; exists {
				fmt.Printf("%c", rune(Goblin))
			} else if _, exists := c.elves[x][y]; exists {
				fmt.Printf("%c", rune(Elf))
			} else {
				fmt.Printf("%c", c.layout[x][y])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
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
	cave := ReadCave(input)
	printCave(cave)
	fmt.Println(time.Since(start))
}
