package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, part1([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}), int64(2))
}

func Test2(t *testing.T) {
	assert.Equal(t, part2([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}), int64(1))
}
