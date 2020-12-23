package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, 92658374, step1(389125467, 10))
	assert.Equal(t, 67384529, step1(389125467, 100))
}

func TestStep2(t *testing.T) {
	first, second, product := step2(389125467)
	assert.Equal(t, int64(934001), first)
	assert.Equal(t, int64(159792), second)
	assert.Equal(t, int64(149245887792), product)
}
