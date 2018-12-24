package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t,
		group{
			id:         1,
			count:      956,
			hp:         7120,
			weakTo:     map[string]bool{"bludgeoning": true, "slashing": true},
			immuneTo:   map[string]bool{},
			damage:     71,
			attack:     "radiation",
			initiative: 7},
		readGroup("956 units each with 7120 hit points (weak to bludgeoning, slashing) with an attack that does 71 radiation damage at initiative 7"))

	assert.Equal(t,
		group{
			id:         2,
			count:      1155,
			hp:         5643,
			weakTo:     map[string]bool{"bludgeoning": true},
			immuneTo:   map[string]bool{"cold": true},
			damage:     42,
			attack:     "slashing",
			initiative: 15},
		readGroup("1155 units each with 5643 hit points (weak to bludgeoning; immune to cold) with an attack that does 42 slashing damage at initiative 15"))

	assert.Equal(t,
		group{
			id:         3,
			count:      1062,
			hp:         11023,
			weakTo:     map[string]bool{"bludgeoning": true},
			immuneTo:   map[string]bool{"cold": true, "radiation": true},
			damage:     93,
			attack:     "fire",
			initiative: 19},
		readGroup("1062 units each with 11023 hit points (immune to cold, radiation; weak to bludgeoning) with an attack that does 93 fire damage at initiative 19"))

	assert.Equal(t,
		group{
			id:         4,
			count:      5009,
			hp:         8078,
			weakTo:     map[string]bool{},
			immuneTo:   map[string]bool{},
			damage:     14,
			attack:     "slashing",
			initiative: 8},
		readGroup("5009 units each with 8078 hit points with an attack that does 14 slashing damage at initiative 8"))
}

func TestWinningArmyCount(t *testing.T) {
	b := readBody([]string{
		"Immune System:",
		"17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2",
		"989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3",
		"",
		"Infection:",
		"801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1",
		"4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4"})

	assert.Equal(t, 5216, b.winningArmyCount())
}
