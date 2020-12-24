package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// Assemble a jigsaw puzzle printed on transparency paper and look at the picture

// Enable debug output
const debug = false

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	tiles := parseTiles(lines)
	fmt.Println(step1(tiles))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(tiles))
	fmt.Println(sw.Elapsed())
}

// valuePair represents the two possible representations of a set of bits
// (current order or reversed)
type valuePair [2]int64

func (v valuePair) current() int64 {
	return v[0]
}

func (v valuePair) flipped() int64 {
	return v[1]
}

func (v valuePair) flip() {
	v[0], v[1] = v[1], v[0]
}

func (v valuePair) min() int64 {
	return aoc.Min(v[0], v[1])
}

func (v valuePair) compatible(v2 valuePair) bool {
	return v[0] == v2[0] || v[0] == v2[1]
}

type tile struct {
	id                    int64    // The identifier of the tile
	grid                  [][]bool // The pixels that make up the tile
	up, down, left, right valuePair
	used                  bool // Has this tile been assembled in a picture?
}

func newTile(id int64, lines []string) *tile {
	t := tile{id: id, grid: make([][]bool, 0), used: false}
	for i, l := range lines {
		if len(l) != len(lines) {
			log.Fatalf("Unexpected line: %q\n", l)
		}
		t.grid = append(t.grid, make([]bool, 0))
		for _, v := range []rune(l) {
			t.grid[i] = append(t.grid[i], v == '#')
		}
	}
	t.calculateSides()
	return &t
}

func newTileFromGrid(id int64, grid [][]bool) *tile {
	for _, r := range grid {
		if len(r) != len(grid) {
			log.Fatal("Invalid tile grid")
		}
	}
	t := tile{id: id, grid: grid, used: false}
	t.calculateSides()
	return &t
}

func binaryAndReversed(s string) valuePair {
	forward, _ := strconv.ParseInt(s, 2, 64)
	bits := []rune(s)
	ls := len(s)
	for i := 0; i < (ls / 2); i++ {
		bits[i], bits[ls-(1+i)] = bits[ls-(1+i)], bits[i]
	}
	reverse, _ := strconv.ParseInt(string(bits), 2, 64)
	return [2]int64{forward, reverse}
}

