package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindConstellations(t *testing.T) {
	pts := readVectors([]string{
		"-1,2,2,0",
		"0,0,2,-2",
		"0,0,0,-2",
		"-1,2,0,0",
		"-2,-2,-2,2",
		"3,0,2,-1",
		"-1,3,2,2",
		"-1,0,-1,0",
		"0,2,1,-2",
		"3,0,0,0",
	})
	assert.Equal(t, 4, findConstellations(pts))
}
