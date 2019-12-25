package intcode

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"../aoc"
)

// A class for simulating an IntCode computer.  IntCode is an imaginary language used in the 2019
// Advent of Code competition

// debug can be set to true to print debugging information as the program runs
const (
	debug                   = false
	debugDumpMemoryEachStep = debug && false
)

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
	// The IO component
	Io InputOutput
	// The Relative Base, used for relative mode instructions
	RelativeBase Instruction
}

// NewComputer creates a new computer which has loaded the provided program
func NewComputer(program Program) Computer {
	c := Computer{Io: nil}
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
	c.RelativeBase = 0
	c.running = true
	c.crashed = false

	if debug {
		fmt.Print("Loaded Program: ")
		fmt.Println(program)
	}
}

// Reset reloads the current program, resetting memory and other bits to their initial state
func (c *Computer) Reset() {
	c.LoadProgram(c.program)
	if c.Io != nil {
		c.Io.Reset()
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
	for c.running {
		c.Step()
	}
	return !c.crashed
}

// RunToNextInput runs the computer until its next input instruction
func (c *Computer) RunToNextInput() bool {
	for c.running {
		_, brokeOnInput := c.step(true)
		if brokeOnInput {
			if debug {
				fmt.Printf("Broke on input at IP %d\n", c.IP)
			}
			return true
		}
	}
	return false
}

// Set the computer into a crashed state
func (c *Computer) crash() {
	c.running = false
	c.crashed = true
	if debug {
		fmt.Printf("*** Crashed at IP %d.  Program Contents: %s\n", c.IP, c.memoryContents())
	}
}

// address modifies an address parameter to account for its mode.
// It returns the updated value and (if debugging is enabled) a
// debug representation of how the parameter was read
func (c *Computer) address(addr Address, mode ParameterMode) (Address, string) {
	target := c.Memory[addr]
	debugTarget := ""
	if mode == PmRelative {
		if debug {
			debugTarget = fmt.Sprintf("i[%d+%d]", target, c.RelativeBase)
		}
		target += c.RelativeBase
	} else if debug {
		debugTarget = fmt.Sprintf("i[%d]", target)
	}
	return target, debugTarget
}

// Step runs a single instruction in the program.  Returns true if the program is still running and false
// if it has ended (due to termination or crash)
func (c *Computer) Step() bool {
	running, _ := c.step(false)
	return running
}

// step runs a single instruction in the program.  Returns: First: true if the program is still running and false
// if it has ended (due to termination or crash).  Second: True if it broke on an input instruction
func (c *Computer) step(breakAtInput bool) (bool, bool) {
	if c.running {
		if debug {
			fmt.Printf("IP %d\t", c.IP)
		}

		op, rawOp, vals, modes, args, debugVals := c.readInstruction()
		if c.DidCrash() {
			return false, false
		}
		if debug {
			fmt.Printf("%d\t", rawOp)
			for i := 0; i < 4; i++ {
				if i < len(args) {
					fmt.Print(c.Memory[args[i]])
				}
				fmt.Print("\t")
			}
		}
		modifiedIP := false

		switch op {
		case OpAdd:
			target, debugTarget := c.address(args[2], modes[2])
			if debug {
				fmt.Printf("%s := %s + %s\n", debugTarget, debugVals[0], debugVals[1])
			}
			c.Memory[target] = vals[0] + vals[1]
		case OpMultiply:
			target, debugTarget := c.address(args[2], modes[2])
			if debug {
				fmt.Printf("%s := %s * %s\n", debugTarget, debugVals[0], debugVals[1])
			}
			c.Memory[target] = vals[0] * vals[1]
		case OpStore:
			if breakAtInput {
				return false, true
			}
			input := c.Io.GetInput()
			target, debugTarget := c.address(args[0], modes[0])
			if debug {
				fmt.Printf("%s := Input(%d)\n", debugTarget, input)
			}
			c.Memory[target] = input
		case OpOutput:
			output := vals[0]
			if debug {
				fmt.Printf("Output << %d\n", output)
			}
			c.Io.Output(output)
		case OpTerminate:
			if debug {
				fmt.Println("TERMINATE")
			}
			c.running = false
		case OpJumpIfTrue:
			if debug {
				fmt.Printf("IF %s GOTO %s: ", debugVals[0], debugVals[1])
			}
			if vals[0] != 0 {
				if debug {
					fmt.Println("Jumped")
				}
				c.IP = vals[1]
				modifiedIP = true
			} else if debug {
				fmt.Println("No")
			}
		case OpJumpIfFalse:
			if debug {
				fmt.Printf("IF !%s GOTO %s: ", debugVals[0], debugVals[1])
			}
			if vals[0] == 0 {
				if debug {
					fmt.Println("Jumped")
				}
				c.IP = vals[1]
				modifiedIP = true
			} else if debug {
				fmt.Println("No")
			}
		case OpLessThan:
			target, debugTarget := c.address(args[2], modes[2])
			if debug {
				fmt.Printf("%s := %s < %s ? 1 : 0 => ", debugTarget, debugVals[0], debugVals[1])
			}
			if vals[0] < vals[1] {
				c.Memory[target] = 1
				if debug {
					fmt.Println("1")
				}
			} else {
				c.Memory[target] = 0
				if debug {
					fmt.Println("0")
				}
			}
		case OpEquals:
			target, debugTarget := c.address(args[2], modes[2])
			if debug {
				fmt.Printf("%s := %s == %s ? 1 : 0 => ", debugTarget, debugVals[0], debugVals[1])
			}
			if vals[0] == vals[1] {
				c.Memory[target] = 1
				if debug {
					fmt.Println("1")
				}
			} else {
				c.Memory[target] = 0
				if debug {
					fmt.Println("0")
				}
			}
		case OpAdjustRelativeBase:
			if debug {
				fmt.Printf("RelativeBase += %s (%d)\n", debugVals[0], c.RelativeBase+vals[0])
			}
			c.RelativeBase += vals[0]
		default:
			fmt.Printf("Unexpected instruction: %d\n", op)
		}
		if !modifiedIP {
			c.IP += (int64(len(args)) + 1) // Add one to account for op itself
		}
		if debugDumpMemoryEachStep {
			fmt.Println(c.memoryContents())
		}
	}

	return c.running, false
}

// readInstruction parses the opcode at the instruction pointer and its parameters.
// Returns opcode, raw opcode, parameter values, parameter modes, memory addresses of parameters, and parameter debug strings for printing
func (c *Computer) readInstruction() (Instruction, Instruction, []Instruction, []ParameterMode, []Address, []string) {
	opCode := c.Memory[c.IP]
	op := opCode % 100
	argc := argCount(op)
	args := []Instruction{}
	modes := []ParameterMode{}
	vals := []Address{}
	debugVals := []string{}

	// Validate position mode arguments
	modeDivisor := int64(10)
	for i := int64(0); i < argc; i++ {
		modeDivisor *= 10
		mode := ParameterMode((opCode / modeDivisor) % 10)
		modes = append(modes, mode)
		arg := c.IP + i + 1
		args = append(args, arg)
		val := c.Memory[arg]
		debugVal := ""
		switch mode {
		case PmImmediate:
			if debug {
				debugVal = fmt.Sprintf("%d", val)
			}
		case PmPosition:
			if debug {
				debugVal = fmt.Sprintf("i%d(%d)", val, c.Memory[val])
			}
			val = c.Memory[val]
		case PmRelative:
			if debug {
				debugVal = fmt.Sprintf("i[%d+%d](%d)", c.RelativeBase, val, c.Memory[val+c.RelativeBase])
			}
			val = c.Memory[val+c.RelativeBase]
		default:
			log.Fatalf("Unexpected mode: %d", mode)
		}

		vals = append(vals, val)
		if debug {
			debugVals = append(debugVals, debugVal)
		}
	}

	return op, opCode, vals, modes, args, debugVals
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
	m += strconv.FormatInt(c.Memory[proglen-1], 10)
	m += "]"

	if int64(len(c.Memory)) > proglen {
		// Find any memory addresses from outside the program space
		addresses := []int64{}
		for addr := range c.Memory {
			if addr < 0 || addr >= proglen {
				addresses = append(addresses, addr)
			}
		}
		sort.Sort(aoc.Int64Slice(addresses))
		// Display them
		for _, addr := range addresses {
			m += "\n" + strconv.FormatInt(addr, 10) + ": " + strconv.FormatInt(c.Memory[addr], 10)
		}
	}

	return m
}
