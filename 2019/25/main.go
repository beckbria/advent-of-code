package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../intcode"
)

const debug = false

func main() {
	p := intcode.ReadIntCode("input.txt")
	c := intcode.NewComputer(p)
	io := newAsciiIo()
	c.Io = io

	// Loop forever, displaying output and reading console input
	reader := bufio.NewReader(os.Stdin)
	for true {
		c.RunToNextInput()
		out := io.OutputAsString()
		if len(out) > 0 {
			fmt.Println(out)
			io.ClearOutput()
		}

		if !c.IsRunning() {
			break
		}

		// Only take more input if the computer has processed all current input and shown output
		if len(out) > 0 {
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if debug {
				fmt.Printf("Read <<%s>>\n", text)
			}
			io.AppendInput(text)
		}
		c.Step()
	}
}

type asciiIo struct {
	inputs         []int64
	nextInputIndex int
	Outputs        []int64
}

// GetInput returns the provided input value
func (io *asciiIo) GetInput() int64 {
	input := io.inputs[io.nextInputIndex]
	if debug {
		fmt.Printf("Read %d\n", input)
	}
	io.nextInputIndex++
	return input
}

// AppendInput adds another input to the queue
func (io *asciiIo) AppendInput(i string) {
	chars := []int64{}
	for _, r := range []rune(i) {
		chars = append(chars, int64(r))
	}
	chars = append(chars, 10) // Trailing newline
	io.inputs = chars
	io.nextInputIndex = 0
	if debug {
		fmt.Print("Sent to input: ")
		fmt.Println(io.inputs)
	}
}

// Output collects the output value into a slice for later use
func (io *asciiIo) Output(o int64) {
	io.Outputs = append(io.Outputs, o)
}

func (io *asciiIo) ClearOutput() {
	io.Outputs = []int64{}
}

func (io *asciiIo) OutputAsString() string {
	chars := []rune{}
	for _, i := range io.Outputs {
		if i == 10 {
			chars = append(chars, '\n')
		} else {
			chars = append(chars, rune(i))
		}
	}
	return string(chars)
}

func (io *asciiIo) Reset() {}

func newAsciiIo() *asciiIo {
	a := asciiIo{inputs: []int64{}, Outputs: []int64{}, nextInputIndex: 0}
	return &a
}
