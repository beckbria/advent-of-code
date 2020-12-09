package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	assert.Equal(t, step1([]string{}), -1)
}

func TestStep2(t *testing.T) {
	assert.Equal(t, step2([]string{}), -1)
}
