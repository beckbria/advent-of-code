package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"sesenwnenenewseeswwswswwnenewsewsw",
	"neeenesenwnwwswnenewnwwsewnenwseswesw",
	"seswneswswsenwwnwse",
	"nwnwneseeswswnenewneswwnewseswneseene",
	"swweswneswnenwsewnwneneseenw",
	"eesenwseswswnenwswnwnwsewwnwsene",
	"sewnenenenesenwsewnenwwwse",
	"wenwwweseeeweswwwnwwe",
	"wsweesenenewnwwnwsenewsenwwsesesenwne",
	"neeswseenwwswnwswswnw",
	"nenwswwsewswnenenewsenwsenwnesesenew",
	"enewnwewneswsewnwswenweswnenwsenwsw",
	"sweneswneswneneenwnewenewwneswswnese",
	"swwesenesewenwneswnwwneseswwne",
	"enesenwswwswneneswsenwnewswseenwsese",
	"wnwnesenesenenwwnenwsewesewsesesew",
	"nenewswnwewswnenesenwnesewesw",
	"eneswnwswnwsenenwnwnwwseeswneewsenese",
	"neswnwewnwnwseenwseesewsenwsweewe",
	"wseweeenwnesenwwwswnew",
}

func TestPathFind(t *testing.T) {
	assert.Equal(t, hex{X: 0, Y: -1, Z: 1}, followPath(parseInst("esew")))
}

func TestStep1(t *testing.T) {
	inst := parseInstructions(input)
	answer, _ := step1(inst)
	assert.Equal(t, int64(10), answer)
}

func TestStep2(t *testing.T) {
	_, grid := step1(parseInstructions(input))
	assert.Equal(t, int64(2208), step2(grid))
}
