package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngredients(t *testing.T) {
	input := parseRecipes([]string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	})
	count, cause := step1(input)
	assert.Equal(t, int64(5), count)
	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", step2(cause))
}
