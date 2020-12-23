package main

import (
	"reflect"
	"testing"

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

	tile := parseTile(1, patterns[0])
	tile.cw()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[90]).grid))
	tile.cw()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[180]).grid))
	tile.cw()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[270]).grid))
	tile.cw()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[0]).grid))
	tile.ccw()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[270]).grid))
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

	tile := parseTile(1, patterns[0])
	tile.flipHorizontal()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[1]).grid))
	tile.flipHorizontal()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[0]).grid))
	tile.flipVertical()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[2]).grid))
	tile.flipVertical()
	assert.True(t, reflect.DeepEqual(tile.grid, parseTile(1, patterns[0]).grid))
}

func TestStep1(t *testing.T) {
	tiles := parseTiles(input)
	assert.Equal(t, int64(20899048083289), step1(tiles))
}

func TestStep2(t *testing.T) {
	tiles := parseTiles(input)
	assert.Equal(t, int64(-1), step2(tiles))
}
