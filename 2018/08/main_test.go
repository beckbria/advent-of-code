package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func TestStringToInts(t *testing.T) {
	assert.Equal(t, []int64{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, StringToInts(data))
}

func TestMetadataSum(t *testing.T) {
	tree := ReadNodes(StringToInts(data))
	assert.Equal(t, int64(138), MetadataSum(tree))
}

func TestValue(t *testing.T) {
	tree := ReadNodes(StringToInts(data))
	assert.Equal(t, int64(66), Value(tree))
}
