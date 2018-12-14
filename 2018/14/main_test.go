package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastTenRecipes(t *testing.T) {
	assert.Equal(t, []int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}, NextTenRecipes(9))
	assert.Equal(t, []int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}, NextTenRecipes(18))
	assert.Equal(t, []int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}, NextTenRecipes(2018))
}

func TestScoreIndex(t *testing.T) {
	assert.Equal(t, 9, ScoreIndex([]int{5, 1, 5, 8, 9}))
	assert.Equal(t, 5, ScoreIndex([]int{0, 1, 2, 4, 5}))
	assert.Equal(t, 18, ScoreIndex([]int{9, 2, 5, 1, 0}))
	assert.Equal(t, 2018, ScoreIndex([]int{5, 9, 4, 1, 4}))
}
