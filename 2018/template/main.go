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
	regEx = regexp.MustCompile("^(\\d+) (\\d+)$")
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadLine parses a line from the input file
func ReadLine(input string) {
	tokens := regEx.FindStringSubmatch(input)
	id, err := strconv.ParseInt(tokens[1], 10, 64)
	check(err)
	x, err := strconv.ParseInt(tokens[2], 10, 64)
	check(err)
}

func ReadLines(input []string) []interface{
	var objs []interface
	for _, s := range input {
		objs = append(objs, ReadLine(s))
	}
	return obj
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
