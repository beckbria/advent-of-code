package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/beckbria/advent-of-code/2019/intcode"
)

const debug = false

func main() {
	p := intcode.ReadIntCode("input.txt")
	c := intcode.NewComputer(p)
	io := intcode.NewASCIIIo()
	c.Io = io

	// Loop forever, displaying output and reading console input
	reader := bufio.NewReader(os.Stdin)
	for true {
		c.RunToNextInput()
		out := io.OutputAsString()
		if len(out) > 0 {
			fmt.Println(out)
			io.ClearOutput()
		}

		if !c.IsRunning() {
			break
		}

		// Only take more input if the computer has processed all current input and shown output
		if len(out) > 0 {
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if debug {
				fmt.Printf("Read <<%s>>\n", text)
			}
			io.AppendInput(text)
		}
		c.Step()
	}
}
