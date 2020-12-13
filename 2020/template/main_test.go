package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, -1, step1([]string{}))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, -1, step2([]string{}))
}
