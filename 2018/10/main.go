package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Projectile describes an object with a position and velocity
type Projectile struct {
	x  int
	y  int
	dx int // X velocity
	dy int // y velocity
}

// ReadProjectile converts a line from the input file into a Projectile object
func ReadProjectile(input string) Projectile {
	x, err := strconv.Atoi(strings.TrimSpace(input[10:16]))
	check(err)
	y, err := strconv.Atoi(strings.TrimSpace(input[18:24]))
	check(err)
	dx, err := strconv.Atoi(strings.TrimSpace(input[36:38]))
	check(err)
	dy, err := strconv.Atoi(strings.TrimSpace(input[40:42]))
	check(err)
	return Projectile{x: x, y: y, dx: dx, dy: dy}
}

func bounds(proj []Projectile) (int, int, int, int) {
	minx := proj[0].x
	maxx := proj[0].x
	miny := proj[0].y
	maxy := proj[0].y
	for _, p := range proj {
		minx = min(minx, p.x)
		miny = min(miny, p.y)
		maxx = max(maxx, p.x)
		maxy = max(maxy, p.y)
	}
	return minx, miny, maxx, maxy
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// TimeToConvergence gives the number of seconds until the
func TimeToConvergence(realProj []Projectile) int {
	// Make a copy to experiment with
	proj := make([]Projectile, len(realProj))
	copy(proj, realProj)
	time := 0
	minX, minY, maxX, maxY := bounds(proj)
	lastWidth := maxX - minX
	lastHeight := maxY - minY
	for {
		advance(proj, 1)
		minX, minY, maxX, maxY = bounds(proj)
		width := maxX - minX
		height := maxY - minY
		if (height > lastHeight) || (width > lastWidth) {
			return time
		}
		lastHeight = height
		lastWidth = width
		time++
	}
}

// Advances the projectiles the desired number of seconds
func advance(proj []Projectile, time int) {
	for t := 0; t < time; t++ {
		for i := 0; i < len(proj); i++ {
			proj[i].x += proj[i].dx
			proj[i].y += proj[i].dy
		}
	}
}

func drawProjectiles(proj []Projectile) {
	minX, minY, maxX, maxY := bounds(proj)
	normProj := normalize(proj, minX, minY)
	normX := maxX - minX
	normY := maxY - minY
	var grid = make(map[int]map[int]bool)
	for x := 0; x <= normX; x++ {
		grid[x] = make(map[int]bool)
		for y := 0; y <= normY; y++ {
			grid[x][y] = false
		}
	}
	for _, np := range normProj {
		grid[np.x][np.y] = true
	}

	for y := 0; y <= normY; y++ {
		for x := 0; x <= normX; x++ {
			if !grid[x][y] {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func normalize(proj []Projectile, minX, minY int) []Projectile {
	var norm []Projectile
	for _, p := range proj {
		norm = append(norm, Projectile{x: p.x - minX, y: p.y - minY, dx: p.dx, dy: p.dy})
	}
	return norm
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []Projectile
	for scanner.Scan() {
		input = append(input, ReadProjectile(scanner.Text()))
	}
	check(scanner.Err())
	start := time.Now()
	convergence := TimeToConvergence(input)
	fmt.Println(convergence)
	advance(input, convergence)
	drawProjectiles(input)
	fmt.Println(time.Since(start))
}
