package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollision(t *testing.T) {
	rails, carts := ReadTrackNetwork([]string{
		"/->-\\        ",
		"|   |  /----\\",
		"| /-+--+-\\  |",
		"| | |  | v  |",
		"\\-+-/  \\-+--/",
		"  \\------/   ",
	})
	x, y, time := FirstCollision(&rails, &carts)
	assert.Equal(t, 7, x)
	assert.Equal(t, 3, y)
	assert.Equal(t, 14, time)
}

func TestLastCartStanding(t *testing.T) {
	rails, carts := ReadTrackNetwork([]string{
		"/>-<\\  ",
		"|   |  ",
		"| /<+-\\",
		"| | | v",
		"\\>+</ |",
		"  |   ^",
		"  \\<->/"})

	x, y, time := LastCartStanding(&rails, &carts)
	assert.Equal(t, 6, x)
	assert.Equal(t, 4, y)
	assert.Equal(t, 3, time)
}

//   0000000000111
//   0123456789012
// 0 /->-\
// 1 |   |  /----\
// 2 | /-+--+-\  |
// 3 | | |  | v  |
// 4 \-+-/  \-+--/
// 5   \------/
