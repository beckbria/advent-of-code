package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	// Input format: "# #"
	beforeRegEx = regexp.MustCompile(`^Before: \[(\d+), (\d+), (\d+), (\d+)\]$`)
	instRegEx   = regexp.MustCompile(`^(\d+) (\d+) (\d+) (\d+)$`)
	afterRegEx  = regexp.MustCompile(`^After:  \[(\d+), (\d+), (\d+), (\d+)\]$`)
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Registers represents the four integer registers of this imaginary CPU
type Registers [4]int

// Instruction represents a CPU instruction, optionally with before/after register state
type Instruction struct {
	regBefore Registers
	regAfter  Registers
	opCode    int
	A         int
	B         int
	dest      int
}

// ReadInstruction parses a line from the input file for part #1
func ReadInstruction(input []string) Instruction {
	var inst Instruction
	beforeTokens := beforeRegEx.FindStringSubmatch(input[0])
	a, err := strconv.Atoi(beforeTokens[1])
	check(err)
	b, err := strconv.Atoi(beforeTokens[2])
	check(err)
	c, err := strconv.Atoi(beforeTokens[3])
	check(err)
	d, err := strconv.Atoi(beforeTokens[4])
	check(err)
	inst.regBefore = Registers{a, b, c, d}

	opCode, A, B, dest := readRawInstruction(input[1])
	inst.opCode, inst.A, inst.B, inst.dest = opCode, A, B, dest

	afterTokens := afterRegEx.FindStringSubmatch(input[2])
	a, err = strconv.Atoi(afterTokens[1])
	check(err)
	b, err = strconv.Atoi(afterTokens[2])
	check(err)
	c, err = strconv.Atoi(afterTokens[3])
	check(err)
	d, err = strconv.Atoi(afterTokens[4])
	check(err)
	inst.regAfter = Registers{a, b, c, d}

	return inst
}

// ReadRawInstruction reads only the four numbers that go into an instruction without any
// values for registers before or after
func readRawInstruction(inst string) (int, int, int, int) {
	instTokens := instRegEx.FindStringSubmatch(inst)
	opCode, err := strconv.Atoi(instTokens[1])
	check(err)
	A, err := strconv.Atoi(instTokens[2])
	check(err)
	B, err := strconv.Atoi(instTokens[3])
	check(err)
	dest, err := strconv.Atoi(instTokens[4])
	check(err)
	return opCode, A, B, dest
}

// ReadRawInstructions reads a list of raw instructions
func ReadRawInstructions(input []string) []Instruction {
	var instructions []Instruction
	for _, s := range input {
		opCode, A, B, dest := readRawInstruction(s)
		var inst Instruction
		inst.opCode, inst.A, inst.B, inst.dest = opCode, A, B, dest
		instructions = append(instructions, inst)
	}
	return instructions
}

// ReadInstructions reads the instructions from part 1 of the data file
func ReadInstructions(input []string) []Instruction {
	// Read input in chunks of 4
	var instructions []Instruction
	for i := 0; i < len(input); i += 4 {
		inst := ReadInstruction(input[i : i+4])
		instructions = append(instructions, inst)
	}
	return instructions
}

// equalReg returns true if the set of registers contain the same contents
func equalReg(a, b Registers) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Addr (add register) stores into register C the result of Adding register A and register B.
func Addr(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] + reg[inst.B]
	return reg
}

// Addi (add immediate) stores into register C the result of Adding register A and value B.
func Addi(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] + inst.B
	return reg
}

// Mulr (multiply register) stores into register C the result of multiplying register A and register B.
func Mulr(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] * reg[inst.B]
	return reg
}

// Muli (multiply immediate) stores into register C the result of multiplying register A and value B.
func Muli(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] * inst.B
	return reg
}

// Banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
func Banr(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] & reg[inst.B]
	return reg
}

// Bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func Bani(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] & inst.B
	return reg
}

// Borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
func Borr(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] | reg[inst.B]
	return reg
}

// Bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
func Bori(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A] | inst.B
	return reg
}

// Setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func Setr(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = reg[inst.A]
	return reg
}

// Seti (set immediate) stores value A into register C. (Input B is ignored.)
func Seti(inst Instruction) Registers {
	reg := inst.regBefore
	reg[inst.dest] = inst.A
	return reg
}

// Gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
func Gtir(inst Instruction) Registers {
	reg := inst.regBefore
	if inst.A > reg[inst.B] {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// Gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func Gtri(inst Instruction) Registers {
	reg := inst.regBefore
	if reg[inst.A] > inst.B {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// Gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
func Gtrr(inst Instruction) Registers {
	reg := inst.regBefore
	if reg[inst.A] > reg[inst.B] {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// Eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
func Eqir(inst Instruction) Registers {
	reg := inst.regBefore
	if inst.A == reg[inst.B] {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// Eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func Eqri(inst Instruction) Registers {
	reg := inst.regBefore
	if reg[inst.A] == inst.B {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// Eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
func Eqrr(inst Instruction) Registers {
	reg := inst.regBefore
	if reg[inst.A] == reg[inst.B] {
		reg[inst.dest] = 1
	} else {
		reg[inst.dest] = 0
	}
	return reg
}

// CouldBeThree counts how many instructions could theoretically be three or more different instructions from the set
func CouldBeThree(inst []Instruction) int {
	count := 0

	for _, i := range inst {
		if len(ValidOpcodes(i)) >= 3 {
			count++
		}
	}

	return count
}

// ValidOpcodes returns each opcode that might be valid for the provided instruction -
// that is, whether (regardless of the opcode) the transform applied would produce the
// provided output
func ValidOpcodes(i Instruction) map[string]bool {
	candidates := make(map[string]bool)

	if equalReg(i.regAfter, Addr(i)) {
		candidates["addr"] = true
	}
	if equalReg(i.regAfter, Addi(i)) {
		candidates["addi"] = true
	}
	if equalReg(i.regAfter, Mulr(i)) {
		candidates["mulr"] = true
	}
	if equalReg(i.regAfter, Muli(i)) {
		candidates["muli"] = true
	}
	if equalReg(i.regAfter, Banr(i)) {
		candidates["banr"] = true
	}
	if equalReg(i.regAfter, Bani(i)) {
		candidates["bani"] = true
	}
	if equalReg(i.regAfter, Borr(i)) {
		candidates["borr"] = true
	}
	if equalReg(i.regAfter, Bori(i)) {
		candidates["bori"] = true
	}
	if equalReg(i.regAfter, Setr(i)) {
		candidates["setr"] = true
	}
	if equalReg(i.regAfter, Seti(i)) {
		candidates["seti"] = true
	}
	if equalReg(i.regAfter, Gtir(i)) {
		candidates["gtir"] = true
	}
	if equalReg(i.regAfter, Gtri(i)) {
		candidates["gtri"] = true
	}
	if equalReg(i.regAfter, Gtrr(i)) {
		candidates["gtrr"] = true
	}
	if equalReg(i.regAfter, Eqir(i)) {
		candidates["eqir"] = true
	}
	if equalReg(i.regAfter, Eqri(i)) {
		candidates["eqri"] = true
	}
	if equalReg(i.regAfter, Eqrr(i)) {
		candidates["eqrr"] = true
	}

	return candidates
}

// Find the intersection of two string lists
func intersection(a, b map[string]bool) map[string]bool {
	seen := make(map[string]bool)
	for k := range a {
		if _, present := b[k]; present {
			seen[k] = true
		}
	}
	return seen
}

func findOpcodes(inst []Instruction) map[int]string {
	validOpcodes := make(map[int]map[string]bool)

	for _, i := range inst {
		vo := ValidOpcodes(i)
		if _, present := validOpcodes[i.opCode]; present {
			validOpcodes[i.opCode] = intersection(validOpcodes[i.opCode], vo)
		} else {
			validOpcodes[i.opCode] = vo
		}
	}

	actualOpcodes := make(map[int]string)

	for len(actualOpcodes) < 16 {
		for intKey, validList := range validOpcodes {
			if len(validList) == 1 {
				for opcode := range validList {
					actualOpcodes[intKey] = opcode
					for _, list := range validOpcodes {
						// Delete from the seen list
						delete(list, opcode)
					}
				}
			}
		}
	}

	return actualOpcodes
}

// This simulates an actual program from raw instructions
func runProgram(opcodes map[int]string, inst []Instruction) Registers {
	reg := Registers{0, 0, 0, 0}
	for _, i := range inst {
		i.regBefore = reg
		switch opcodes[i.opCode] {
		case "eqri":
			reg = Eqri(i)
		case "banr":
			reg = Banr(i)
		case "bori":
			reg = Bori(i)
		case "mulr":
			reg = Mulr(i)
		case "seti":
			reg = Seti(i)
		case "bani":
			reg = Bani(i)
		case "muli":
			reg = Muli(i)
		case "gtrr":
			reg = Gtrr(i)
		case "setr":
			reg = Setr(i)
		case "addi":
			reg = Addi(i)
		case "gtir":
			reg = Gtir(i)
		case "borr":
			reg = Borr(i)
		case "addr":
			reg = Addr(i)
		case "eqrr":
			reg = Eqrr(i)
		case "gtri":
			reg = Gtri(i)
		case "eqir":
			reg = Eqir(i)
		}
	}
	return reg
}

func main() {
	file, err := os.Open("2018/16/input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	file2, err := os.Open("input2.txt")
	check(err)
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	var input2 []string
	for scanner.Scan() {
		input2 = append(input2, scanner.Text())
	}

	// Calculate Part 1
	start := time.Now()
	inst := ReadInstructions(input)
	fmt.Println(CouldBeThree(inst))
	fmt.Println(time.Since(start))
	start = time.Now()

	// Calculate Part 2
	opcodes := findOpcodes(inst)
	inst2 := ReadRawInstructions(input2)
	reg := runProgram(opcodes, inst2)
	fmt.Println(reg[0])
	fmt.Println(time.Since(start))
}
