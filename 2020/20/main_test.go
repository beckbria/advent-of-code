package main

import (
	"reflect"
	"testing"

	"../../aoc"
	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Tile 2311:",
	"..##.#..#.",
	"##..#.....",
	"#...##..#.",
	"####.#...#",
	"##.##.###.",
	"##...#.###",
	".#.#.#..##",
	"..#....#..",
	"###...#.#.",
	"..###..###",
	"",
	"Tile 1951:",
	"#.##...##.",
	"#.####...#",
	".....#..##",
	"#...######",
	".##.#....#",
	".###.#####",
	"###.##.##.",
	".###....#.",
	"..#.#..#.#",
	"#...##.#..",
	"",
	"Tile 1171:",
	"####...##.",
	"#..##.#..#",
	"##.#..#.#.",
	".###.####.",
	"..###.####",
	".##....##.",
	".#...####.",
	"#.##.####.",
	"####..#...",
	".....##...",
	"",
	"Tile 1427:",
	"###.##.#..",
	".#..#.##..",
	".#.##.#..#",
	"#.#.#.##.#",
	"....#...##",
	"...##..##.",
	"...#.#####",
	".#.####.#.",
	"..#..###.#",
	"..##.#..#.",
	"",
	"Tile 1489:",
	"##.#.#....",
	"..##...#..",
	".##..##...",
	"..#...#...",
	"#####...#.",
	"#..#.#.#.#",
	"...#.#.#..",
	"##.#...##.",
	"..##.##.##",
	"###.##.#..",
	"",
	"Tile 2473:",
	"#....####.",
	"#..#.##...",
	"#.##..#...",
	"######.#.#",
	".#...#.#.#",
	".#########",
	".###.#..#.",
	"########.#",
	"##...##.#.",
	"..###.#.#.",
	"",
	"Tile 2971:",
	"..#.#....#",
	"#...###...",
	"#.#.###...",
	"##.##..#..",
	".#####..##",
	".#..####.#",
	"#..#.#..#.",
	"..####.###",
	"..#.#.###.",
	"...#.#.#.#",
	"",
	"Tile 2729:",
	"...#.#.#.#",
	"####.#....",
	"..#.#.....",
	"....#..#.#",
	".##..##.#.",
	".#.####...",
	"####.#.#..",
	"##.####...",
	"##..#.##..",
	"#.##...##.",
	"",
	"Tile 3079:",
	"#.#.#####.",
	".#..######",
	"..#.......",
	"######....",
	"####.#..#.",
	".#...#.##.",
	"#.#####.##",
	"..#.###...",
	"..#.......",
	"..#.###...",
}

func TestRotate(t *testing.T) {
	patterns := make(map[int][]string)
	patterns[0] = []string{
		"##.",
		"#..",
		"..."}
	patterns[90] = []string{
		".##",
		"..#",
		"..."}
	patterns[180] = []string{
		"...",
		"..#",
		".##"}
	patterns[270] = []string{
		"...",
		"#..",
		"##."}

	tl := newTile(1, patterns[0])
	tl.cw()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[90]).grid))
	tl.cw()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[180]).grid))
	tl.cw()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[270]).grid))
	tl.cw()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[0]).grid))
	tl.ccw()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[270]).grid))
}

func TestFlip(t *testing.T) {
	patterns := make(map[int][]string)
	patterns[0] = []string{
		"##.",
		"..#",
		".##"}
	patterns[1] = []string{
		".##",
		"#..",
		"##."}
	patterns[2] = []string{
		".##",
		"..#",
		"##."}

	tl := newTile(1, patterns[0])
	tl.flipHorizontal()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[1]).grid))
	tl.flipHorizontal()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[0]).grid))
	tl.flipVertical()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[2]).grid))
	tl.flipVertical()
	assert.True(t, reflect.DeepEqual(tl.grid, newTile(1, patterns[0]).grid))
}

func TestFindEdges(t *testing.T) {
	ts := parseTiles(input)
	edges, corners := ts.findEdges()
	var c []int64
	for id := range corners {
		c = append(c, id)
	}

	expectedCorners := []int64{1171, 1951, 2971, 3079}
	assert.ElementsMatch(t, expectedCorners, c)

	expectedEdges := []int64{2311, 2729, 2473, 1489}
	var e []int64
	for id := range edges {
		e = append(e, id)
	}
	assert.ElementsMatch(t, expectedEdges, e)
}

func TestFindSeaMonsters(t *testing.T) {
	tile := newTile(1, []string{
		".####...#####..#...###..",
		"#####..#..#.#.####..#.#.",
		".#.#...#.###...#.##.##..",
		"#.#.##.###.#.##.##.#####",
		"..##.###.####..#.####.##",
		"...#.#..##.##...#..#..##",
		"#.##.#..#.#..#..##.#.#..",
		".###.##.....#...###.#...",
		"#.####.#.#....##.#..#.#.",
		"##...#..#....#..#...####",
		"..#.##...###..#.#####..#",
		"....#.##.#.#####....#...",
		"..##.##.###.....#.##..#.",
		"#...#...###..####....##.",
		".#.##...#.##.#.#.###...#",
		"#.###.#..####...##..#...",
		"#.###...#.##...#.######.",
		".###.###.#######..#####.",
		"..##.#..#..#.#######.###",
		"#.#..##.########..#..##.",
		"#.#####..#.#...##..#....",
		"#....##..#.#########..##",
		"#...#.....#..##...###.##",
		"#..###....##.#...##.##.#",
	})

	monsters := findSeaMonsters(tile.grid)
	expected := make(map[aoc.Point]bool)
	expectMonster(expected, aoc.Point{X: 2, Y: 2})
	expectMonster(expected, aoc.Point{X: 1, Y: 16})
	assert.True(t, reflect.DeepEqual(monsters, expected))
}

func expectMonster(expected map[aoc.Point]bool, origin aoc.Point) {
	for y, line := range seaMonster {
		for x, c := range line {
			if c == '#' {
				expected[aoc.Point{X: origin.X + int64(x), Y: origin.Y + int64(y)}] = true
			}
		}
	}
}

func TestStep1(t *testing.T) {
	ts := parseTiles(input)
	assert.Equal(t, int64(20899048083289), step1(ts))
}

func TestStep2(t *testing.T) {
	ts := parseTiles(input)
	assert.Equal(t, int64(273), step2(ts))
}
