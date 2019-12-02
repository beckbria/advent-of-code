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
	OpAdd 		Instruction = 1
	OpMultiply 	Instruction = 2
	OpTerminate	Instruction = 99
)

// Returns the number of arguments for a given instruction
func argCount(i Instruction) int64 {
	switch i {
	case OpAdd, OpMultiply:
		return 3
	case OpTerminate:
		return 0
	default:
		log.Fatalf("Unknown instruction passed to argCount: %d\n", i)
	}
	return 0
}
