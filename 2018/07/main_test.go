package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	A = 'A'
	B = 'B'
	C = 'C'
	D = 'D'
	E = 'E'
	F = 'F'
)

var (
	allNodes = []byte{A, B, C, D, E, F}

	input = []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin."}
)

func TestRead(t *testing.T) {

	dependsOn, isDependedOnBy := ReadDependencies(input)
	expectedDependsOn := make(CharacterMap)
	expectedIsDependedOnBy := make(CharacterMap)

	for _, n := range allNodes {
		expectedDependsOn[n] = make(map[byte]bool)
		expectedIsDependedOnBy[n] = make(map[byte]bool)
	}

	expectedDependsOn[A][C] = true
	expectedDependsOn[B][A] = true
	expectedDependsOn[D][A] = true
	expectedDependsOn[E][B] = true
	expectedDependsOn[E][D] = true
	expectedDependsOn[E][F] = true
	expectedDependsOn[F][C] = true
	assert.Equal(t, expectedDependsOn, dependsOn)

	expectedIsDependedOnBy[A][B] = true
	expectedIsDependedOnBy[A][D] = true
	expectedIsDependedOnBy[B][E] = true
	expectedIsDependedOnBy[C][A] = true
	expectedIsDependedOnBy[C][F] = true
	expectedIsDependedOnBy[D][E] = true
	expectedIsDependedOnBy[F][E] = true
	assert.Equal(t, expectedIsDependedOnBy, isDependedOnBy)
}

func TestExecutionOrder(t *testing.T) {
	dependsOn, isDependedOnBy := ReadDependencies(input)
	assert.Equal(t, "CABDFE", ExecutionOrder(dependsOn, isDependedOnBy))
}

func TestNodeTime(t *testing.T) {
	assert.Equal(t, 1, NodeTime('A', 0))
	assert.Equal(t, 13, NodeTime('M', 0))
	assert.Equal(t, 26, NodeTime('Z', 0))
	assert.Equal(t, 61, NodeTime('A', 60))
	assert.Equal(t, 73, NodeTime('M', 60))
	assert.Equal(t, 86, NodeTime('Z', 60))
}

func TestExecutionTime(t *testing.T) {
	dependsOn, isDependedOnBy := ReadDependencies(input)
	assert.Equal(t, 15, ExecutionTime(dependsOn, isDependedOnBy, 2, 0))
}
