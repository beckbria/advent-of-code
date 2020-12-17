package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(112), step1([]string{".#.", "..#", "###"}))
}

func TestAdvance(t *testing.T) {
	assert.Equal(t, int64(848), step2([]string{".#.", "..#", "###"}))
}
