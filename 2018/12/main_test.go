package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	state := ReadInitialState("initial state: #..#.#..##......###...###")
	rules := MakeRuleSet(ReadRules([]string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}))
	afterTwenty := Advance(&state, rules, 20)
	assert.Equal(t, int64(325), Score(&afterTwenty))
}
