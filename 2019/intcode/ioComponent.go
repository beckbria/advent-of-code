package intcode

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
	// Per program definition, always input 1
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
