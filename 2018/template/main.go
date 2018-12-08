package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	// Input format: "# #"
	fabricRegEx = regexp.MustCompile("^(\\d+) (\\d+)$")
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadFabric converts a line from the input file into a Fabric object
func ReadNumbers(input string) {
	tokens := fabricRegEx.FindStringSubmatch(input)
	id, err := strconv.ParseInt(tokens[1], 10, 64)
	check(err)
	x, err := strconv.ParseInt(tokens[2], 10, 64)
	check(err)
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	fmt.Println(input)
	fmt.Println(time.Since(start))
}
