package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodCount(t *testing.T) {

	//    44444455555555
	//    9999990000000
	//    45678901234567
	//
	// 01 ......1.....#.
	// 02 .#..#g2hi...#.
	// 03 .#..#f3#j.....
	// 04 .#..#e4#k.....
	// 05 .#dcba5#l.....
	// 06 .#09876#m.....
	// 07 .#######n.....
	// 08 ........o.....
	// 09 ...a9876p|||..
	// 10 ...b#432q5#|..
	// 11 ...c#zyxr1#|..
	// 12 ...d#vutsw#|..
	// 13 ...e#######|..

	running, standing := FloodCount([]string{
		"x=495, y=2..7",
		"y=7, x=495..501",
		"x=501, y=3..7",
		"x=498, y=2..4",
		"x=506, y=1..2",
		"x=498, y=10..13",
		"x=504, y=10..13",
		"y=13, x=498..504"})
	assert.Equal(t, 57, running+standing)
	assert.Equal(t, 29, standing)
}
