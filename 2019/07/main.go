package main

import (
	"fmt"
	"math"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

// https://adventofcode.com/2019/day/7
// Chain multiple IntCode computers together via their IO components
// Find the ordering to produce a maximum value

const ampCount = 5

func main() {
	program := intcode.ReadIntCode("input.txt")
	sw := lib.NewStopwatch()
	// Part 1
	maxSignal, _ := maxThrusterSignal(program)
	fmt.Println(maxSignal)
	// Part 2
	maxSignalLoopback, _ := maxThrusterSignalLoopback(program)
	fmt.Println(maxSignalLoopback)
	fmt.Println(sw.Elapsed())
}

func maxThrusterSignal(program intcode.Program) (int64, []int64) {
	maxSignal := int64(math.MinInt64)
	maxSignalOrder := make([]int64, ampCount)
	amps := []intcode.Computer{}
	phases := []int64{}
	for i := 0; i < ampCount; i++ {
		amps = append(amps, intcode.NewComputer(program))
		phases = append(phases, int64(i))
	}

	// Iterate through all computer orders
	for ok := true; ok; ok = lib.NextPermutation(phases) {
		previousOutput := int64(0)
		for i := 0; i < ampCount; i++ {
			amps[i].Reset()
			io := intcode.NewStreamInputOutput([]int64{phases[i], previousOutput})
			amps[i].Io = io
			amps[i].Run()
			previousOutput = io.LastOutput()
		}

		if previousOutput > maxSignal {
			maxSignal = previousOutput
			copy(maxSignalOrder, phases)
		}
	}
	return maxSignal, maxSignalOrder
}

func maxThrusterSignalLoopback(program intcode.Program) (int64, []int64) {
	maxSignal := int64(math.MinInt64)
	maxSignalOrder := make([]int64, ampCount)

	amps := []intcode.Computer{}
	phases := []int64{}
	for i := 0; i < ampCount; i++ {
		amps = append(amps, intcode.NewComputer(program))
		phases = append(phases, int64(5+i))
	}

	// Iterate through all phase orders
	for ok := true; ok; ok = lib.NextPermutation(phases) {
		io := []*intcode.StreamInputOutput{}
		for i := 0; i < ampCount; i++ {
			amps[i].Reset()
			io = append(io, intcode.NewStreamInputOutput([]int64{phases[i]}))
			amps[i].Io = io[i]
		}
		io[0].AppendInput(0) // Add the initial signal

		for amps[4].IsRunning() {
			// Run each program until it outputs something
			for i := 0; i < ampCount; i++ {
				oldOutputCount := len(io[i].Outputs)
				for amps[i].IsRunning() && len(io[i].Outputs) == oldOutputCount {
					amps[i].Step()
				}
				if len(io[i].Outputs) > oldOutputCount {
					io[(i+1)%ampCount].AppendInput(io[i].LastOutput())
				}
			}
		}
		signal := io[4].LastOutput()
		if signal > maxSignal {
			maxSignal = signal
			copy(maxSignalOrder, phases)
		}
	}
	return maxSignal, maxSignalOrder
}
