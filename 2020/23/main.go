package main

import (
	"fmt"

	"github.com/beckbria/advent-of-code/2020/lib"
)

// https://adventofcode.com/2020/day/23
// A cup shuffle game

func main() {
	sw := lib.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(364297581, 100))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	_, _, s2 := step2(364297581)
	fmt.Println(s2)
	fmt.Println(sw.Elapsed())
}

func step1(input, iterations int) int {
	digits := splitDigits(input)
	c := newCups(digits)
	for i := 0; i < iterations; i++ {
		c.move()
	}
	return c.value()
}

func step2(input int) (int64, int64, int64) {
	dimension := 1000000
	iterations := 10000000
	digits := splitDigits(input)
	padded := make([]int, dimension)
	copy(padded, digits)
	for i := len(digits); i < dimension; i++ {
		padded[i] = i + 1
	}
	c := newCups(padded)
	for i := 0; i < iterations; i++ {
		c.move()
	}

	one := c.nodes[1]
	first := int64(one.next.val)
	second := int64(one.next.next.val)
	return first, second, first * second
}

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	n := node{val: val, next: nil}
	return &n
}

type cups struct {
	head    *node
	current *node
	nodes   map[int]*node
	l       int // Length of the cups array.  Also the max number
}

func splitDigits(input int) []int {
	var val []int
	for input > 0 {
		val = append([]int{input % 10}, val...)
		input /= 10
	}
	return val
}

func newCups(input []int) *cups {
	// Create the first node
	val := input[0]
	prev := newNode(val)
	c := cups{
		head:    prev,
		current: prev,
		nodes:   make(map[int]*node),
		l:       len(input)}
	c.nodes[val] = prev

	for i := 1; i < c.l; i++ {
		val = input[i]
		curr := newNode(val)
		prev.next = curr
		c.nodes[val] = curr
		prev = curr
	}
	// Complete the circle
	c.nodes[input[c.l-1]].next = c.head

	return &c
}

func (c *cups) value() int {
	v := 0

	one := c.nodes[1]
	for curr := one.next; curr != one; curr = curr.next {
		v *= 10
		v += curr.val
	}

	return v
}

func (c *cups) move() {
	pickupHead := c.current.next
	pickupTail := pickupHead.next.next
	c.current.next = pickupTail.next

	puVal := [3]int{pickupHead.val, pickupHead.next.val, pickupTail.val}

	dest := c.current.val - 1
	if dest < 1 {
		dest += c.l
	}
	for foundDest := true; foundDest; {
		foundDest = false
		for _, v := range puVal {
			if v == dest {
				foundDest = true
				break
			}
		}
		if foundDest {
			dest--
			if dest < 1 {
				dest += c.l
			}
		}
	}

	// Insert the pickup after the destination
	d := c.nodes[dest]
	pickupTail.next = d.next
	d.next = pickupHead

	// Move the current node
	c.current = c.current.next
}
