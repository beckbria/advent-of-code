package main

import (
	"fmt"

	"../aoc"
)

// https://adventofcode.com/2019/day/8
// Read a 25x6 image layered image.  Do some statistics
// and then render the image, taking transparency into account

const width = 25
const height = 6
const layerSize = width * height
const zero = '0'
const one = '1'
const two = '2'
const black = zero
const white = one
const transparent = two

type layer struct {
	grid [height][width]rune
}

func main() {
	input := aoc.ReadFileLines("input.txt")[0]
	sw := aoc.NewStopwatch()
	layers := readLayers(input)
	fmt.Println(bestLayerScore(layers))
	image := parseImage(layers)
	printImage(image)
	fmt.Println(sw.Elapsed())
}

func parseImage(layers []layer) layer {
	image := layers[0]
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
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

func readLayers(input string) []layer {
	chars := []rune(input)
	layers := make([]layer, len(chars)/layerSize)

	for i, c := range chars {
		l := i / layerSize
		layerPos := i % layerSize
		row := layerPos / width
		col := layerPos % width
		layers[l].grid[row][col] = c
	}
	return layers
}

func bestLayerScore(layers []layer) int {
	minZeroes := layerSize + 1
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
