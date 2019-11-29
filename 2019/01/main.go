package main

import (
	"fmt"

	"../aoc"
)

func main() {
	sw := aoc.NewStopwatch()
	fmt.Println(sw.Elapsed())
}