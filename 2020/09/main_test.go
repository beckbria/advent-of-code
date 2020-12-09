package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int64{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestStep1(t *testing.T) {
	assert.Equal(t, step1(input, 5), int64(127))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, step2(input, int64(127)), int64(62))
}
