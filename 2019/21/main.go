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
	p := intcode.ReadIntCode("2019/21/input.txt")
	c := intcode.NewComputer(p)
	io := intcode.NewASCIIIo()
	c.Io = io

	// Loop forever, displaying output and reading console input
	reader := bufio.NewReader(os.Stdin)
	for true {
		c.RunToNextInput()
		out := io.OutputAsString()

		if !c.IsRunning() {
			fmt.Println(io.Outputs[len(io.Outputs)-1])
			break
		}

		if len(out) > 0 {
			fmt.Println(out)
			io.ClearOutput()
		}

		// Only take more input if the computer has processed all current input and shown output
		if len(out) > 0 {
			fmt.Print("> ")
			inst := ""
			for inst != "WALK" && inst != "RUN" {
				text, _ := reader.ReadString('\n')
				inst = strings.Replace(text, "\n", "", -1)
				if debug {
					fmt.Printf("Read <<%s>>\n", inst)
				}
				io.AppendInput(inst)
			}
		}
		c.Step()
	}
}
