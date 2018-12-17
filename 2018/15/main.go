package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const debug = false

// spot is an "enum" of Wall/Cavern
type spot rune

// unitType is an "enum" of Goblin/Elf
type unitType rune

// I already told you what these are but golint wants a comment here
const (
	Wall   = '#'
	Cavern = '.'

	Goblin = 'G'
	Elf    = 'E'
)

// We want to modify exactly one set of units and never copy them,
// so pass around pointers
type Units []*Unit

// Cave represents the units in a cave network
type Cave struct {
	layout caveLayout
	units  Units
	width  int
	height int
}

// Unit represents a goblin or elf
type Unit struct {
	kind   unitType
	hp     int
	attack int
	id     int
	x      int
	y      int
}

var nextUnitID = 0

type caveLayout map[int]map[int]spot // Maps X->Y->spot
type unitLocationMap map[int]map[int]*Unit
type unitIDMap map[int]*Unit

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func makeUnit(kind unitType, hp, attack, x, y int) Unit {
	nextUnitID++
	return Unit{kind: kind, hp: hp, attack: attack, id: nextUnitID, x: x, y: y}
}

// Returns a map of unit location to the unit
func makeUnitLocationMap(c *Cave) unitLocationMap {
	uMap := make(unitLocationMap)
	for i := 0; i < c.width; i++ {
		uMap[i] = make(map[int]*Unit)
	}
	for _, u := range c.units {
		uMap[u.x][u.y] = u
	}
	return uMap
}

func makeUnitIDMap(units Units) unitIDMap {
	uMap := make(unitIDMap)
	for _, u := range units {
		uMap[u.id] = u
	}
	return uMap
}

func makeGoblin(x, y int) Unit {
	return makeUnit(Goblin, 200, 3, x, y)
}

func makeElf(x, y int) Unit {
	return makeUnit(Elf, 200, 3, x, y)
}

func alive(u Unit) bool {
	return u.hp > 0
}

// ReadCave parses the input into a cave network
func ReadCave(input []string) Cave {
	layout := make(caveLayout)
	units := make(Units, 0)
	for y, s := range input {
		for x, c := range []rune(s) {
			if _, exists := layout[x]; !exists {
				layout[x] = make(map[int]spot)
			}

			switch c {
			case Wall, Cavern:
				layout[x][y] = spot(c)

			case Elf:
				layout[x][y] = Cavern
				elf := makeElf(x, y)
				units = append(units, &elf)

			case Goblin:
				layout[x][y] = Cavern
				goblin := makeGoblin(x, y)
				units = append(units, &goblin)

			default:
				log.Fatalf("Unknown character in input: %c", c)
			}
		}
	}
	return Cave{
		layout: layout,
		units:  units,
		width:  len(input[0]),
		height: len(input)}
}

func printCave(c *Cave) {
	uMap := makeUnitLocationMap(c)
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if _, exists := uMap[x][y]; exists {
				fmt.Printf("%c", rune(uMap[x][y].kind))
			} else {
				fmt.Printf("%c", c.layout[x][y])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Outcome returns the number of rounds multiplied by the hit points
// of the winning team
func Outcome(input []string) int {
	cave := ReadCave(input)
	if debug {
		printCave(&cave)
	}

	for round := 0; ; round++ {
		done := performRound(&cave)
		if debug {
			printCave(&cave)
		}
		if done {
			return round
		}
	}

	return 0
}

// Where the magic happens.  For a full description of the process,
// see the README
func performRound(c *Cave) bool {
	turnOrder := sortedByPosition(c.units)

	for _, currentUnit := range turnOrder {
		// Phase 0: Count enemies.  If none exist, return true
		if len(enemies(currentUnit, c.units)) == 0 {
			return true
		}

		// Phase 1: Movement
		// If in adjacent to enemy, do not move
		unitLocations := makeUnitLocationMap(c)
		adjacent := findAdjacent(currentUnit, &unitLocations)
		if len(enemies(currentUnit, adjacent)) < 1 {
			// Identify all open squares adjacent to all enemies
			// If no open squares adjacent to enemies, do not move
			// Djikstra distance to all of the in-range squares
			// Pick the square that could be moved to in fewest steps
			// Move one space towards that square
		}

		// Phase 2: Attack
		// If no target in range, end turn
		adjacent = findAdjacent(currentUnit, &unitLocations)
		targetCandidates := enemies(currentUnit, adjacent)
		if len(targetCandidates) >= 1 {
			// Take target w/ lowest HP, tiebreak in READING ORDER
			// If unit killed, remove from units list
		}
	}

	// TODO: Once ready to test, uncomment
	//return false
	return true
}

func enemies(current *Unit, others Units) Units {
	enemy := make(Units, 0)
	for _, u := range others {
		if u.kind != current.kind {
			enemy = append(enemy, u)
		}
	}
	return enemy
}

func findAdjacent(current *Unit, locations *unitLocationMap) Units {

}

func sortedByPosition(units Units) Units {
	sorted := make(Units, 0)
	for _, u := range units {
		sorted = append(sorted, u)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].y < sorted[j].y {
			return true
		} else if sorted[i].y > sorted[j].y {
			return false
		}
		return sorted[i].x < sorted[j].x
	})
	return sorted
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
	fmt.Println(Outcome(input))
	fmt.Println(time.Since(start))
}
