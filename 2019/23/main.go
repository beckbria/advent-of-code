package main

import (
	"fmt"
	"sync"

	"github.com/beckbria/advent-of-code/2019/lib"
	"github.com/beckbria/advent-of-code/2019/intcode"
)

const debug = false

func main() {
	p := intcode.ReadIntCode("input.txt")

	sw := lib.NewStopwatch()
	// Part 1
	//fmt.Println(firstTo255(p))
	//fmt.Println(sw.Elapsed())

	// Part 2
	sw.Reset()
	fmt.Println(duplicateNat(p))
	fmt.Println(sw.Elapsed())
}

// networkIo is a non-blocking Intcode IO component for the networked computers
type networkIo struct {
	inputs         []int64
	nextInputIndex int
	Outputs        []int64
	inputMux       sync.Mutex
	outputMux      sync.Mutex
	rout           *router
}

// GetInput returns the provided input value
func (io *networkIo) GetInput() int64 {
	input := int64(-1)
	io.inputMux.Lock()
	if io.nextInputIndex < len(io.inputs) {
		input = io.inputs[io.nextInputIndex]
		io.nextInputIndex++
	}
	io.inputMux.Unlock()
	return input
}

// AppendInput adds another input to the queue
func (io *networkIo) AppendInput(i ...int64) {
	io.inputMux.Lock()
	io.inputs = append(io.inputs, i...)
	io.inputMux.Unlock()
}

// Output collects the output value into a slice for later use
func (io *networkIo) Output(o int64) {
	send := false
	dest := int64(-1)
	x := int64(-1)
	y := int64(-1)
	io.outputMux.Lock()
	io.Outputs = append(io.Outputs, o)
	lo := len(io.Outputs)
	if lo%3 == 0 {
		send = true
		dest = io.Outputs[lo-3]
		x = io.Outputs[lo-2]
		y = io.Outputs[lo-1]
	}
	io.outputMux.Unlock()
	if send {
		io.rout.send(dest, x, y)
	}
}

func (io *networkIo) Reset() {}

func (io *networkIo) idle() bool {
	//i := false
	//io.inputMux.Lock()
	return io.nextInputIndex >= len(io.inputs)
	//io.inputMux.Unlock()
	//return i
}

func newNetworkIo(r *router) *networkIo {
	io := networkIo{Outputs: []int64{}, inputs: []int64{}, nextInputIndex: 0, rout: r}
	return &io
}

type router struct {
	io      [50]*networkIo
	lost    map[int64][]lib.Point
	lostMux sync.Mutex
}

func (r *router) send(dest, x, y int64) {
	if dest < int64(len(r.io)) {
		r.io[dest].AppendInput(x, y)
	} else {
		r.lostMux.Lock()
		if _, found := r.lost[dest]; !found {
			r.lost[dest] = []lib.Point{}
		}
		r.lost[dest] = append(r.lost[dest], lib.Point{X: x, Y: y})
		r.lostMux.Unlock()
	}
}

func (r *router) hasLostMailFor(dest int64) bool {
	r.lostMux.Lock()
	_, found := r.lost[dest]
	r.lostMux.Unlock()
	return found
}

func (r *router) idle() bool {
	for _, i := range r.io {
		if !i.idle() {
			return false
		}
	}
	return true
}

func newRouter() *router {
	r := router{lost: make(map[int64][]lib.Point)}
	for i := range r.io {
		r.io[i] = newNetworkIo(&r)
		// Assign the initial network address
		r.io[i].AppendInput(int64(i))
	}
	return &r
}

func firstTo255(p intcode.Program) int64 {
	r := newRouter()

	// Create 50 computers
	c := make([]intcode.Computer, 50)
	for i := range c {
		c[i] = intcode.NewComputer(p)
		c[i].Io = r.io[i]
		go func(c intcode.Computer) {
			c.Run()
		}(c[i])
	}

	for !r.hasLostMailFor(255) {
	}
	return r.lost[255][0].Y
}

func duplicateNat(p intcode.Program) int64 {
	r := newRouter()

	// Create 50 computers
	c := make([]intcode.Computer, 50)
	for i := range c {
		c[i] = intcode.NewComputer(p)
		c[i].Io = r.io[i]
		go func(c intcode.Computer) {
			c.Run()
		}(c[i])
	}

	// NAT loop
	lastNat := lib.Point{X: -1, Y: -1423}
	for true {
		oldLen := len(r.lost[255])
		// Wait for idle
		for !r.idle() {
		}
		// Did we repeat the Y value?
		newLen := len(r.lost[255])
		if newLen > oldLen {
			nat := r.lost[255][newLen-1]
			if lastNat.Y == nat.Y {
				return lastNat.Y
			}
			// Otherwise resume the network
			r.send(0, nat.X, nat.Y)
		}
	}
	return -1
}
