package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"position=<     9,      1> velocity=< 0,  2>",
	"position=<     7,      0> velocity=<-1,  0>",
	"position=<     3,     -2> velocity=<-1,  1>",
	"position=<     6,     10> velocity=<-2, -1>",
	"position=<     2,     -4> velocity=< 2,  2>",
	"position=<    -6,     10> velocity=< 2, -2>",
	"position=<     1,      8> velocity=< 1, -1>",
	"position=<     1,      7> velocity=< 1,  0>",
	"position=<    -3,     11> velocity=< 1, -2>",
	"position=<     7,      6> velocity=<-1, -1>",
	"position=<    -2,      3> velocity=< 1,  0>",
	"position=<    -4,      3> velocity=< 2,  0>",
	"position=<    10,     -3> velocity=<-1,  1>",
	"position=<     5,     11> velocity=< 1, -2>",
	"position=<     4,      7> velocity=< 0, -1>",
	"position=<     8,     -2> velocity=< 0,  1>",
	"position=<    15,      0> velocity=<-2,  0>",
	"position=<     1,      6> velocity=< 1,  0>",
	"position=<     8,      9> velocity=< 0, -1>",
	"position=<     3,      3> velocity=<-1,  1>",
	"position=<     0,      5> velocity=< 0, -1>",
	"position=<    -2,      2> velocity=< 2,  0>",
	"position=<     5,     -2> velocity=< 1,  2>",
	"position=<     1,      4> velocity=< 2,  1>",
	"position=<    -2,      7> velocity=< 2, -2>",
	"position=<     3,      6> velocity=<-1, -1>",
	"position=<     5,      0> velocity=< 1,  0>",
	"position=<    -6,      0> velocity=< 2,  0>",
	"position=<     5,      9> velocity=< 1, -2>",
	"position=<    14,      7> velocity=<-2,  0>",
	"position=<    -3,      6> velocity=< 2, -1>",
}

func TestRead(t *testing.T) {
	assert.Equal(t, Projectile{x: 30444, y: 50599, dx: 3, dy: 5}, ReadProjectile("position=< 30444,  50599> velocity=< 3,  5>"))
	assert.Equal(t, Projectile{x: -30052, y: -9918, dx: 3, dy: 1}, ReadProjectile("position=<-30052,  -9918> velocity=< 3,  1>"))
	assert.Equal(t, Projectile{x: 30444, y: 50599, dx: -3, dy: -5}, ReadProjectile("position=< 30444,  50599> velocity=<-3, -5>"))
}

func TestConvergence(t *testing.T) {
	var proj []Projectile
	for _, i := range input {
		proj = append(proj, ReadProjectile(i))
	}
	assert.Equal(t, 3, TimeToConvergence(proj))
}
