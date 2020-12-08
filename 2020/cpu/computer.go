package cpu

// Computer represents a CPU capable of running programs in this
// arbitrary assembler language
type Computer struct {
	// IP represents the instruction pointer.  Acc represents the value in the accumulator register.
	IP, Acc    Data
	Inst       []Instruction
	Terminated bool
}

// NewComputer initializes a new Computer object with the provided program.
func NewComputer(program []Instruction) Computer {
	c := Computer{
		Inst: program,
	}
	c.Reset()
	return c
}

// Reset rests the program and registers to their initial states
func (c *Computer) Reset() {
	c.IP = 0
	c.Acc = 0
	c.Terminated = false
}

// Step executes the current instruction.  Returns true if the program should
// continue executing and false if it has terminated.
func (c *Computer) Step() bool {
	i := &c.Inst[c.IP]
	switch i.Op {
	case OpNop:
		c.IP++
	case OpAcc:
		c.Acc += i.Num
		c.IP++
	case OpJmp:
		c.IP += i.Num
	}

	// Termination is indicated by jumping to the instruction immediately after
	// the program ends
	c.Terminated = c.IP == Data(len(c.Inst))
	return !c.Terminated
}

// FindInfiniteLoop runs the current program until it finds an infinite loop.
// Returns true if a loop found and false if the program terminates.
func (c *Computer) FindInfiniteLoop() bool {
	seen := make(map[Data]bool)
	for !seen[c.IP] && !c.Terminated {
		seen[c.IP] = true
		c.Step()
	}
	return !c.Terminated
}
