package intcode

import "fmt"

// InputOutput represents the inputs and outputs to a computer
type InputOutput interface {
	// GetInput should return the next value input to the computer
	GetInput() int64
	// Output is called each time the computer wants to output a value
	Output(int64)
	// Reset is called when the input and output buffers should be reset to their initial state
	Reset()
}

// ConstantInputOutput represents an I/O interface that always returns a constant
// value whenever input is requested
type ConstantInputOutput struct {
	input   int64
	Outputs []int64
}

// GetInput returns the provided input value
func (io *ConstantInputOutput) GetInput() int64 {
	return io.input
}

// Output collects the output value into a slice for later use
func (io *ConstantInputOutput) Output(o int64) {
	io.Outputs = append(io.Outputs, o)
}

// Reset resets the output buffer
func (io *ConstantInputOutput) Reset() {
	io.Outputs = []int64{}
}

// NewConstantInputOutput creates a new I/O module which returns a constant input value
func NewConstantInputOutput(input int64) *ConstantInputOutput {
	io := ConstantInputOutput{Outputs: []int64{}, input: input}
	return &io
}

// StreamInputOutput represents an IO component which returns inputs from a stream until it is exhausted
type StreamInputOutput struct {
	initialInputs  []int64
	inputs         []int64
	nextInputIndex int
	Outputs        []int64
	Debug          bool
}

// GetInput returns the provided input value
func (io *StreamInputOutput) GetInput() int64 {
	input := io.inputs[io.nextInputIndex]
	if io.Debug {
		fmt.Printf("Input: %d\n", input)
	}
	io.nextInputIndex++
	return input
}

// AppendInput adds another input to the queue
func (io *StreamInputOutput) AppendInput(i int64) {
	io.inputs = append(io.inputs, i)
}

// Output collects the output value into a slice for later use
func (io *StreamInputOutput) Output(o int64) {
	if io.Debug {
		fmt.Printf("Output: %d\n", o)
	}
	io.Outputs = append(io.Outputs, o)
}

// Reset resets the output buffer
func (io *StreamInputOutput) Reset() {
	io.inputs = make([]int64, len(io.initialInputs))
	copy(io.inputs, io.initialInputs)
	io.Outputs = []int64{}
	io.nextInputIndex = 0
}

// LastOutput returns the most recent value output into the buffer
func (io *StreamInputOutput) LastOutput() int64 {
	return io.Outputs[len(io.Outputs)-1]
}

// NewStreamInputOutput creates an IO component which returns a fixed series of inputs
func NewStreamInputOutput(input []int64) *StreamInputOutput {
	io := StreamInputOutput{Outputs: []int64{}, inputs: input, nextInputIndex: 0}
	return &io
}

// ProducerConsumerInputOutput is an IO consumer with Go channels designed to be interconnected
type ProducerConsumerInputOutput struct {
	InputChan  chan int64
	OutputChan chan int64
	OutputLog  []int64
}

// GetInput returns the provided input value
func (io *ProducerConsumerInputOutput) GetInput() int64 {
	i := <-io.InputChan
	return i
}

// Output collects the output value into a slice for later use
func (io *ProducerConsumerInputOutput) Output(o int64) {
	io.OutputLog = append(io.OutputLog, o)
	io.OutputChan <- o
}

// Reset resets the output buffer
func (io *ProducerConsumerInputOutput) Reset() {
	io.InputChan = make(chan int64)
	io.OutputChan = make(chan int64)
	io.OutputLog = []int64{}
}

// LastOutput returns the most recent value output into the buffer
func (io *ProducerConsumerInputOutput) LastOutput() int64 {
	return io.OutputLog[len(io.OutputLog)-1]
}

// NewProducerConsumerInputOutput creates an IO component with Go channels for interconnection
func NewProducerConsumerInputOutput() *ProducerConsumerInputOutput {
	io := ProducerConsumerInputOutput{}
	io.Reset()
	return &io
}

// ASCIIIo is an I/O component which reads and writes ASCII strings to an Intcode computer
type ASCIIIo struct {
	inputs         []int64
	nextInputIndex int
	Outputs        []int64
}

// GetInput returns the provided input value
func (io *ASCIIIo) GetInput() int64 {
	input := io.inputs[io.nextInputIndex]
	if debug {
		fmt.Printf("Read %d\n", input)
	}
	io.nextInputIndex++
	return input
}

// AppendInput adds a string message to the input queue
func (io *ASCIIIo) AppendInput(i string) {
	chars := []int64{}
	for _, r := range []rune(i) {
		chars = append(chars, int64(r))
	}
	chars = append(chars, 10) // Trailing newline
	io.inputs = append(io.inputs, chars...)
}

// Output collects the output value into a slice for later use
func (io *ASCIIIo) Output(o int64) {
	io.Outputs = append(io.Outputs, o)
}

// ClearOutput clears the output buffer
func (io *ASCIIIo) ClearOutput() {
	io.Outputs = []int64{}
}

// OutputAsString converts the output buffer into a string for display
func (io *ASCIIIo) OutputAsString() string {
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

// Reset resets the input and output buffers
func (io *ASCIIIo) Reset() {
	io.inputs = []int64{}
	io.Outputs = []int64{}
	io.nextInputIndex = 0
}

// NewASCIIIo creates a new ASCIIIo component
func NewASCIIIo() *ASCIIIo {
	a := ASCIIIo{}
	a.Reset()
	return &a
}
