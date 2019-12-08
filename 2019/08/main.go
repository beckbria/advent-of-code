package main

import (
	"fmt"

	"../aoc"
)

// https://adventofcode.com/2019/day/8
// Read a 25x6 image layered image.  Do some statistics
// and then render the image, taking transparency into account

const (
	zero        = '0'
	one         = '1'
	two         = '2'
	black       = zero
	white       = one
	transparent = two
)

type layer struct {
	grid [][]rune
}

func newLayer(height, width int) layer {
	l := layer{grid: make([][]rune, height)}
	for i := 0; i < height; i++ {
		l.grid[i] = make([]rune, width)
	}
	return l
}

func readLayers(input string, height, width int) []layer {
	layerSize := height * width
	chars := []rune(input)
	layers := []layer{}

	for i := 0; i < len(chars)/layerSize; i++ {
		layers = append(layers, newLayer(height, width))
	}

	for i, c := range chars {
		l := i / layerSize
		layerPos := i % layerSize
		row := layerPos / width
		col := layerPos % width
		layers[l].grid[row][col] = c
	}
	return layers
}

func main() {
	input := aoc.ReadFileLines("input.txt")[0]
	sw := aoc.NewStopwatch()
	layers := readLayers(input, 6, 25)
	fmt.Println(bestLayerScore(layers))
	image := composeImage(layers)
	printImage(image)
	fmt.Println(sw.Elapsed())
}

func composeImage(layers []layer) layer {
	image := layers[0]
	for y := 0; y < len(image.grid); y++ {
		for x := 0; x < len(image.grid[0]); x++ {
			image.grid[y][x] = transparent
			for _, l := range layers {
				if l.grid[y][x] != transparent {
					image.grid[y][x] = l.grid[y][x]
					break
				}
			}
		}
	}

	return image
}

func printImage(image layer) {
	for _, row := range image.grid {
		for _, c := range row {
			switch c {
			case white:
				fmt.Print("X")
			case black:
				fmt.Print(" ")
			default:
				fmt.Print("?")
			}
		}
		fmt.Print("\n")
	}
}

func countValues(l layer) (int, int, int) {
	z := 0
	o := 0
	t := 0
	for _, row := range l.grid {
		for _, c := range row {
			switch c {
			case zero:
				z++
			case one:
				o++
			case two:
				t++
			}
		}
	}
	return z, o, t
}

func bestLayerScore(layers []layer) int {
	minZeroes := len(layers[0].grid)*len(layers[0].grid[0]) + 1
	score := 0
	for _, l := range layers {
		z, o, t := countValues(l)
		if z < minZeroes {
			minZeroes = z
			score = o * t
		}
	}
	return score
}
