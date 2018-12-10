package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t, Projectile{x: 30444, y: 50599, dx: 3, dy: 5}, ReadProjectile("position=< 30444,  50599> velocity=< 3,  5>"))
	assert.Equal(t, Projectile{x: -30052, y: -9918, dx: 3, dy: 1}, ReadProjectile("position=<-30052,  -9918> velocity=< 3,  1>"))
	assert.Equal(t, Projectile{x: 30444, y: 50599, dx: -3, dy: -5}, ReadProjectile("position=< 30444,  50599> velocity=<-3, -5>"))
}
