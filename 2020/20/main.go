package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"

	"../../aoc"
)

// https://adventofcode.com/2020/day/#
// TODO: Description

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	tiles := parseTiles(lines)
	fmt.Println(step1(tiles))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	//fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

type tile struct {
	id   int64
	grid [][]bool
	used bool
}

func binaryAndReversed(s string) []int64 {
	forward, _ := strconv.ParseInt(s, 2, 64)
	bits := []rune(s)
	ls := len(s)
	for i := 0; i < (ls / 2); i++ {
		bits[i], bits[ls-(1+i)] = bits[ls-(1+i)], bits[i]
	}
	reverse, _ := strconv.ParseInt(string(bits), 2, 64)
	return []int64{forward, reverse}
}

func (t *tile) top() []int64 {
	bits := ""
	for _, b := range t.grid[0] {
		val := "0"
		if b {
			val = "1"
		}
		bits = bits + val
	}
	return binaryAndReversed(bits)
}

func (t *tile) left() []int64 {
	bits := ""
	for _, r := range t.grid {
		val := "0"
		if r[0] {
			val = "1"
		}
		bits = bits + val
	}
	return binaryAndReversed(bits)
}

func (t *tile) right() []int64 {
	bits := ""
	for _, r := range t.grid {
		val := "0"
		if r[len(r)-1] {
			val = "1"
		}
		bits = bits + val
	}
	return binaryAndReversed(bits)
}

func (t *tile) bottom() []int64 {
	bits := ""
	for _, b := range t.grid[len(t.grid)-1] {
		val := "0"
		if b {
			val = "1"
		}
		bits = bits + val
	}
	return binaryAndReversed(bits)
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
}

func (t *tile) ccw() {
	t.cw()
	t.cw()
	t.cw()
}

func (t *tile) flipHorizontal() {
	for x := range t.grid[0] {
		lg := len(t.grid)
		for i := 0; i < (lg / 2); i++ {
			t.grid[i][x], t.grid[lg-1-i][x] = t.grid[lg-1-i][x], t.grid[i][x]
		}
	}
}

func (t *tile) flipVertical() {
	for y := range t.grid {
		ly := len(t.grid[y])
		for i := 0; i < (ly / 2); i++ {
			t.grid[y][i], t.grid[y][ly-1-i] = t.grid[y][ly-1-i], t.grid[y][i]
		}
	}
}

var (
	inputRegex = regexp.MustCompile(`^Tile (\d+):$`)
)

func parseTiles(lines []string) map[int64]*tile {
	tiles := make(map[int64]*tile)
	for i := 0; i < len(lines); i += 12 {
		tokens := inputRegex.FindStringSubmatch(lines[i])
		id, _ := strconv.ParseInt(tokens[1], 10, 64)
		tiles[id] = parseTile(id, lines[i+1:i+11])
	}
	return tiles
}

func parseTile(id int64, lines []string) *tile {
	t := tile{id: id, grid: make([][]bool, 0)}
	for i, l := range lines {
		if len(l) != 10 {
			log.Fatalf("Unexpected line: %q\n", l)
		}
		t.grid = append(t.grid, make([]bool, 0))
		for _, v := range []rune(l) {
			t.grid[i] = append(t.grid[i], v == '#')
		}
	}
	return &t
}

func findCorners(tiles map[int64]*tile) map[int64]*tile {
	dimensions := int(math.Sqrt(float64(len(tiles))))
	if dimensions*dimensions != len(tiles) {
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
	for _, t := range tiles {
		for _, s := range [][]int64{t.top(), t.left(), t.right(), t.bottom()} {
			val := aoc.Min(s[0], s[1])
			if tileEdges[val] == nil {
				tileEdges[val] = make([]*tile, 0)
			}
			tileEdges[val] = append(tileEdges[val], t)
		}
	}

	edges := make(map[int64]*tile)
	corners := make(map[int64]*tile)
	for _, te := range tileEdges {
		if len(te) == 1 {
			t := te[0]
			if _, found := edges[t.id]; found {
				corners[t.id] = t
			}
			edges[t.id] = t
		}
	}

	return corners
}

const debug = true

/*func placeTiles(tiles map[int64]*tile) [][]bool {
	dimensions := int(math.Sqrt(float64(len(tiles))))
	if dimensions*dimensions != len(tiles) {
		log.Fatalf("Not a square grid")
	}

	// Find unique patterns and mark them as the edges
	tileEdges := make(map[int64][]*tile)
	for _, t := range tiles {
		for _, s := range [][]int64{t.top(), t.left(), t.right(), t.bottom()} {
			val := aoc.Min(s[0], s[1])
			if tileEdges[val] == nil {
				tileEdges[val] = make([]*tile, 0)
			}
			tileEdges[val] = append(tileEdges[val], t)
		}
	}

	edges := make(map[int64]*tile)
	corners := make(map[int64]*tile)
	for _, te := range tileEdges {
		if len(te) == 1 {
			t := te[0]
			if _, found := edges[t.id]; found {
				corners[t.id] = t
			}
			edges[t.id] = t
		}
	}

	top, left, right, bottom := 0, 0, dimensions-1, dimensions-1
	sq := make([][]int64, dimensions)
	for i := range sq {
		sq[i] = make([]int64, dimensions)
		for j := range sq[i] {
			sq[i][j] = -1
		}
	}

	// Arbitrarily pick a first corner
	sq[top][left] =
}*/

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
					if seaMonster[y][x] != '#' {
						continue
					}
					if !picture[y+yd][x+xd] {
						// This is not a valid sea monster
						continue PICTURE
					}
				} // xd
			} // yd
			// This is a valid sea monster.  Loop back through and mark all of the spots
			for yd := 0; yd < len(seaMonster); yd++ {
				for xd := 0; xd < len(seaMonster[0]); xd++ {
					monsters[aoc.Point{X: int64(x + xd), Y: int64(y + yd)}] = true
				} // xd
			} // yd
		} // x
	} // y
	return monsters
}

const width = 4

func printTiles(tiles map[int64]*tile) {
	var ids aoc.Int64Slice
	for id := range tiles {
		ids = append(ids, id)
	}
	sort.Sort(ids)

}

func step1(tiles map[int64]*tile) int64 {
	corners := findCorners(tiles)
	product := int64(1)
	for id := range corners {
		product *= id
	}
	return product
}

func step2(tiles map[int64]*tile) int64 {
	/*origPicture := placeTiles(tiles)
	picture := tile{grid: origPicture, id: -1, used: true}
	monsters := make(map[aoc.Point]bool)
	for it := 0; true; it++ {
		monsters := findSeaMonsters(picture.grid)
		if len(monsters) > 0 {
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
			if picture.grid[y][x] && !monsters[aoc.Point{X: int64(x), Y: int64(y)}] {
				count++
			}
		}
	}*/
	printTiles(tiles, "tiles.txt")
	return -1
}
