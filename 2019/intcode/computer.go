package intcode

import (
	"fmt"
	"sort"
	"strconv"
)

// A class for simulating an IntCode computer.  IntCode is an imaginary language used in the 2019
// Advent of Code competition

// debug can be set to true to print debugging information as the program runs
const debug = false

// Computer represents an IntCode computer, capable of simulating instructions.  Internals such
// as the state of memory, registers, etc. are directly exposed for manipulation while the machine
// is running
type Computer struct {
	// Memory represents the current contents of memory
	Memory map[Address]Instruction
	// IP contains the instruction pointer (the address of the instruction that will next be read)
	IP Address
	// A copy of the program initially loaded into this machine
	program Program
	// Whether the program is still running or has terminated (naturally or by crashing)
	running bool
	// Whether the program crashed
	crashed bool
}

// NewComputer creates a new computer which has loaded the provided program
func NewComputer(program Program) Computer {
	c := Computer{}
	c.LoadProgram(program)
	return c
}

// LoadProgram loads a new program into the computer, clearing all memory and state
func (c *Computer) LoadProgram(program Program) {
	// Discard the current program and memory
	c.Memory = make(map[Address]Instruction)
	c.program = make(Program, len(program))
	copy(c.program, program)
	// Copy the program into memory
	for idx, val := range program {
		c.Memory[Address(idx)] = val
	}
	c.IP = 0
	c.running = true
	c.crashed = false

	if (debug) {
		fmt.Print("Loaded Program: ")
		fmt.Println(program)
	}
}

// IsRunning returns true if the program has not ended
func (c *Computer) IsRunning() bool {
	return c.running
}

// DidCrash returns true if the program crashed (as opposed to terminating normally)
func (c *Computer) DidCrash() bool {
	return c.crashed
}

// Run runs the current program to completion.  Returns true if the program successfully terminated
// and false if it crashed (by accessing an unknown memory address, for instance)
func (c *Computer) Run() bool {
	for ; c.running; {
		c.Step()
	}
	return !c.crashed
}

// Set the computer into a crashed state
func (c *Computer) crash() {
	c.running = false
	c.crashed = true
	if (debug) {
		fmt.Printf("*** Crashed at IP %d.  Program Contents: %s\n", c.IP, c.memoryContents())
	}
}

// Step runs a single instruction in the program.  Returns true if the program is still running and false
// if it has ended (due to termination or crash)
func (c *Computer) Step() bool {
	if (c.running) {
		if (debug) {
			fmt.Printf("IP %d: ", c.IP)
		}

		if (!c.validAddress(c.IP)) {
			c.crash()
			return false
		}
		op := c.Memory[c.IP]
		arg1 := c.IP + 1
		arg2 := c.IP + 2
		arg3 := c.IP + 3
		switch op {
		case OpAdd:
			if !(c.validAddress(arg1) && c.validAddress(arg2) && c.validAddress(arg3)) {
				c.crash()
				return false
			}
			target := c.Memory[arg3]
			s1 := c.Memory[arg1]
			s2 := c.Memory[arg2]
			// TODO: If writing to previously-unknown memory address is supported, remove the target check
			if !(c.validAddress(target) && c.validAddress(s1) && c.validAddress(s2)) {
				c.crash()
				return false
			}
			if (debug) {
				fmt.Printf("i%d := i%d(%d) + i%d(%d)\n", target, s1, c.Memory[s1], s2, c.Memory[s2])
			}
			c.Memory[target] = c.Memory[s1] + c.Memory[s2]
		case OpMultiply:
			if !(c.validAddress(arg1) && c.validAddress(arg2) && c.validAddress(arg3)) {
				c.crash()
				return false
			}
			target := c.Memory[arg3]
			s1 := c.Memory[arg1]
			s2 := c.Memory[arg2]
			// TODO: If writing to previously-unknown memory address is supported, remove the target check
			if !(c.validAddress(target) && c.validAddress(s1) && c.validAddress(s2)) {
				c.crash()
				return false
			}
			if (debug) {
				fmt.Printf("i%d := i%d(%d) * i%d(%d)\n", target, s1, c.Memory[s1], s2, c.Memory[s2])
			}
			c.Memory[target] = c.Memory[s1] * c.Memory[s2]
		case OpTerminate:
			if (debug) {
				fmt.Println("TERMINATE")
			}
			c.running = false
		default:
			fmt.Printf("Unexpected instruction: %d\n", op)
		}
		c.IP += (argCount(op) + 1)	// Add one to account for op itself
		if (debug) {
			fmt.Println(c.memoryContents())
		}
	}

	return c.running
}

// validAddress indicates whether a memory address is valid
func (c *Computer) validAddress(a Address) bool {
	_, exists := c.Memory[a]
	return exists
}

// programMemoryContents returns a string containing the contents of the instructions corresponding to the program
// in memory, in order
func (c *Computer) memoryContents() string {
	proglen := int64(len(c.program))
	
	// Get the contents of the memory for the program address space
	m := "["
	for i := int64(0); i < (proglen - 1); i++ {
		m += strconv.FormatInt(c.Memory[i], 10)
		m += ","
	}
	// Add the last element
	m += strconv.FormatInt(c.Memory[proglen - 1], 10)
	m += "]"

	if int64(len(c.Memory)) > proglen {
		// Find any memory addresses from outside the program space
		addresses := []int64{}
		for addr := range c.Memory {
			if addr < 0 || addr >= proglen {
				addresses = append(addresses, addr)
			}
		}
		// Sort them
		sort.Slice(addresses, func(i, j int) bool {
			return addresses[i] < addresses[j]
		})
		// Display them
		for _, addr := range addresses {
			m += "\n" + strconv.FormatInt(addr, 10) + ": " + strconv.FormatInt(c.Memory[addr], 10)
		}
	}

	return m
}