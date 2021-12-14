package main

import (
	"testing"

	"github.com/beckbria/advent-of-code/2019/lib"

	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/10
func TestMonitoringStation1(t *testing.T) {
	input := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##"}
	loc, count, allCounts := bestMonitoringStation(newMap(input))
	expectedCounts := make(map[lib.Point]int)
	expectedCounts[lib.Point{X: 1, Y: 0}] = 7
	expectedCounts[lib.Point{X: 4, Y: 0}] = 7
	expectedCounts[lib.Point{X: 0, Y: 2}] = 6
	expectedCounts[lib.Point{X: 1, Y: 2}] = 7
	expectedCounts[lib.Point{X: 2, Y: 2}] = 7
	expectedCounts[lib.Point{X: 3, Y: 2}] = 7
	expectedCounts[lib.Point{X: 4, Y: 2}] = 5
	expectedCounts[lib.Point{X: 4, Y: 3}] = 7
	expectedCounts[lib.Point{X: 3, Y: 4}] = 8
	expectedCounts[lib.Point{X: 4, Y: 4}] = 7

	assert.Equal(t, expectedCounts, allCounts)
	assert.Equal(t, lib.Point{X: 3, Y: 4}, loc)
	assert.Equal(t, 8, count)
}

func TestMonitoringStation2(t *testing.T) {
	input := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####"}
	loc, count, _ := bestMonitoringStation(newMap(input))
	assert.Equal(t, lib.Point{X: 5, Y: 8}, loc)
	assert.Equal(t, 33, count)
}

func TestMonitoringStation3(t *testing.T) {
	input := []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###."}
	loc, count, _ := bestMonitoringStation(newMap(input))
	assert.Equal(t, lib.Point{X: 1, Y: 2}, loc)
	assert.Equal(t, 35, count)
}

func TestMonitoringStation4(t *testing.T) {
	input := []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#.."}
	loc, count, _ := bestMonitoringStation(newMap(input))
	assert.Equal(t, lib.Point{X: 6, Y: 3}, loc)
	assert.Equal(t, 41, count)
}

func TestMonitoringStation5(t *testing.T) {
	input := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##"}
	loc, count, _ := bestMonitoringStation(newMap(input))
	assert.Equal(t, lib.Point{X: 11, Y: 13}, loc)
	assert.Equal(t, 210, count)
}

func TestDestructionOrder(t *testing.T) {
	input := []string{
		".#....#####...#..",
		"##...##.#####..##",
		"##...#...#.#####.",
		"..#.....#...###..",
		"..#.#.....#....##"}
	from := lib.Point{X: 8, Y: 3}
	order := destructionOrder(newMap(input), &from)
	assert.Equal(t, 1, order[lib.Point{X: 8, Y: 1}])
	assert.Equal(t, 2, order[lib.Point{X: 9, Y: 0}])
	assert.Equal(t, 3, order[lib.Point{X: 9, Y: 1}])
	assert.Equal(t, 4, order[lib.Point{X: 10, Y: 0}])
	assert.Equal(t, 5, order[lib.Point{X: 9, Y: 2}])
	assert.Equal(t, 6, order[lib.Point{X: 11, Y: 1}])
	assert.Equal(t, 7, order[lib.Point{X: 12, Y: 1}])
	assert.Equal(t, 8, order[lib.Point{X: 11, Y: 2}])
	assert.Equal(t, 9, order[lib.Point{X: 15, Y: 1}])
}
