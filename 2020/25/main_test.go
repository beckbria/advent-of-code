package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	cardLoop, doorLoop, encryptionKey := step1([]int64{5764801, 17807724})
	assert.Equal(t, int64(8), cardLoop)
	assert.Equal(t, int64(11), doorLoop)
	assert.Equal(t, int64(14897079), encryptionKey)
}
