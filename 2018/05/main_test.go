package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, "dabCBAcaDA", Reduce("dabAcCaCBAcCcaDA"))
}

func testShortest(t *testing.T) {
	assert.Equal(t, "daDA", ShortestRemoveOne("dabAcCaCBAcCcaDA"))
}
