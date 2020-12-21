package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = parseRecipes([]string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
})

func TestStep1(t *testing.T) {
	assert.Equal(t, int64(5), step1(input))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, int64(-1), step2(input))
}
