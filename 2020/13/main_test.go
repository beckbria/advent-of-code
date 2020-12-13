package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const start = 939

var departures = []int64{7, 13, -1, -1, 59, -1, 31, 19}

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(295), step1(start, departures))
}

func TestStep2Naive(t *testing.T) {
	assert.Equal(t, int64(1068781), step2Naive(departures, 0))
	assert.Equal(t, int64(1202161486), step2Naive([]int64{1789, 37, 47, 1889}, 0))
}
