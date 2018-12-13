package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const debug = false

// Track is an "enum" of Vertical/Horizontal/Intersection/Backslash/Slash/Empty
// Yes, this language has no concept of an enum.  I feel like I'm writing QBASIC again.
type Track rune

// Direction is an "enum" of Up/Left/Right/Down
type Direction rune

// I already told you what these are but golint wants a comment here
const (
	Vertical     = '|'
	Horizontal   = '-'
	Intersection = '+'
	Backslash    = '\\'
	Slash        = '/'
	Empty        = ' '

	Up    = '^'
	Left  = '<'
	Right = '>'
	Down  = 'v'
)

// RailCart tracks a cart's position and direction
type RailCart struct {
	currentX  int
	currentY  int
	previousX int
	previousY int
	id        int
	dir       Direction
	// How many times the cart has turned.  At any intersection, for turnCount % 3:
	// 0 -> Turn left, 1 -> Straight, 2 -> Right.  Then increment turnCount
	turnCount int
	alive     bool
}

// RailMap represents the grid of tracks
type RailMap map[int]map[int]Track // Maps X->Y->Track

// CartMap represents the grid of where every cart is
type CartMap map[int]map[int][]RailCart // Maps X->Y->Track

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadTrackNetwork parses the input into a track network
func ReadTrackNetwork(input []string) (RailMap, CartMap) {
	rails := make(RailMap)
	carts := make(CartMap)
	for y, s := range input {
		for x, c := range []rune(s) {
			if _, exists := rails[x]; !exists {
				rails[x] = make(map[int]Track)
			}

			switch c {
			case Vertical, Horizontal, Intersection, Backslash, Slash, Empty:
				rails[x][y] = Track(c)

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
	return rails, carts
}

// FirstCollision finds the first time two rail carts collide and returns the location
// Returns x, y, time
func FirstCollision(rails *RailMap, originalCarts *CartMap) (int, int, int) {
	carts := *originalCarts
	for time := 1; ; time++ {
		if debug {
			fmt.Printf("Time: %d\n", time)
		}
		newCarts := advance(rails, &carts)

		// First, check for a direct same-square collision
		for x := range newCarts {
			for y := range newCarts[x] {
				if len(newCarts[x][y]) > 1 {
					return x, y, time
				}
			}
		}

		// Then, check for two carts which swapped locations (and thus crashed in the middle)
		for x := range newCarts {
			for y, carts := range newCarts[x] {
				for _, cart := range carts {
					// Look for a cart in the previous location
					if _, present := newCarts[cart.previousX]; present {
						if _, exists := newCarts[cart.previousX][cart.previousY]; exists {
							for _, oldCart := range newCarts[cart.previousX][cart.previousY] {
								if (oldCart.previousX == cart.currentX) && (oldCart.previousY == cart.currentY) {
									return x, y, time
								}
							}
						}
					}
				}
			}
		}
		carts = newCarts
	}
}

// LastCartStanding finds the local of the final cart still alive after the others collide
func LastCartStanding(rails *RailMap, originalCarts *CartMap) (int, int, int) {
	carts := *originalCarts
	for time := 1; ; time++ {
		if debug {
			fmt.Printf("Time: %d\n", time)
		}
		newCarts := advance(rails, &carts)

		// First, check for a direct same-square collision
		for x := range newCarts {
			for y := range newCarts[x] {
				if len(newCarts[x][y]) > 1 {
					for i := 0; i < len(newCarts[x][y]); i++ {
						newCarts[x][y][i].alive = false
					}
				}
			}
		}

		// Then, check for two carts which swapped locations (and thus crashed in the middle)
		for x := range newCarts {
			for y, carts := range newCarts[x] {
				for newIndex, cart := range carts {
					// Look for a cart in the previous location
					if _, present := newCarts[cart.previousX]; present {
						if _, exists := newCarts[cart.previousX][cart.previousY]; exists {
							for oldIndex, oldCart := range newCarts[cart.previousX][cart.previousY] {
								if (oldCart.previousX == cart.currentX) && (oldCart.previousY == cart.currentY) {
									newCarts[x][y][newIndex].alive = false
									newCarts[cart.previousX][cart.previousY][oldIndex].alive = false
								}
							}
						}
					}
				}
			}
		}

		aliveX, aliveY, aliveCount := -1, -1, 0
		for x := range newCarts {
			for y, carts := range newCarts[x] {
				for _, cart := range carts {
					if cart.alive {
						aliveX, aliveY = x, y
						aliveCount++
					}
				}
			}
		}

		if aliveCount == 1 {
			return aliveX, aliveY, time
		}

		carts = newCarts
	}
}

func turnLeft(d Direction) Direction {
	switch d {
	case Up:
		return Left
	case Left:
		return Down
	case Down:
		return Right
	case Right:
		return Up
	}
	log.Fatalf("Unknown direction to turnLeft: %c", d)
	return Up
}

func turnRight(d Direction) Direction {
	switch d {
	case Up:
		return Right
	case Left:
		return Up
	case Down:
		return Left
	case Right:
		return Down
	}
	log.Fatalf("Unknown direction to turnRight: %c", d)
	return Up
}

func advance(rails *RailMap, originalCarts *CartMap) CartMap {
	newCarts := make(CartMap)
	for x, row := range *originalCarts {
		for y, carts := range row {
			for _, c := range carts {
				if c.alive {
					cart := RailCart(c)
					cart.previousX = cart.currentX
					cart.previousY = cart.currentY
					// Determine the new direction
					switch (*rails)[x][y] {
					case Vertical:
						if (cart.dir != Up) && (cart.dir != Down) {
							log.Fatalf("Cart on vertical rail with non-vertical direction: %d,%d %d", c.currentX, c.currentY, c.id)
						}
					case Horizontal:
						if (cart.dir != Left) && (cart.dir != Right) {
							log.Fatalf("Cart on horizontal rail with non-horizontal direction %d,%d %d", c.currentX, c.currentY, c.id)
						}
					case Backslash:
						switch cart.dir {
						case Left:
							cart.dir = Up
						case Right:
							cart.dir = Down
						case Up:
							cart.dir = Left
						case Down:
							cart.dir = Right
						}
					case Slash:
						switch cart.dir {
						case Left:
							cart.dir = Down
						case Right:
							cart.dir = Up
						case Up:
							cart.dir = Right
						case Down:
							cart.dir = Left
						}
					case Intersection:
						switch cart.turnCount % 3 {
						case 0:
							cart.dir = turnLeft(c.dir)
						// case 1: go straight
						case 2:
							cart.dir = turnRight(c.dir)
						}
						cart.turnCount++
					case Empty:
						log.Fatalf("Cart on empty space %d,%d %d", cart.currentX, cart.currentY, cart.id)
					default:
						log.Fatalf("Cart on unknown space %d,%d %d", cart.currentX, cart.currentY, cart.id)
					}

					switch cart.dir {
					case Up:
						cart.currentY--
					case Down:
						cart.currentY++
					case Left:
						cart.currentX--
					case Right:
						cart.currentX++
					default:
						log.Fatalf("Unknown cart direction: %d,%d %d %c", cart.currentX, cart.currentY, cart.id, cart.dir)
					}

					if debug {
						fmt.Printf("Cart %d now at (%d, %d) travelling %c\n", cart.id, cart.currentX, cart.currentY, rune(cart.dir))
					}

					// Add the cart to the new list
					if _, present := newCarts[cart.currentX]; !present {
						newCarts[cart.currentX] = make(map[int][]RailCart)
					}
					if _, present := newCarts[cart.currentX][cart.currentY]; !present {
						newCarts[cart.currentX][cart.currentY] = make([]RailCart, 0)
					}
					newCarts[cart.currentX][cart.currentY] = append(newCarts[cart.currentX][cart.currentY], cart)
				}
			}
		}
	}
	return newCarts
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
	rails, carts := ReadTrackNetwork(input)
	x, y, t := FirstCollision(&rails, &carts)
	fmt.Printf("First Collision at (%d, %d) at time %d\n", x, y, t)
	fmt.Println(time.Since(start))
	start = time.Now()
	x, y, t = LastCartStanding(&rails, &carts)
	fmt.Printf("Last Cart at (%d, %d) at time %d\n", x, y, t)
	fmt.Println(time.Since(start))
}
