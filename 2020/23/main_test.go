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
	assert.Equal(t, -1, step2(389125467))
}
