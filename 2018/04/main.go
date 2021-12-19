package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Action represents something a guard does
type Action int

const (
	// ShiftStart represents Starting a Shift
	ShiftStart Action = 1
	// WakeUp represents exactly what you think it does
	WakeUp Action = 2
	// Sleep represents how unhappy I am with Go's mandatory comment rules
	Sleep Action = 3
)

// Event represents an occurrence during the night
type Event struct {
	timestamp time.Time
	action    Action
	guard     int64
}

// ParseAction parses an action from an input line
func ParseAction(line string) (Action, int64, error) {
	if line == "wakes up" {
		return WakeUp, 0, nil
	} else if line == "falls asleep" {
		return Sleep, 0, nil
	} else if line[0:5] != "Guard" {
		return 0, 0, fmt.Errorf("Unexpected Action: %s -- \"%s\"", line, line[0:6])
	} else {
		// Guard #2917 begins shift
		// TODO: Use Regex to parse this stuff
		words := strings.Split(line, " ")
		guard, err := strconv.ParseInt(words[1][1:], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return ShiftStart, guard, nil
	}
}

// ParseEvent parses an event from an input line
func ParseEvent(line string) (Event, error) {
	// Format: [1518-05-21 23:59] Guard #2917 begins shift
	var evt Event
	t, err := time.Parse("2006-01-02 15:04", line[1:17])
	if err != nil {
		return evt, err
	}

	act, guard, err := ParseAction(line[19:])
	if err != nil {
		return evt, err
	}
	evt.timestamp = t
	evt.action = act
	evt.guard = guard
	return evt, nil
}

// ParseEvents parses more than one event
func ParseEvents(scanner *bufio.Scanner) ([]Event, error) {
	// Read the events
	var input []Event
	for scanner.Scan() {
		evt, err := ParseEvent(scanner.Text())
		if err != nil {
			return []Event{}, err
		}
		input = append(input, evt)
	}
	err := scanner.Err()
	if err != nil {
		return []Event{}, err
	}

	// Sort them chronologically
	sort.Slice(input, func(a, b int) bool {
		return input[a].timestamp.Before(input[b].timestamp)
	})

	return input, nil
}

// SleepMap maps from the guard ID to a map of the minute (12:xx) to the number of days the guard was seen asleep at that minute
type SleepMap map[int64]map[int]int64

// CountSleep counts the total number of minutes each guard is asleep.  Returns a
func CountSleep(input []Event) SleepMap {
	sleepCount := make(SleepMap)
	currentGuard := int64(-1)
	startedSleeping := time.Now()
	for _, evt := range input {
		switch evt.action {
		case ShiftStart:
			currentGuard = evt.guard
		case Sleep:
			startedSleeping = evt.timestamp
		case WakeUp:
			for i := startedSleeping.Minute(); i < evt.timestamp.Minute(); i++ {
				if sleepCount[currentGuard] == nil {
					sleepCount[currentGuard] = make(map[int]int64)
				}
				sleepCount[currentGuard][i]++
			}
		}
	}
	return sleepCount
}

// MostAsleep finds the guard who was asleep for the longest time
func MostAsleep(sleep SleepMap) int64 {
	guard := int64(-1)
	mostSleep := int64(-1000000000) // TODO: Does golang REALLY have no concept of INT_MIN?
	for g, s := range sleep {
		sleepCount := int64(0)
		for _, z := range s {
			sleepCount += z
		}
		if sleepCount > mostSleep {
			mostSleep = sleepCount
			guard = g
		}
	}
	return guard
}

// MostCommonAsleepMinute finds the minute where a guard was most often asleep
func MostCommonAsleepMinute(sleep SleepMap, guard int64) int {
	zzzMinute := -1
	zzzCount := int64(-1000000000)
	for minute, count := range sleep[guard] {
		if count > zzzCount {
			zzzCount = count
			zzzMinute = minute
		}
	}
	return zzzMinute
}

// MostReliablyAsleep returns guard, minute
func MostReliablyAsleep(sleep SleepMap) (int64, int) {
	zzzCount := int64(-1000000000)
	zzzMinute := -1
	zzzGuard := int64(-1)
	for guard, s := range sleep {
		for minute, count := range s {
			if count > zzzCount {
				zzzGuard = guard
				zzzMinute = minute
				zzzCount = count
			}
		}
	}
	return zzzGuard, zzzMinute
}

func main() {
	file, err := os.Open("2018/04/input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input, err := ParseEvents(scanner)
	check(err)

	// Count the minutes each guard is asleep
	start := time.Now()
	sleep := CountSleep(input)
	mostAsleepGuard := MostAsleep(sleep)
	minute := MostCommonAsleepMinute(sleep, mostAsleepGuard)

	fmt.Printf("Guard #%d at Minute %d == %d\n", mostAsleepGuard, minute, mostAsleepGuard*int64(minute))
	fmt.Println(time.Since(start))

	start = time.Now()
	rGuard, rMinute := MostReliablyAsleep(sleep)
	fmt.Printf("Most Reliably Guard #%d at Minute %d == %d\n", rGuard, rMinute, rGuard*int64(rMinute))
	fmt.Println(time.Since(start))
}
