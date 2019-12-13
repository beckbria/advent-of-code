package main

import (
	"testing"

	"../aoc"
	"github.com/stretchr/testify/assert"
)

// Examples of program input and output from https://adventofcode.com/2019/day/12
func TestGravity(t *testing.T) {
	ms := readMoons([]string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>"})
	ms.step()

	assert.Equal(t, ms[0].v, aoc.Point3{X: 3, Y: -1, Z: -1})
	assert.Equal(t, ms[1].v, aoc.Point3{X: 1, Y: 3, Z: 3})
	assert.Equal(t, ms[2].v, aoc.Point3{X: -3, Y: 1, Z: -3})
	assert.Equal(t, ms[3].v, aoc.Point3{X: -1, Y: -3, Z: 1})
}

func TestStep(t *testing.T) {
	ms := readMoons([]string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>"})
	for i := 0; i < 10; i++ {
		ms.step()
	}

	assert.Equal(t, ms[0].v, aoc.Point3{X: -3, Y: -2, Z: 1})
	assert.Equal(t, ms[1].v, aoc.Point3{X: -1, Y: 1, Z: 3})
	assert.Equal(t, ms[2].v, aoc.Point3{X: 3, Y: 2, Z: -3})
	assert.Equal(t, ms[3].v, aoc.Point3{X: 1, Y: -1, Z: -1})

	assert.Equal(t, ms[0].p, aoc.Point3{X: 2, Y: 1, Z: -3})
	assert.Equal(t, ms[1].p, aoc.Point3{X: 1, Y: -8, Z: 0})
	assert.Equal(t, ms[2].p, aoc.Point3{X: 3, Y: -6, Z: 1})
	assert.Equal(t, ms[3].p, aoc.Point3{X: 2, Y: 0, Z: 4})
}

func TestEnergy(t *testing.T) {
	ms := readMoons([]string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>"})
	for i := 0; i < 10; i++ {
		ms.step()
	}

	assert.Equal(t, int64(179), ms.energy())
}

func TestFirstCycle(t *testing.T) {
	ms := readMoons([]string{
		"<x=-8, y=-10, z=0>",
		"<x=5, y=5, z=10>",
		"<x=2, y=-7, z=3>",
		"<x=9, y=-8, z=-3>"})
	assert.Equal(t, int64(4686774924), firstCycle(ms))
}
