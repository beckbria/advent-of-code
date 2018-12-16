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

var (
	// Input format: "# #"
	beforeRegEx = regexp.MustCompile("^Before: \\[(\\d+), (\\d+), (\\d+), (\\d+)\\]$")
	instRegEx   = regexp.MustCompile("$(\\d+) (\\d+) (\\d+) (\\d+)$") // TODO: Determine why this doesn't work
	afterRegEx  = regexp.MustCompile("^After:  \\[(\\d+), (\\d+), (\\d+), (\\d+)\\]$")
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
	instTokens := strings.Split(inst, " ")
	opCode, err := strconv.Atoi(instTokens[0])
	check(err)
	A, err := strconv.Atoi(instTokens[1])
	check(err)
	B, err := strconv.Atoi(instTokens[2])
	check(err)
	dest, err := strconv.Atoi(instTokens[3])
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
func ValidOpcodes(i Instruction) []string {
	candidates := make([]string, 0)

	if equalReg(i.regAfter, Addr(i)) {
		candidates = append(candidates, "addr")
	}
	if equalReg(i.regAfter, Addi(i)) {
		candidates = append(candidates, "addi")
	}
	if equalReg(i.regAfter, Mulr(i)) {
		candidates = append(candidates, "mulr")
	}
	if equalReg(i.regAfter, Muli(i)) {
		candidates = append(candidates, "muli")
	}
	if equalReg(i.regAfter, Banr(i)) {
		candidates = append(candidates, "banr")
	}
	if equalReg(i.regAfter, Bani(i)) {
		candidates = append(candidates, "bani")
	}
	if equalReg(i.regAfter, Borr(i)) {
		candidates = append(candidates, "borr")
	}
	if equalReg(i.regAfter, Bori(i)) {
		candidates = append(candidates, "bori")
	}
	if equalReg(i.regAfter, Setr(i)) {
		candidates = append(candidates, "setr")
	}
	if equalReg(i.regAfter, Seti(i)) {
		candidates = append(candidates, "seti")
	}
	if equalReg(i.regAfter, Gtir(i)) {
		candidates = append(candidates, "gtir")
	}
	if equalReg(i.regAfter, Gtri(i)) {
		candidates = append(candidates, "gtri")
	}
	if equalReg(i.regAfter, Gtrr(i)) {
		candidates = append(candidates, "gtrr")
	}
	if equalReg(i.regAfter, Eqir(i)) {
		candidates = append(candidates, "eqir")
	}
	if equalReg(i.regAfter, Eqri(i)) {
		candidates = append(candidates, "eqri")
	}
	if equalReg(i.regAfter, Eqrr(i)) {
		candidates = append(candidates, "eqrr")
	}

	return candidates
}

// Find the intersection of two string lists
func intersection(a, b []string) []string {
	seen := make(map[string]bool)
	for _, s := range a {
		seen[s] = true
	}
	intersect := make([]string, 0)
	for _, s := range b {
		if _, exists := seen[s]; exists {
			intersect = append(intersect, s)
		}
	}
	return intersect
}

// dumpOpcodes prints out a list of what opcodes a given number could be from the
// test data.  Manual analysis gives the following list:
//0: [eqri]
//1: [banr]
//2: [bori]
//3: [mulr]
//4: [seti]
//5: [bani]
//6: [muli]
//7: [gtrr]
//8: [setr]
//9: [addi]
//10: [gtir]
//11: [borr]
//12: [addr]
//13: [eqrr]
//14: [gtri]
//15: [eqir]
func dumpOpcodes(inst []Instruction) {
	validOpcodes := make(map[int][]string)

	for _, i := range inst {
		vo := ValidOpcodes(i)
		if _, present := validOpcodes[i.opCode]; present {
			validOpcodes[i.opCode] = intersection(validOpcodes[i.opCode], vo)
		} else {
			validOpcodes[i.opCode] = vo
		}
	}

	for k, v := range validOpcodes {
		op := fmt.Sprintln(v)
		fmt.Printf("%d: %s", k, op)
	}
}

// This simulates an actual program from raw instructions
func runProgram(inst []Instruction) Registers {
	reg := Registers{0, 0, 0, 0}
	for _, i := range inst {
		i.regBefore = reg
		switch i.opCode {
		case 0:
			reg = Eqri(i)
		case 1:
			reg = Banr(i)
		case 2:
			reg = Bori(i)
		case 3:
			reg = Mulr(i)
		case 4:
			reg = Seti(i)
		case 5:
			reg = Bani(i)
		case 6:
			reg = Muli(i)
		case 7:
			reg = Gtrr(i)
		case 8:
			reg = Setr(i)
		case 9:
			reg = Addi(i)
		case 10:
			reg = Gtir(i)
		case 11:
			reg = Borr(i)
		case 12:
			reg = Addr(i)
		case 13:
			reg = Eqrr(i)
		case 14:
			reg = Gtri(i)
		case 15:
			reg = Eqir(i)
		}
	}
	return reg
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
	// To see what opcodes are candidates for what numbers (and recreate the manual work), uncomment:
	//dumpOpcodes(inst)

	// Calculate Part 2
	inst2 := ReadRawInstructions(input2)
	reg := runProgram(inst2)
	fmt.Println(reg[0])
	fmt.Println(time.Since(start))
}
