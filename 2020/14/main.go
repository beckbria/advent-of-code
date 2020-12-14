package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/14
// Obfuscated memory management

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	inst := readInstructions(lines)
	fmt.Println("Step 1:")
	fmt.Println(step1(inst))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(inst))
	fmt.Println(sw.Elapsed())
}

var (
	memRegex  = regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)
	maskRegex = regexp.MustCompile(`^mask = ([10X]+)$`)
)

type instruction struct {
	addr, value     uint64
	mask            string
	andMask, orMask uint64
}

func readInstructions(lines []string) []instruction {
	inst := make([]instruction, 0)
	for _, l := range lines {
		i := instruction{}
		tokens := memRegex.FindStringSubmatch(l)
		if tokens == nil {
			i.mask = maskRegex.FindStringSubmatch(l)[1]
			i.andMask, i.orMask = generateMasks(i.mask)
		} else {
			i.addr, _ = strconv.ParseUint(tokens[1], 10, 64)
			i.value, _ = strconv.ParseUint(tokens[2], 10, 64)
		}
		inst = append(inst, i)
	}
	return inst
}

const fullAndMask = uint64(0xffffffffffffffff)

func generateMasks(mask string) (uint64, uint64) {
	andMask := fullAndMask
	orMask := uint64(0)

	for i := 0; i < len(mask); i++ {
		bit := uint64(1) << (len(mask) - (i + 1))
		switch mask[i : i+1] {
		case "0":
			andMask &= ^bit
		case "1":
			orMask |= bit
		}
	}

	return andMask, orMask
}

// step1 applies the provided mask to the values being written to memory
func step1(inst []instruction) uint64 {
	andMask := fullAndMask
	orMask := uint64(0)
	mem := make(map[uint64]uint64)
	for _, i := range inst {
		if len(i.mask) > 0 {
			andMask = i.andMask
			orMask = i.orMask
		} else {
			mem[i.addr] = (i.value | orMask) & andMask
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

// step2 applies the provided mask to the addresses being written to
func step2(inst []instruction) uint64 {
	mask := ""
	mem := make(map[uint64]uint64)
	for _, i := range inst {
		if len(i.mask) > 0 {
			mask = i.mask
		} else {
			for _, a := range maskedAddresses(i.addr, mask) {
				mem[a] = i.value
			}
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func maskedAddresses(addr uint64, mask string) []uint64 {
	// Set any 1 bits to 1
	for i, c := range []rune(mask) {
		if c == '1' {
			bit := uint64(1) << (len(mask) - (i + 1))
			addr |= bit
		}
	}
	return maskedAddressesImpl(addr, mask)
}

func maskedAddressesImpl(addr uint64, mask string) []uint64 {
	x := strings.Index(mask, "X")
	if x == -1 {
		return []uint64{addr}
	}
	// Otherwise, try both values
	newMask := mask[0:x] + "0" + mask[x+1:]
	bit := uint64(1) << (len(mask) - (x + 1))
	zero := addr & ^bit
	one := addr | bit
	return append(maskedAddressesImpl(zero, newMask), maskedAddressesImpl(one, newMask)...)
}
