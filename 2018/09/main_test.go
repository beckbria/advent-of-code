package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	players, points := ReadInput("10 players; last marble is worth 1618 points")
	assert.Equal(t, 10, players)
	assert.Equal(t, int64(1618), points)
}

func TestHighScore(t *testing.T) {
	assert.Equal(t, int64(32), HighScore(9, 25))
	assert.Equal(t, int64(8317), HighScore(10, 1618))
	assert.Equal(t, int64(146373), HighScore(13, 7999))
	assert.Equal(t, int64(2764), HighScore(17, 1104))
	assert.Equal(t, int64(54718), HighScore(21, 6111))
	assert.Equal(t, int64(37305), HighScore(30, 5807))
}
