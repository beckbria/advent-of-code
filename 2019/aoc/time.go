package aoc

import (
	"time"
)

// stopwatch is a basic timer which records time since its construction or most recent reset.
type stopwatch struct {
	start time.Time
}

// NewStopwatch creates a stopwatch class which records the time since its construction or most recent reset.
func NewStopwatch() stopwatch {
	sw := stopwatch { time.Now() }
	return sw
}

// Reset resets the start time for the stopwatch
func (sw *stopwatch) Reset() {
	sw.start = time.Now()
}

// Elapsed returns how long the stopwatch has been running
func (sw *stopwatch) Elapsed() time.Duration {
	return time.Since(sw.start)
}