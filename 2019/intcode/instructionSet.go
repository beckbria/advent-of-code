package intcode

import "log"

// Address represents a memory address in the IntCode computer
type Address = int64

// Instruction represents an IntCode instruction or memory value
type Instruction = int64

// Program represents an IntCode program.  IntCode programs are loaded into memory at position
// 0 and consist of a series of numbers representing the instructions/initial memory values
type Program = []Instruction

// Opcodes in the IntCode architecture
const (
	// OpAdd reads from two parameter and stores their sum in the third parameter
	OpAdd 			Instruction = 1
	// OpMultiply reads from two parameter and stores their product in the third parameter
	OpMultiply 		Instruction = 2
	// OpStore reads a integer from the I/O Component and saves it to the first parameter
	OpStore			Instruction = 3
	// OpOutput reads a value from the first parameter and outputs it to the I/O Component
	OpOutput		Instruction = 4
	// OpJumpIfTrue sets the IP to the second parameter if the first parameter is non-zero
	OpJumpIfTrue	Instruction = 5
	// OpJumpIfFalse sets the IP to the second parameter if the first parameter is zero
	OpJumpIfFalse	Instruction = 6
	// OpLessThan stores 1 in the third parameter if the first parameter is less than the second parameter, otherwise 0
	OpLessThan		Instruction = 7
	// OpEquals stores 1 in the third parameter if the first parameter is equal to the second parameter, otherwise 0
	OpEquals		Instruction = 8
	// OpTerminate immediately halts the program
	OpTerminate		Instruction = 99
)

// argCount returns the number of arguments for a given instruction
func argCount(i Instruction) int64 {
	switch i {
	case OpAdd, OpMultiply, OpLessThan, OpEquals:
		return 3
	case OpJumpIfTrue, OpJumpIfFalse:
		return 2
	case OpStore, OpOutput:
		return 1
	case OpTerminate:
		return 0
	default:
		log.Fatalf("Unknown instruction passed to argCount: %d\n", i)
	}
	return 0
}
