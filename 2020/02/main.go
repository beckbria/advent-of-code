package main

import (
	"fmt"
	"regexp"
	"strconv"

	"../../aoc"
)

// https://adventofcode.com/2020/day/2
// Validate passwords against a schema

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(part1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(part2(lines))
	fmt.Println(sw.Elapsed())
}

type restriction struct {
	char byte
	min  int64
	max  int64
}

type password struct {
	r  restriction
	pw string
}

func (p *password) valid() bool {
	count := int64(0)
	for i := 0; i < len(p.pw); i++ {
		if p.pw[i] == p.r.char {
			count++
		}
	}
	return count >= p.r.min && count <= p.r.max
}

func (p *password) valid2() bool {
	first := p.r.min <= int64(len(p.pw)) && p.pw[p.r.min-1] == p.r.char
	second := p.r.max <= int64(len(p.pw)) && p.pw[p.r.max-1] == p.r.char
	return first != second
}

func part1(lines []string) int64 {
	pws := parsePasswords(lines)
	valid := int64(0)
	for _, p := range pws {
		if p.valid() {
			valid++
		}
	}
	return valid
}

func part2(lines []string) int64 {
	pws := parsePasswords(lines)
	valid := int64(0)
	for _, p := range pws {
		if p.valid2() {
			valid++
		}
	}
	return valid
}

var (
	inputRegEx = regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
)

func parsePasswords(lines []string) []password {
	pws := make([]password, 0)
	for _, l := range lines {
		tokens := inputRegEx.FindStringSubmatch(l)
		var pw password
		pw.pw = tokens[4]
		pw.r.char = tokens[3][0]
		pw.r.min, _ = strconv.ParseInt(tokens[1], 10, 64)
		pw.r.max, _ = strconv.ParseInt(tokens[2], 10, 64)
		pws = append(pws, pw)
	}
	return pws
}
