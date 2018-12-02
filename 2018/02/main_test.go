package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	assert.Equal(t, 12, Checksum(input))
}

func TestEditDistanceOnce(t *testing.T) {
	input := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	a, b, err := EditDistanceOne(input)
	assert.Equal(t, nil, err)
	assert.Equal(t, "fghij", a)
	assert.Equal(t, "fguij", b)
}

func TestCommonLetters(t *testing.T) {
	c, err := CommonLetters("fghij", "fguij")
	assert.Equal(t, nil, err)
	assert.Equal(t, "fgij", c)
}
