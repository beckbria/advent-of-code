package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(7), step1([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(5), step2([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}))
}
