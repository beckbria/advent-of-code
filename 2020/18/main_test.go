package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(26), tokenize("2 * 3 + (4 * 5)").evaluate(false))
	assert.Equal(t, int64(437), tokenize("5 + (8 * 3 + 9 + 3 * 4 * 3)").evaluate(false))
	assert.Equal(t, int64(13632), tokenize("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2").evaluate(false))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(46), tokenize("2 * 3 + (4 * 5)").evaluate(true))
	assert.Equal(t, int64(1445), tokenize("5 + (8 * 3 + 9 + 3 * 4 * 3)").evaluate(true))
	assert.Equal(t, int64(23340), tokenize("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2").evaluate(true))
}
