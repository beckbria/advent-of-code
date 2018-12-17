package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	Unreachable = math.MaxInt32
)

// Units should never be copied so pass around pointers
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

type point struct {
	x int
	y int
}

// path represents the distance to a space and the space before it in
// the path
type path struct {
	distance  int
	preceding point
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

func alive(u *Unit) bool {
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
}

// Compute the shortest path to all reachable squares on the map
func shortestPaths(unit *Unit, unitLocations unitLocationMap, cave *Cave) map[int]map[int]path {
	paths := make(map[int]map[int]path)
	for i := 0; i < cave.width; i++ {
		paths[i] = make(map[int]path)
		for j := 0; j < cave.height; j++ {
			paths[i][j] = path{distance: Unreachable}
		}
	}

	// TODO: Djikstra's Algorithm

	return paths
}

// Where the magic happens.  For a full description of the process,
// see the README
func performRound(cave *Cave) bool {
	turnOrder := sortedByPosition(cave.units)

	for _, currentUnit := range turnOrder {
		// Phase 0: Count enemies.  If none exist, return true
		allEnemies := enemies(currentUnit, cave.units)
		if len(allEnemies) == 0 {
			return true
		}

		// Phase 1: Movement
		// If in adjacent to enemy, do not move
		unitLocations := makeUnitLocationMap(cave)
		adjacentUnits, _ := findAdjacent(currentUnit, unitLocations, cave)
		if len(enemies(currentUnit, adjacentUnits)) < 1 {
			// Identify all open squares adjacent to all enemies
			destinations := make([]point, 0)
			for _, e := range allEnemies {
				_, adjacent := findAdjacent(e, unitLocations, cave)
				destinations = append(destinations, adjacent...)
			}
			// If no open squares adjacent to enemies, do not move
			if len(destinations) > 0 {
				// Djikstra distance to all of the in-range squares
				distances := shortestPaths(currentUnit, unitLocations, cave)
				// Pick the square that could be moved to in fewest steps
				target := destinations[0]
				for _, d := range destinations {
					dt := distances[target.x][target.y].distance
					dd := distances[d.x][d.y].distance
					if dd < dt {
						target = d
					} else if dd == dt {
						ud := Unit{x: d.x, y: d.y}
						ut := Unit{x: target.x, y: target.y}
						if readOrderLess(&ud, &ut) {
							target = d
						}
					}
				}
				// Move one space towards that square
				// Walk the path backwards to find the first step
				p := distances[target.x][target.y]
				newLocation := point{x: target.x, y: target.y}
				for (p.preceding.x != currentUnit.x) || (p.preceding.y != currentUnit.y) {
					newLocation = p.preceding
					p = distances[p.preceding.x][p.preceding.y]
				}
				currentUnit.x, currentUnit.y = newLocation.x, newLocation.y
			}
		}

		// Phase 2: Attack
		// If no target in range, end turn
		adjacentUnits, _ = findAdjacent(currentUnit, unitLocations, cave)
		targetCandidates := enemies(currentUnit, adjacentUnits)
		if len(targetCandidates) >= 1 {
			// Take target w/ lowest HP, tiebreak in READING ORDER
			target := targetCandidates[0]
			for _, t := range targetCandidates {
				if (t.hp < target.hp) || ((t.hp == target.hp) && readOrderLess(t, target)) {
					target = t
				}
			}

			// Attack!
			target.hp -= currentUnit.attack

			// Bury the dead
			cave.units = aliveOnly(cave.units)
		}
	}

	// TODO: Once ready to test, uncomment
	//return false
	return true
}

// Removes any deceased units from a list
func aliveOnly(units Units) Units {
	living := make(Units, 0)
	for _, u := range units {
		if alive(u) {
			living = append(living, u)
		}
	}
	return living
}

func enemies(current *Unit, others Units) Units {
	enemy := make(Units, 0)
	for _, u := range others {
		if alive(u) && (u.kind != current.kind) {
			enemy = append(enemy, u)
		}
	}
	return enemy
}

func findAdjacent(current *Unit, locations unitLocationMap, cave *Cave) (Units, []point) {
	x := current.x
	y := current.y
	units := make(Units, 0)
	emptyPoints := make([]point, 0)
	candidates := []point{
		point{x: x, y: y - 1},
		point{x: x - 1, y: y},
		point{x: x + 1, y: y},
		point{x: x, y: y + 1}}

	for _, pt := range candidates {
		if (x >= 0) && (y <= 0) && (x < cave.width) && (y < cave.height) {
			if hasLocation(locations, pt.x, pt.y) {
				units = append(units, locations[pt.x][pt.y])
			} else {
				emptyPoints = append(emptyPoints, pt)
			}
		}
	}

	return units, emptyPoints
}

func hasLocation(locations unitLocationMap, x, y int) bool {
	if _, present := locations[x]; !present {
		return false
	}
	_, present := locations[x][y]
	return present
}

// Returns true if Unit a<b in reading order (left to right, top to bottom)
func readOrderLess(a, b *Unit) bool {
	if a.y < b.y {
		return true
	} else if a.y > b.y {
		return false
	}
	return a.x < b.x
}

func sortedByPosition(units Units) Units {
	sorted := make(Units, 0)
	for _, u := range units {
		sorted = append(sorted, u)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return readOrderLess(sorted[i], sorted[j])
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
