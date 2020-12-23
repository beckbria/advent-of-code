package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/23
// A cup shuffle game

// Enable debug tracing
const debug = false
const debugIt = true

func main() {
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(364297581, 100))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(364297581))
	fmt.Println(sw.Elapsed())
}

func step1(input, iterations int) int {
	c := newCups(input)
	for i := 0; i < iterations; i++ {
		if debug {
			fmt.Printf("-- move %d --\n", i+1)
		}
		c.move()
	}
	return c.value()
}

func step2(input int) (int64, int64, int64) {
	c := newCupsPadded(input, 1000000)
	for i := 0; i < 10000000; i++ {
		if debug {
			fmt.Printf("-- move %d --\n", i+1)
		}
		c.move()
		if debugIt && i%10000 == 0 {
			fmt.Printf("-- move %d --\n", i)
		}
	}

	one := 0
	for ; c.c[one] != one; one++ {
	}
	first := int64(c.c[(one+1)%len(c.c)])
	second := int64(c.c[(one+2)%len(c.c)])
	return first, second, first * second
}

type cups struct {
	c       []int
	current int
	l       int // Length of the cups array.  Also the max number
}

func newCups(input int) *cups {
	var c []int
	for input > 0 {
		c = append([]int{input % 10}, c...)
		input /= 10
	}
	cu := cups{c: c, current: 0, l: len(c)}
	return &cu
}

func newCupsPadded(input int, dimension int) *cups {
	tmp := newCups(input)
	c := make([]int, dimension)
	for i, v := range tmp.c {
		c[i] = v
	}
	for i := len(tmp.c); i < dimension; i++ {
		c[i] = i + 1
	}
	cu := cups{c: c, current: 0, l: dimension}
	return &cu
}

func (c *cups) String() string {
	var b strings.Builder
	for i, v := range c.c {
		if i == c.current {
			fmt.Fprintf(&b, "(%d) ", v)

		} else {
			fmt.Fprintf(&b, "%d ", v)
		}
	}
	return b.String()
}

func (c *cups) value() int {
	v := 0
	one := 0
	for c.c[one] != 1 {
		one++
	}

	for i := one + 1; i < c.l; i++ {
		v *= 10
		v += c.c[i]
	}

	for i := 0; i < one; i++ {
		v *= 10
		v += c.c[i]
	}
	return v
}

func (c *cups) move() {
	if debug {
		fmt.Printf("cups: %s\n", c)
	}
	result := make([]int, len(c.c))
	pickup := make([]int, 3)

	result[c.current] = c.c[c.current]
	cIt := (c.current + 1) % c.l
	resultIt := cIt
	for i := 0; i < 3; i++ {
		pickup[i] = c.c[cIt]
		cIt = (cIt + 1) % c.l
	}
	if debug {
		fmt.Print("pick up: ")
		for i, v := range pickup {
			fmt.Print(v)
			if i < 2 {
				fmt.Print(", ")
			}
		}
		fmt.Println("")
	}

	dest := c.c[c.current] - 1
	if dest < 1 {
		dest += c.l
	}
	for foundDest := true; foundDest; {
		foundDest = false
		for _, v := range pickup {
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
	if debug {
		fmt.Printf("destination: %d\n\n", dest)
	}

	// Copy values until we get to the destination
	for foundDest := false; !foundDest; {
		result[resultIt] = c.c[cIt]
		foundDest = dest == c.c[cIt]
		cIt = (cIt + 1) % c.l
		resultIt = (resultIt + 1) % c.l
	}

	// Copy the picked up values
	for i := 0; i < 3; i++ {
		result[resultIt] = pickup[i]
		resultIt = (resultIt + 1) % c.l
	}

	// Copy any remaining falues
	for resultIt != c.current {
		result[resultIt] = c.c[cIt]
		cIt = (cIt + 1) % c.l
		resultIt = (resultIt + 1) % c.l
	}

	// Move the current node
	c.c = result
	c.current = (c.current + 1) % c.l
}
