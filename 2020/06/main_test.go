package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func TestStep1(t *testing.T) {
	assert.Equal(t, step1(input), 11)
}

func TestStep2(t *testing.T) {
	assert.Equal(t, step2(input), 6)
}
