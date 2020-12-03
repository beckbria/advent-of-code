package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trees = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func Test(t *testing.T) {
	assert.Equal(t, findTrees(trees, 1, 1), 2)
	assert.Equal(t, findTrees(trees, 3, 1), 7)
	assert.Equal(t, findTrees(trees, 5, 1), 3)
	assert.Equal(t, findTrees(trees, 7, 1), 4)
	assert.Equal(t, findTrees(trees, 1, 2), 2)
}
