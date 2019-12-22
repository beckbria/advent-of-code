package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"../aoc"
)

func main() {
	input := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	// Part 1
	d := newDeck(10007)
	inst := parseInstructions(input)
	d.shuffle(inst)
	fmt.Println(d.find(2019))
	fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()

	fmt.Println(sw.Elapsed())
}

type operation int

const (
	newStack      operation = 0
	dealIncrement operation = 1
	cut           operation = 2
)

type instruction struct {
	op    operation
	value int
}

type deck []int

func newDeck(size int) deck {
	d := make(deck, size)
	for i := range d {
		d[i] = i
	}
	return d
}

func (d deck) shuffle(inst []instruction) {
	for _, i := range inst {
		switch i.op {
		case newStack:
			d.reverse()
		case dealIncrement:
			d.increment(i.value)
		case cut:
			d.cut(i.value)
		default:
			log.Fatalf("Unexpected shuffle operation: %d\n", i.op)
		}
	}
}

func (d deck) reverse() {
	start := 0
	end := len(d) - 1
	for end > start {
		d[start], d[end] = d[end], d[start]
		start++
		end--
	}
}

func (d deck) increment(val int) {
	new := make(deck, len(d))
	to := 0
	for from := 0; from < len(d); from++ {
		if new[to] != 0 {
			log.Fatalf("Tried to write to index %d twice (previous value %d)\n", to, new[to])
		}
		new[to] = d[from]
		to = (to + val) % len(d)
	}
	copy(d, new)
}

func (d deck) cut(val int) {
	pivot := 0
	if val < 0 {
		pivot = len(d) + val
	} else {
		pivot = val
	}
	new := append(d[pivot:], d[0:pivot]...)
	copy(d, new)
}

func (d deck) find(needle int) int {
	for i, v := range d {
		if v == needle {
			return i
		}
	}
	return -1
}

var (
	cutRegEx       = regexp.MustCompile("^cut (-?[0-9]+)$")
	incrementRegEx = regexp.MustCompile("^deal with increment ([0-9]+)$")
	stackRegEx     = regexp.MustCompile("^deal into new stack$")
)

func parseInstructions(input []string) []instruction {
	inst := []instruction{}
	for _, s := range input {
		tokens := stackRegEx.FindStringSubmatch(s)
		if tokens != nil {
			inst = append(inst, instruction{op: newStack})
			continue
		}

		tokens = cutRegEx.FindStringSubmatch(s)
		if tokens != nil {
			val, err := strconv.Atoi(tokens[1])
			aoc.Check(err)
			inst = append(inst, instruction{op: cut, value: val})
			continue
		}

		tokens = incrementRegEx.FindStringSubmatch(s)
		if tokens != nil {
			val, err := strconv.Atoi(tokens[1])
			aoc.Check(err)
			inst = append(inst, instruction{op: dealIncrement, value: val})
			continue
		}

		log.Fatalf("Failed to parse instruction: \"%s\"\n", s)
	}
	return inst
}
