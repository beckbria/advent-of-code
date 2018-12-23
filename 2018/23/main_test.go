package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert.Equal(t, nanobot{x: 0, y: 0, z: 0, r: 4}, readNanobot("pos=<0,0,0>, r=4"))
	assert.Equal(t,
		fleet{
			nanobot{x: 76659180, y: 55463797, z: 20890147, r: 80344142},
			nanobot{x: -2084092, y: 73605216, z: 31684616, r: 79399057}},
		readFleet([]string{
			"pos=<76659180,55463797,20890147>, r=80344142",
			"pos=<-2084092,73605216,31684616>, r=79399057"}))
}

func TestStrongestRange(t *testing.T) {
	f := readFleet([]string{
		"pos=<0,0,0>, r=4",
		"pos=<1,0,0>, r=1",
		"pos=<4,0,0>, r=3",
		"pos=<0,2,0>, r=1",
		"pos=<0,5,0>, r=3",
		"pos=<0,0,3>, r=1",
		"pos=<1,1,1>, r=1",
		"pos=<1,1,2>, r=1",
		"pos=<1,3,1>, r=1",
	})
	assert.Equal(t, 7, f.inRangeOfStrongest())
}
