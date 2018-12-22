package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiskToTarget(t *testing.T) {
	c := makeCave(510, point{x: 10, y: 10})
	assert.Equal(t, 114, c.riskToTarget())
}

func TestTimeToTarget(t *testing.T) {
	c := makeCave(510, point{x: 10, y: 10})
	assert.Equal(t, 45, c.timeToTarget())
}
