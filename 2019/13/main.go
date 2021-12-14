package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

func main() {
	program := intcode.ReadIntCode("input.txt")

	sw := lib.NewStopwatch()
	// Part 1
	fmt.Println(countBlocks(program))
	fmt.Println(sw.Elapsed())
	sw.Reset()

	// Part 2
	fmt.Println(finalGameScore(program))
	fmt.Println(sw.Elapsed())
}

type tile int64

const (
	empty  tile = 0
	wall   tile = 1
	block  tile = 2
	paddle tile = 3
	ball   tile = 4
)

func countBlocks(p intcode.Program) int64 {
	c := intcode.NewComputer(p)
	io := intcode.NewConstantInputOutput(0)
	c.Io = io
	c.Run()
	screen := make(map[lib.Point]tile)
	for i := 0; i < len(io.Outputs); i += 3 {
		screen[lib.Point{X: io.Outputs[i], Y: io.Outputs[i+1]}] = tile(io.Outputs[i+2])
	}
	count := int64(0)
	for _, v := range screen {
		if v == block {
			count++
		}
	}
	return count
}

func finalGameScore(p intcode.Program) int64 {
	// Fire up the game
	c := intcode.NewComputer(p)
	// Insert tokens
	c.Memory[0] = 2
	io := newGameIo([]int64{})
	c.Io = io
	// Play to completion
	c.Run()
	// Render the screen one more time
	io.updateGameState()

	return io.score
}

type gameIo struct {
	inputs         []int64
	nextInputIndex int
	Outputs        []int64
	screen         map[lib.Point]tile
	score          int64
}

const (
	left        = -1
	stay        = 0
	right       = 1
	manualInput = false
)

func (io *gameIo) updateGameState() {
	// Display the screen
	for i := 0; i < len(io.Outputs); i += 3 {
		x := io.Outputs[i]
		y := io.Outputs[i+1]
		if x == -1 && y == 0 {
			io.score = io.Outputs[i+2]
		} else {
			io.screen[lib.Point{X: x, Y: y}] = tile(io.Outputs[i+2])
		}
	}
}

// GetInput returns the provided input value
func (io *gameIo) GetInput() int64 {
	// Pre-run the game
	if io.nextInputIndex < len(io.inputs) {
		i := io.inputs[io.nextInputIndex]
		io.nextInputIndex++
		return i
	}

	io.updateGameState()
	if manualInput {
		io.drawScreen()
	}
	io.ResetOutputs()
	if manualInput {
		reader := bufio.NewReader(os.Stdin)
		for true {
			input, _ := reader.ReadString('\n')
			switch []rune(input)[0] {
			case 'a', 'A':
				io.inputs = append(io.inputs, left)
				return left
			case 's', 'S':
				io.inputs = append(io.inputs, stay)
				return stay
			case 'd', 'D':
				io.inputs = append(io.inputs, right)
				return right
			case 'q', 'Q':
				fmt.Println(io.inputs)
			}
		}
		return 0
	} else {
		ballX := int64(0)
		paddleX := int64(0)
		for pt, t := range io.screen {
			if t == ball {
				ballX = pt.X
			} else if t == paddle {
				paddleX = pt.X
			}
		}
		if paddleX < ballX {
			return right
		} else if paddleX > ballX {
			return left
		} else {
			return stay
		}
	}
}

func (io *gameIo) drawScreen() {
	minX := int64(9999)
	minY := int64(9999)
	maxX := int64(-9999)
	maxY := int64(-9999)
	for pt := range io.screen {
		minX = lib.Min(minX, pt.X)
		maxX = lib.Max(maxX, pt.X)
		minY = lib.Min(minY, pt.Y)
		maxY = lib.Max(maxY, pt.Y)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			t := io.screen[lib.Point{X: x, Y: y}]
			switch t {
			case empty:
				fmt.Print(" ")
			case wall:
				fmt.Print("#")
			case block:
				fmt.Print("-")
			case paddle:
				fmt.Print("=")
			case ball:
				fmt.Print("O")
			default:
				fmt.Print("?")
			}
		}
		fmt.Print("\n")
	}
}

// Output collects the output value into a slice for later use
func (io *gameIo) Output(o int64) {
	io.Outputs = append(io.Outputs, o)
}

func (io *gameIo) ResetOutputs() {
	io.Outputs = []int64{}
}

// Reset resets the output buffer
func (io *gameIo) Reset() {
	io.Outputs = []int64{}
	io.nextInputIndex = 0
}

// newGameIo creates an IO component which returns a fixed series of inputs
func newGameIo(initialInput []int64) *gameIo {
	io := gameIo{Outputs: []int64{}, inputs: initialInput, nextInputIndex: 0, screen: make(map[lib.Point]tile)}
	return &io
}
