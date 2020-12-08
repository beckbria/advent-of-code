package cpu

import "log"

// Data represents a numerical literal in the computer
type Data = int64

// OpCode represents an instruction
type OpCode = int64

// Instruction represents a single instruction to the computer
type Instruction struct {
	Op  OpCode
	Num Data
}

// Program represents a program to be ran in the CPU
type Program = []Instruction

const (
	// OpNop (No-op) does nothing and proceeds to the next instruction.
	OpNop = 1
	// OpAcc (Accumulate) adjusts the accumulator register by the value specified in
	// Num and proceeds to the next instruction.
	OpAcc = 2
	// OpJmp (Relative Jump) adjusts the instruction pointer by the amount specified
	// in Num.
	OpJmp = 3
)

func opToString(op OpCode) string {
	switch op {
	case OpNop:
		return "nop"
	case OpAcc:
		return "acc"
	case OpJmp:
		return "jmp"
	}
	log.Fatalf("Unknown opcode: %d", op)
	return ""
}

func opFromString(op string) OpCode {
	switch op {
	case "nop":
		return OpNop
	case "acc":
		return OpAcc
	case "jmp":
		return OpJmp
	}
	log.Fatalf("Unknown opcode: %q", op)
	return OpNop
}
