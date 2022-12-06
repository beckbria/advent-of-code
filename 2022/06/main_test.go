package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestStep1(t *testing.T) {
	expectedAnswers := []int64{7,5,6,10,11}
	for i := range testInput {
		assert.Equal(t, expectedAnswers[i], step1(testInput[i]))
	}
}

func TestStep2(t *testing.T) {
	expectedAnswers := []int64{19,23,23,29,26}
	for i := range testInput {
		assert.Equal(t, expectedAnswers[i], step2(testInput[i]))
	}
}
