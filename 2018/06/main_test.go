package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPoint(t *testing.T) {
	assert.Equal(t, Point{X: 158, Y: 163}, ReadPoint("158, 163"))
	assert.Equal(t, []Point{{X: 158, Y: 163}, {X: 287, Y: 68}}, ReadPoints([]string{"158, 163", "287, 68"}))
}

func TestMaxArea(t *testing.T) {
	pts := ReadPoints([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	assert.Equal(t, int64(17), MaximumArea(pts))
}

func TestMaxSafeArea(t *testing.T) {
	pts := ReadPoints([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	assert.Equal(t, int64(16), MaximumSafeArea(pts, 32))
}
