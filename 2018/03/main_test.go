package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t, Fabric{ID: 1, Left: 808, Top: 550, Right: 819, Bottom: 571}, ReadFabric("#1 @ 808,550: 12x22"))
}

func TestOverlapping(t *testing.T) {
	input := []Fabric{ReadFabric("#1 @ 1,3: 4x4"), ReadFabric("#2 @ 3,1: 4x4"), ReadFabric("#3 @ 5,5: 2x2")}
	assert.Equal(t, 4, OverlappingArea(input))
}

func TestDistinct(t *testing.T) {
	input := []Fabric{ReadFabric("#1 @ 1,3: 4x4"), ReadFabric("#2 @ 3,1: 4x4"), ReadFabric("#3 @ 5,5: 2x2")}
	distinct, err := DistinctFabric(input)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(3), distinct)
}
