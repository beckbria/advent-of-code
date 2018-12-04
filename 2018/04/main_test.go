package main

import (
	"bufio"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateTestData(t *testing.T) []Event {
	input, err := ParseEvents(bufio.NewScanner(strings.NewReader(
		`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`)))
	assert.Equal(t, nil, err)
	return input
}

func TestParseEvent(t *testing.T) {
	event, err := ParseEvent("[1518-09-01 00:58] wakes up")
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(0), event.guard)
	assert.Equal(t, WakeUp, event.action)
	assert.Equal(t, 1518, event.timestamp.Year())
	assert.Equal(t, time.September, event.timestamp.Month())
	assert.Equal(t, 1, event.timestamp.Day())
	assert.Equal(t, 0, event.timestamp.Hour())
	assert.Equal(t, 58, event.timestamp.Minute())
}

func TestCountSleep(t *testing.T) {
	input := generateTestData(t)
	sleep := CountSleep(input)
	// Count the minutes slept by guard #10
	count := int64(0)
	for _, t := range sleep[int64(10)] {
		count += t
	}
	assert.Equal(t, int64(50), count)
	count = int64(0)
	for _, t := range sleep[int64(99)] {
		count += t
	}
	assert.Equal(t, int64(30), count)
}

func TestMostAsleep(t *testing.T) {
	input := generateTestData(t)
	sleep := CountSleep(input)
	assert.Equal(t, int64(10), MostAsleep(sleep))
}

func TestMostAsleepMinute(t *testing.T) {
	input := generateTestData(t)
	sleep := CountSleep(input)
	guard := MostAsleep(sleep)
	minute := MostCommonAsleepMinute(sleep, guard)
	assert.Equal(t, 24, minute)
}

func TestMostReliablyAsleep(t *testing.T) {
	input := generateTestData(t)
	sleep := CountSleep(input)
	guard, minute := MostReliablyAsleep(sleep)
	assert.Equal(t, guard, int64(99))
	assert.Equal(t, minute, 45)
}
