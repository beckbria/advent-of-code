package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const debug = false

var (
	// Input format: "# #"
	ipRegEx   = regexp.MustCompile("^\\#ip (\\d+)$")
	instRegEx = regexp.MustCompile("^([a-z]+) (\\d+) (\\d+) (\\d+)$")
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Op is an enum
type Op int

// Op enum entries
const (
	OpEqri = iota + 1
	OpBanr = iota + 1
	OpBori = iota + 1
	OpMulr = iota + 1
	OpSeti = iota + 1
	OpBani = iota + 1
	OpMuli = iota + 1
	OpGtrr = iota + 1
	OpSetr = iota + 1
	OpAddi = iota + 1
	OpGtir = iota + 1
	OpBorr = iota + 1
	OpAddr = iota + 1
	OpEqrr = iota + 1
	OpGtri = iota + 1
	OpEqir = iota + 1
)

// Registers represents the four integer registers of this imaginary CPU
type Registers [6]int

// Instruction represents a CPU instruction, optionally with before/after register state
type Instruction struct {
	opCode Op
	A      int
	B      int
	dest   int
}

// Computer represents a program and the CPU state
type Computer struct {
	ir               int // The Instruction Register
	ip               int // The instruction pointer - what the next instruction is
	reg              Registers
	inst             []Instruction
	instructionCount int
}

// ReadComputer parses the computer
func ReadComputer(input []string) Computer {
	ir, err := strconv.Atoi(strings.Split(input[0], " ")[1])
	check(err)
	var reg Registers // Zero-initialize
	inst := make([]Instruction, 0)
	for _, s := range input[1:] {
		inst = append(inst, readInstruction(s))
	}
	if debug {
		fmt.Printf("ir: %d\n", ir)
	}
	return Computer{ir: ir, ip: 0, reg: reg, inst: inst, instructionCount: 0}
}

// readInstruction parses a line from the input file for part #1
func readInstruction(input string) Instruction {
	instTokens := instRegEx.FindStringSubmatch(input)
	opCode, err := parseOpCode(instTokens[1])
	check(err)
	A, err := strconv.Atoi(instTokens[2])
	check(err)
	B, err := strconv.Atoi(instTokens[3])
	check(err)
	dest, err := strconv.Atoi(instTokens[4])
	check(err)

	return Instruction{opCode: opCode, A: A, B: B, dest: dest}
}

type opInfo struct {
	op  Op
	str string
}

var ops = []opInfo{
	{op: OpEqri, str: "eqri"},
	{op: OpBanr, str: "banr"},
	{op: OpBori, str: "bori"},
	{op: OpMulr, str: "mulr"},
	{op: OpSeti, str: "seti"},
	{op: OpBani, str: "bani"},
	{op: OpMuli, str: "muli"},
	{op: OpGtrr, str: "gtrr"},
	{op: OpSetr, str: "setr"},
	{op: OpAddi, str: "addi"},
	{op: OpGtir, str: "gtir"},
	{op: OpBorr, str: "borr"},
	{op: OpAddr, str: "addr"},
	{op: OpEqrr, str: "eqrr"},
	{op: OpGtri, str: "gtri"},
	{op: OpEqir, str: "eqir"}}

func parseOpCode(o string) (Op, error) {
	for _, opi := range ops {
		if o == opi.str {
			return opi.op, nil
		}
	}
	return -1, fmt.Errorf("Unknown opcode: %s", o)
}

func opCodeToString(o Op) string {
	for _, opi := range ops {
		if o == opi.op {
			return opi.str
		}
	}
	return fmt.Sprintf("Unknown opcode: %d", o)
}

// Addr (add register) stores into register C the result of Adding register A and register B.
func (c *Computer) Addr(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] + c.reg[inst.B]
}

// Addi (add immediate) stores into register C the result of Adding register A and value B.
func (c *Computer) Addi(inst Instruction) {
	if debug {
		fmt.Printf("reg[%d] = reg[%d] + %d = %d\n", inst.dest, inst.A, inst.B, c.reg[inst.A]+inst.B)
	}
	c.reg[inst.dest] = c.reg[inst.A] + inst.B
}

// Mulr (multiply register) stores into register C the result of multiplying register A and register B.
func (c *Computer) Mulr(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] * c.reg[inst.B]
}

// Muli (multiply immediate) stores into register C the result of multiplying register A and value B.
func (c *Computer) Muli(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] * inst.B
}

// Banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
func (c *Computer) Banr(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] & c.reg[inst.B]
}

// Bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func (c *Computer) Bani(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] & inst.B
}

// Borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
func (c *Computer) Borr(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] | c.reg[inst.B]
}

// Bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
func (c *Computer) Bori(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A] | inst.B
}

// Setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func (c *Computer) Setr(inst Instruction) {
	c.reg[inst.dest] = c.reg[inst.A]
}

// Seti (set immediate) stores value A into register C. (Input B is ignored.)
func (c *Computer) Seti(inst Instruction) {
	c.reg[inst.dest] = inst.A
}

// Gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
func (c *Computer) Gtir(inst Instruction) {
	if inst.A > c.reg[inst.B] {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

// Gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func (c *Computer) Gtri(inst Instruction) {
	if c.reg[inst.A] > inst.B {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

// Gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
func (c *Computer) Gtrr(inst Instruction) {
	if c.reg[inst.A] > c.reg[inst.B] {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

// Eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
func (c *Computer) Eqir(inst Instruction) {
	if inst.A == c.reg[inst.B] {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

// Eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func (c *Computer) Eqri(inst Instruction) {
	if c.reg[inst.A] == inst.B {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

// Eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
func (c *Computer) Eqrr(inst Instruction) {
	if c.reg[inst.A] == c.reg[inst.B] {
		c.reg[inst.dest] = 1
	} else {
		c.reg[inst.dest] = 0
	}
}

func (c *Computer) currentInstruction() Instruction {
	return c.inst[c.reg[c.ir]]
}

func (c *Computer) printReg() {
	fmt.Printf("[%d, %d, %d, %d, %d, %d]", c.reg[0], c.reg[1], c.reg[2], c.reg[3], c.reg[4], c.reg[5])
}

func printInstruction(i Instruction) {
	fmt.Printf(" %s %d %d %d ", opCodeToString(i.opCode), i.A, i.B, i.dest)
}

var regValues = make([]int, 0)
var regValueSeen = make(map[int]bool)

// This simulates an actual program from raw instructions
func (c *Computer) stepInstruction() bool {
	if (c.ip < 0) || (c.ip >= len(c.inst)) {
		if debug {
			fmt.Printf("ip=%d, terminating\n", c.ip)
		}
		return false
	}

	// From the assembly, the only time register 0 is used is to compare against
	// register 1 at line 30.  Thus, the only interesting values for r0 are whatever
	// r1 is when ip=30.  Log them - the first value will be the shortest terminating
	// condition, and the last value before it loops to a value seen before will be
	// the longest
	if c.ip == 30 {
		if _, present := regValueSeen[c.reg[1]]; present {
			// We've looped
			return false
		}
		regValues = append(regValues, c.reg[1])
		regValueSeen[c.reg[1]] = true
	}

	// Load the IP into its register
	c.reg[c.ir] = c.ip
	i := c.currentInstruction()
	c.instructionCount++

	if debug {
		fmt.Printf("ip=%d ", c.ip)
		c.printReg()
		printInstruction(i)
	}

	switch i.opCode {
	case OpEqri:
		c.Eqri(i)
	case OpBanr:
		c.Banr(i)
	case OpBori:
		c.Bori(i)
	case OpMulr:
		c.Mulr(i)
	case OpSeti:
		c.Seti(i)
	case OpBani:
		c.Bani(i)
	case OpMuli:
		c.Muli(i)
	case OpGtrr:
		c.Gtrr(i)
	case OpSetr:
		c.Setr(i)
	case OpAddi:
		c.Addi(i)
	case OpGtir:
		c.Gtir(i)
	case OpBorr:
		c.Borr(i)
	case OpAddr:
		c.Addr(i)
	case OpEqrr:
		c.Eqrr(i)
	case OpGtri:
		c.Gtri(i)
	case OpEqir:
		c.Eqir(i)
	}

	if debug {
		c.printReg()
		fmt.Printf("\n")
	}

	// Advance the instruction pointer
	c.ip = c.reg[c.ir] + 1
	return true
}

func (c *Computer) run() {
	for c.stepInstruction() {
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())

	// Calculate Part 1
	start := time.Now()
	comp := ReadComputer(input)
	comp.run()
	fmt.Println(regValues[0])                // First terminating value
	fmt.Println(regValues[len(regValues)-1]) // Final terminating value
	fmt.Println(time.Since(start))

}
