package main

import (
	"fmt"

	"../aoc"
	"../intcode"
)

// https://adventofcode.com/2019/day/5
// Implement new instructions in the IntCode computer, run a program,
// return the diagnostic code it emits immediately prior to terminating

func main() {
	program := intcode.ReadIntCode("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	maxSignal, _ := maxThrusterSignal(program)
	fmt.Println(maxSignal)
	// Part 2
	maxSignalLoopback, _ := maxThrusterSignalLoopback(program)
	fmt.Println(maxSignalLoopback)
	fmt.Println(sw.Elapsed())
}

const ampCount = 5

func maxThrusterSignal(program intcode.Program) (int64, []int64) {
	maxSignal := int64(-9999999)
	maxSignalOrder := make([]int64, ampCount)

	computers := []intcode.Computer{}
	order := []int64{}
	for i := 0; i < ampCount; i++ {
		computers = append(computers, intcode.NewComputer(program))
		order = append(order, int64(i))
	}

	// Iterate through all computer orders
	for ok := true; ok; ok = aoc.NextPermutation(order) {
		for i := 0; i < len(computers); i++ {
			computers[i].Reset()
		}
		inputA := intcode.NewFixedInputOutput([]int64{order[0], 0})
		computers[0].Io = inputA
		computers[0].Run()
		// Connect the output of A to the input of B
		inputB := intcode.NewFixedInputOutput([]int64{order[1], inputA.LastOutput()})
		computers[1].Io = inputB
		computers[1].Run()
		inputC := intcode.NewFixedInputOutput([]int64{order[2], inputB.LastOutput()})
		computers[2].Io = inputC
		computers[2].Run()
		inputD := intcode.NewFixedInputOutput([]int64{order[3], inputC.LastOutput()})
		computers[3].Io = inputD
		computers[3].Run()
		inputE := intcode.NewFixedInputOutput([]int64{order[4], inputD.LastOutput()})
		computers[4].Io = inputE
		computers[4].Run()

		signal := inputE.LastOutput()
		if signal > maxSignal {
			maxSignal = signal
			for i := range order {
				maxSignalOrder[i] = order[i]
			}
		}
	}

	return maxSignal, maxSignalOrder
}

func maxThrusterSignalLoopback(program intcode.Program) (int64, []int64) {
	maxSignal := int64(-9999999)
	maxSignalOrder := make([]int64, ampCount)

	computers := []intcode.Computer{}
	order := []int64{}
	for i := 0; i < ampCount; i++ {
		computers = append(computers, intcode.NewComputer(program))
		order = append(order, int64(5+i))
	}

	// Iterate through all computer orders
	for ok := true; ok; ok = aoc.NextPermutation(order) {
		for i := 0; i < len(computers); i++ {
			computers[i].Reset()
		}

		io := []*intcode.FixedInputOutput{
			intcode.NewFixedInputOutput([]int64{order[0]}),
			intcode.NewFixedInputOutput([]int64{order[1]}),
			intcode.NewFixedInputOutput([]int64{order[1]}),
			intcode.NewFixedInputOutput([]int64{order[1]}),
			intcode.NewFixedInputOutput([]int64{order[1]})}
		for i := 0; i < 5; i++ {
			computers[i].Io = io[i]
		}

		signal := int64(0)
		for done := false; !done; {
			// Run each program until it outputs something
			io[0].AddInput(signal)
			oc := len(io[0].Outputs)
			for computers[0].IsRunning() && len(io[0].Outputs) == oc {
				computers[0].Step()
			}
			if len(io[0].Outputs) > oc {
				io[1].AddInput(io[0].LastOutput())
			}
			oc = len(io[1].Outputs)
			for computers[1].IsRunning() && len(io[1].Outputs) == oc {
				computers[1].Step()
			}
			if len(io[1].Outputs) > oc {
				io[2].AddInput(io[1].LastOutput())
			}
			oc = len(io[2].Outputs)
			for computers[2].IsRunning() && len(io[2].Outputs) == oc {
				computers[2].Step()
			}
			if len(io[2].Outputs) > oc {
				io[3].AddInput(io[2].LastOutput())
			}
			oc = len(io[3].Outputs)
			for computers[3].IsRunning() && len(io[3].Outputs) == oc {
				computers[3].Step()
			}
			if len(io[3].Outputs) > oc {
				io[4].AddInput(io[3].LastOutput())
			}
			oc = len(io[4].Outputs)
			for computers[4].IsRunning() && len(io[4].Outputs) == oc {
				computers[4].Step()
			}
			signal = io[4].LastOutput()

			done = !(computers[0].IsRunning() && computers[1].IsRunning() && computers[2].IsRunning() && computers[3].IsRunning() && computers[4].IsRunning())
		}

		if signal > maxSignal {
			maxSignal = signal
			for i := range order {
				maxSignalOrder[i] = order[i]
			}
		}
	}

	return maxSignal, maxSignalOrder
}