func bit(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func (t *tile) calculateSides() {
	u, d, l, r := "", "", "", ""
	for _, b := range t.grid[0] {
		u += bit(b)
	}
	t.up = binaryAndReversed(u)
	for _, b := range t.grid[len(t.grid)-1] {
		d += bit(b)
	}
	t.down = binaryAndReversed(d)
	for _, row := range t.grid {
		l += bit(row[0])
		r += bit(row[len(row)-1])
	}
	t.left = binaryAndReversed(l)
	t.right = binaryAndReversed(r)
}

func (t *tile) size() int {
	return len(t.grid)
}

func (t *tile) cw() {
	lg := len(t.grid)
	lr := len(t.grid[0])
	rot := make([][]bool, lr)
	for i := range t.grid {
		rot[i] = make([]bool, lg)
	}

	for y, r := range t.grid {
		for x, val := range r {
			rot[x][lr-1-y] = val
		}
	}

	t.grid = rot
	t.right, t.down, t.left, t.up = t.up, t.right, t.down, t.left
}

func (t *tile) ccw() {
	t.cw()
	t.cw()
	t.cw()
}

func (t *tile) flipVertical() {
	for x := range t.grid[0] {
		lg := len(t.grid)
		for y := 0; y < (lg / 2); y++ {
			t.grid[y][x], t.grid[lg-1-y][x] = t.grid[lg-1-y][x], t.grid[y][x]
		}
	}
	t.left.flip()
	t.right.flip()
	t.up, t.down = t.down, t.up
}

func (t *tile) flipHorizontal() {
	for y := range t.grid {
		ly := len(t.grid[y])
		for x := 0; x < (ly / 2); x++ {
			t.grid[y][x], t.grid[y][ly-1-x] = t.grid[y][ly-1-x], t.grid[y][x]
		}
	}
	t.up.flip()
	t.down.flip()
	t.left, t.right = t.right, t.left
}

func (t *tile) trim() {
	t.grid = t.grid[1 : len(t.grid)-1]
	for i := range t.grid {
		t.grid[i] = t.grid[i][1 : len(t.grid[i])-1]
	}
}

func (t *tile) side(dir aoc.Direction) valuePair {
	switch dir {
	case aoc.Up:
		return t.up
	case aoc.Down:
		return t.down
	case aoc.Left:
		return t.left
	case aoc.Right:
		return t.right
	}
	log.Fatalf("Invalid direction: %d", int64(dir))
	return valuePair{}
}

func (t *tile) rotateExteriorEdges(dir []aoc.Direction, tileEdges map[int64][]*tile) {
	for matchedAll := false; !matchedAll; {
		matchedAll = true
		for _, d := range dir {
			matchedAll = matchedAll && len(tileEdges[t.side(d).min()]) == 1
		}
		if matchedAll {
			return
		}
		t.cw()
	}
}

func (t *tile) rotateValueTo(val valuePair, dir aoc.Direction) {
	for !t.side(dir).compatible(val) {
		t.cw()
	}
}

type tiles map[int64]*tile

var (
	inputRegex = regexp.MustCompile(`^Tile (\d+):$`)
)

func parseTiles(lines []string) tiles {
	t := make(tiles)
	for i := 0; i < len(lines); i += 12 {
		tokens := inputRegex.FindStringSubmatch(lines[i])
		id, _ := strconv.ParseInt(tokens[1], 10, 64)
		t[id] = newTile(id, lines[i+1:i+11])
	}
	return t
}

// findEdges returns the edges and corners, as well as a list of all tiles with a certain edge
func (ts tiles) findEdges() (tiles, tiles, map[int64][]*tile) {
	dimensions := int(math.Sqrt(float64(len(ts))))
	if dimensions*dimensions != len(ts) {
		log.Fatalf("Not a square grid")
	}

	sq := make([][]int64, dimensions)
	for i := range sq {
		sq[i] = make([]int64, dimensions)
		for j := range sq[i] {
			sq[i][j] = -1
		}
	}

	// Find unique patterns and mark them as the edges
	tileEdges := make(map[int64][]*tile)
	for _, t := range ts {
		for _, s := range []valuePair{t.up, t.left, t.right, t.down} {
			val := s.min()
			if tileEdges[val] == nil {
				tileEdges[val] = make([]*tile, 0)
			}
			tileEdges[val] = append(tileEdges[val], t)
		}
	}

	edges := make(map[int64]*tile)
	corners := make(map[int64]*tile)
	for e, te := range tileEdges {
		if len(te) == 1 {
			t := te[0]
			if _, found := edges[t.id]; found {
				corners[t.id] = t
			}
			edges[t.id] = t
		}
		if debug && len(te) > 2 {
			// See if all common edges are unique to a pair of tiles.
			// Testing on the input indicates they are
			fmt.Printf("Found %d hits for pattern %d (tiles ", len(te), e)
			for _, t := range te {
				fmt.Printf("%d ", t.id)
			}
			fmt.Println(")")
		}
	}
	for id := range corners {
		delete(edges, id)
	}

	return edges, corners, tileEdges
}

func (ts tiles) assemble() *tile {
	// To assemble the picture, it must be a square
	dim := int(math.Sqrt(float64(len(ts))))
	if dim*dim != len(ts) || len(ts) < 1 {
		log.Fatalf("Invalid grid size: %d", len(ts))
	}

	// Create a 2D grid to place the tiles in
	pic := make([][]*tile, dim)
	for i := range pic {
		pic[i] = make([]*tile, dim)
	}

	_, corners, tileEdges := ts.findEdges()
	// Pick an arbitrary corner
	for _, c := range corners {
		pic[0][0] = c
		pic[0][0].used = true
		break
	}

	// Helper values for rotation
	toUpLeft := []aoc.Direction{aoc.Up, aoc.Left}
	toUp := []aoc.Direction{aoc.Up}
	toUpRight := []aoc.Direction{aoc.Up, aoc.Right}
	toLeft := []aoc.Direction{aoc.Left}
	toRight := []aoc.Direction{aoc.Right}
	toDownLeft := []aoc.Direction{aoc.Down, aoc.Left}
	toDown := []aoc.Direction{aoc.Down}
	toDownRight := []aoc.Direction{aoc.Down, aoc.Right}

	// Rotate the corner so that its unique edges are at the top and the left
	pic[0][0].rotateExteriorEdges(toUpLeft, tileEdges)

	// Fill in the top row
	for x := 1; x < dim; x++ {
		found := false
		target := pic[0][x-1].right
		for _, t := range tileEdges[target.min()] {
			if t.used {
				continue
			}

			if x < (dim - 1) {
				t.rotateExteriorEdges(toUp, tileEdges)
				if !t.left.compatible(target) {
					t.flipHorizontal()
					if !t.left.compatible(target) {
						log.Fatalf("(1) Cannot place edge %d next to tile %d", t.id, pic[0][x-1].id)
					}
				}
			} else {
				// This is a corner
				t.rotateExteriorEdges(toUpRight, tileEdges)
				if !t.left.compatible(target) {
					t.flipHorizontal()
					t.rotateExteriorEdges(toUpRight, tileEdges)
					if !t.left.compatible(target) {
						log.Fatalf("(2) Cannot place corner %d next to edge %d", t.id, pic[0][x-1].id)
					}
				}
			}

			pic[0][x] = t
			t.used = true
			found = true
		}
		if !found {
			log.Fatalf("Unable to place tile at Y=0 X=%d", x)
		}
	}

	for y := 1; y < dim-1; y++ {
		// Fill in the left edge
		found := false
		target := pic[y-1][0].down
		for _, t := range tileEdges[target.min()] {
			if t.used {
				continue
			}

			t.rotateExteriorEdges(toLeft, tileEdges)
			if !t.up.compatible(target) {
				t.flipVertical()
				if !t.up.compatible(target) {
					log.Fatalf("(3) Cannot place edge %d next to tile %d", t.id, pic[y-1][0].id)
				}
			}

			pic[y][0] = t
			t.used = true
			found = true
		}
		if !found {
			log.Fatalf("Unable to place tile at Y=%d X=0", y)
		}

		// Fill in the right edge
		found = false
		target = pic[y-1][dim-1].down
		for _, t := range tileEdges[target.min()] {
			if t.used {
				continue
			}

			t.rotateExteriorEdges(toRight, tileEdges)
			if !t.up.compatible(target) {
				t.flipVertical()
				if !t.up.compatible(target) {
					log.Fatalf("(4) Cannot place edge %d next to tile %d", t.id, pic[y-1][dim-1].id)
				}
			}

			pic[y][dim-1] = t
			t.used = true
			found = true
		}
		if !found {
			log.Fatalf("Unable to place tile at Y=%d X=%d", y, dim-1)
		}
	}

	// Fill in the bottom row
	for x := 0; x < dim; x++ {
		y := dim - 1
		found := false
		if x == 0 {
			// DL corner
			up := pic[y-1][x].down
			for _, t := range tileEdges[up.min()] {
				if t.used {
					continue
				}
				t.rotateExteriorEdges(toDownLeft, tileEdges)
				if !t.up.compatible(up) {
					t.flipHorizontal()
					t.rotateExteriorEdges(toDownLeft, tileEdges)
					if !t.up.compatible(up) {
						log.Fatalf("(5) Cannot place corner %d next to tile %d", t.id, pic[y-1][x].id)
					}
				}
				pic[y][x] = t
				t.used = true
				found = true
			}
		} else if x == dim-1 {
			// DR corner
			up := pic[y-1][x].down
			for _, t := range tileEdges[up.min()] {
				if t.used {
					continue
				}
				t.rotateExteriorEdges(toDownRight, tileEdges)
				if !t.up.compatible(up) {
					t.flipHorizontal()
					t.rotateExteriorEdges(toDownRight, tileEdges)
					if !t.up.compatible(up) {
						log.Fatalf("(6) Cannot place corner %d next to tile %d", t.id, pic[y-1][x].id)
					}
				}
				pic[y][x] = t
				t.used = true
				found = true
			}
		} else {
			// Bottom edge, Interior
			left := pic[y][x-1].right
			for _, t := range tileEdges[left.min()] {
				if t.used {
					continue
				}
				t.rotateExteriorEdges(toDown, tileEdges)
				if !t.left.compatible(left) {
					t.flipHorizontal()
					if !t.left.compatible(left) {
						log.Fatalf("(7) Cannot place edge %d next to tile %d", t.id, pic[y][x-1].id)
					}
				}
				pic[y][x] = t
				t.used = true
				found = true
			}
		}
		if !found {
			log.Fatalf("Unable to place tile at Y=%d X=%d", y, x)
		}
	}

	// Fill in the interior
	for y := 1; y < dim-1; y++ {
		for x := 1; x < dim-1; x++ {
			up := pic[y-1][x].down
			left := pic[y][x-1].right
			for _, t := range tileEdges[up.min()] {
				if t.used {
					continue
				}
				t.rotateValueTo(up, aoc.Up)
				if !t.left.compatible(left) {
					t.flipHorizontal()
					if !t.left.compatible(left) {
						log.Fatalf("(8) Cannot place edge %d next to tile %d", t.id, pic[y][x-1].id)
					}
				}
				pic[y][x] = t
				t.used = true
			}
		}
	}

	if debug {
		for _, row := range pic {
			for _, t := range row {
				fmt.Printf("%d\t", t.id)
			}
			fmt.Println("")
		}
	}

	// Create a new combined picture
	cellDim := pic[0][0].size() - 2
	fullDim := cellDim * dim
	grid := make([][]bool, fullDim)
	for i := range grid {
		grid[i] = make([]bool, fullDim)
	}
	for yy, row := range pic {
		for xx, t := range row {
			// Trim off the alignment edges
			for y := 1; y < len(t.grid)-1; y++ {
				for x := 1; x < len(t.grid[y])-1; x++ {
					grid[yy*cellDim+y-1][xx*cellDim+x-1] = t.grid[y][x]
				}
			}
		}
	}

	// Return the combined picture
	t := newTileFromGrid(0, grid)
	return t
}

var seaMonster = [][]rune{
	[]rune("                  # "),
	[]rune("#    ##    ##    ###"),
	[]rune(" #  #  #  #  #  #   "),
}

func findSeaMonsters(picture [][]bool) map[aoc.Point]bool {
	monsters := make(map[aoc.Point]bool)
	for y := 0; y <= len(picture)-len(seaMonster); y++ {
	PICTURE:
		for x := 0; x <= len(picture[y])-len(seaMonster[0]); x++ {
			for yd := 0; yd < len(seaMonster); yd++ {
				for xd := 0; xd < len(seaMonster[0]); xd++ {
					if seaMonster[yd][xd] != '#' {
						continue
					}
					if !picture[y+yd][x+xd] {
						// This is not a valid sea monster
						continue PICTURE
					}
				} // xd
			} // yd
			// This is a valid sea monster.  Loop back through and mark all of the spots
			if debug {
				fmt.Printf("Found monster at Y=%d,X=%d\n", y, x)
			}
			for yd := 0; yd < len(seaMonster); yd++ {
				for xd := 0; xd < len(seaMonster[0]); xd++ {
					if seaMonster[yd][xd] == '#' {
						monsters[aoc.Point{X: int64(x + xd), Y: int64(y + yd)}] = true
					}
				} // xd
			} // yd
		} // x
	} // y
	return monsters
}

func step1(t tiles) int64 {
	_, corners, _ := t.findEdges()
	product := int64(1)
	for id := range corners {
		product *= id
	}
	return product
}

func step2(ts tiles) int64 {
	picture := ts.assemble()
	monsters := make(map[aoc.Point]bool)
	for it := 0; true; it++ {
		monsters = findSeaMonsters(picture.grid)
		if len(monsters) > 0 {
			if debug {
				fmt.Printf("Found %d monster squares on iteration %d\n", len(monsters), it)
			}
			break
		}
		picture.cw()
		it++
		if it == 4 || it == 12 {
			picture.flipHorizontal()
		} else if it == 8 {
			picture.flipVertical()
		}
		if it > 16 {
			log.Fatalf("Found no monsters")
		}
	}
	count := int64(0)
	for y, r := range picture.grid {
		for x, c := range r {
			if c && !monsters[aoc.Point{X: int64(x), Y: int64(y)}] {
				count++
			}
		}
	}
	return count
}
