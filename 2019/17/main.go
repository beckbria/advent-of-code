package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

const debug = false

func main() {
	p := intcode.ReadIntCode("input.txt")

	sw := lib.NewStopwatch()
	// Part 1
	fmt.Println(sw.Elapsed())
	fmt.Println(alignmentParams(p))

	// Part 2
	sw.Reset()
	fmt.Println(collectDust(p))
	fmt.Println(sw.Elapsed())
}

const (
	scaffold = '#'
	open     = '.'
	newLine  = int64(10)
)

type grid [][]rune

func (g grid) print() {
	for _, row := range g {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
}

func (g grid) alignmentParams() []int64 {
	ap := []int64{}
	for y := 1; y < (len(g) - 1); y++ {
		for x := 1; x < (len(g[y]) - 1); x++ {
			if g[y][x] == scaffold && g[y-1][x] == scaffold && g[y+1][x] == scaffold && g[y][x-1] == scaffold && g[y][x+1] == scaffold {
				ap = append(ap, int64(x*y))
			}
		}
	}
	return ap
}

func readGrid(output []int64) grid {
	g := grid{}

	startLine := true
	for _, char := range output {
		if startLine {
			g = append(g, []rune{})
			startLine = false
		}

		if char == newLine {
			startLine = true
		} else {
			g[len(g)-1] = append(g[len(g)-1], rune(char))
		}
	}

	// Filter empty rows
	gf := grid{}
	for _, row := range g {
		if len(row) > 0 {
			gf = append(gf, row)
		}
	}

	return gf
}

func alignmentParams(p intcode.Program) int64 {
	c := intcode.NewComputer(p)
	io := intcode.NewStreamInputOutput([]int64{})
	c.Io = io
	c.Run()
	g := readGrid(io.Outputs)
	ap := g.alignmentParams()
	sum := int64(0)
	for _, i := range ap {
		sum += i
	}
	return sum
}

func collectDust(p intcode.Program) int64 {
	c := intcode.NewComputer(p)
	c.Memory[0] = int64(2)
	// Manually compressed
	input := robotProgram([]string{
		"A,B,A,C,B,A,B,C,C,B", // Main program
		"L,12,L,12,R,4",       // A
		"R,10,R,6,R,4,R,4",    // B
		"R,6,L,12,L,12"})      // C

	io := intcode.NewStreamInputOutput(input)
	c.Io = io
	c.Run()
	return io.LastOutput()
}

func robotProgram(routines []string) []int64 {
	prog := []int64{}
	for _, r := range routines {
		for _, c := range []rune(r) {
			prog = append(prog, int64(c))
		}
		prog = append(prog, newLine)
	}
	// Reject visual mode
	prog = append(prog, int64('r'))
	prog = append(prog, newLine)
	return prog
}
