package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/18
// It's Math, Jim, but not as we know it

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	fmt.Println(step1(lines))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(step2(lines))
	fmt.Println(sw.Elapsed())
}

type operation rune

const (
	multiply = '*'
	add      = '+'
	lParen   = '('
	rParen   = ')'
	literal  = '#'
)

type token struct {
	op  operation
	val int64
}

func (t *token) print() {
	switch t.op {
	case multiply, add, lParen, rParen:
		fmt.Print(string(t.op))
	case literal:
		fmt.Print(t.val)
	}
}

func newToken(op operation) token {
	return token{op: op, val: 0}
}

func newLiteral(val int64) token {
	return token{op: literal, val: val}
}

type expression []token

func (e expression) print() {
	for _, t := range e {
		t.print()
		fmt.Print(" ")
	}
	fmt.Println("")
}

func (e expression) evaluate(advanced bool) int64 {
	// Step 1: Condense any parenthesis
	for foundParen := true; foundParen; {
		foundParen = false
		start := -1
		end := -1
		depth := 0
		for i, c := range e {
			switch c.op {
			case lParen:
				if start < 0 {
					start = i
				}
				foundParen = true
				depth++
			case rParen:
				if start < 0 {
					log.Fatalf("Unmatched rParen at %d", i)
				}
				depth--
				if depth == 0 {
					end = i
				}
			}
			if end >= 0 {
				break
			}
		}
		if foundParen {
			val := e[start+1 : end].evaluate(advanced)
			begin := expression{}
			if start > 0 {
				begin = e[:start]
			}
			finish := expression{}
			if end < (len(e) - 1) {
				finish = e[end+1:]
			}
			e = append(begin, newLiteral(val))
			e = append(e, finish...)
		}
	}

	// Now there should be no parenthesis left
	if e[0].op != literal {
		log.Fatalf("Invalid post-paren first op: %s", string(e[0].op))
	}

	if advanced {
		// Do addition first
		for foundAdd := true; foundAdd; {
			foundAdd = false
			for i, c := range e {
				switch c.op {
				case add:
					foundAdd = true
					val := e[i-1].val + e[i+1].val
					start := expression{}
					if i > 1 {
						start = e[:i-1]
					}
					finish := expression{}
					if i < (len(e) - 2) {
						finish = e[i+2:]
					}
					e = append(start, newLiteral(val))
					e = append(e, finish...)
				}
				if foundAdd {
					break
				}
			}
		}
	}

	lastOp := operation(literal)
	result := e[0].val

	for i := 1; i < len(e); i++ {
		switch e[i].op {
		case literal:
			switch lastOp {
			case multiply:
				result *= e[i].val
			case add:
				result += e[i].val
			default:
				log.Fatalf("Invalid op before literal: %s", string(lastOp))
			}
		case multiply, add:
			if lastOp != literal {
				log.Fatalf("Invalid op before math: %s", string(lastOp))
			}
		default:
			log.Fatalf("Unexpected postparen op: %s", string(e[i].op))
		}
		lastOp = e[i].op
	}

	return result
}

func tokenizeSlice(lines []string) []expression {
	var t []expression
	for _, l := range lines {
		t = append(t, tokenize(l))
	}
	return t
}

func tokenize(s string) expression {
	s = strings.ReplaceAll(s, " ", "")
	start := -1
	var tokens []token
	for i, c := range []rune(s) {
		switch c {
		case multiply, add, lParen, rParen:
			if start >= 0 {
				val, _ := strconv.ParseInt(s[start:i], 10, 64)
				tokens = append(tokens, newLiteral(val))
			}
			tokens = append(tokens, newToken(operation(c)))
			start = -1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if start < 0 {
				start = i
			}
		default:
			log.Fatalf("Unexpected character: %s", string(c))
		}
	}
	if start > 0 {
		val, _ := strconv.ParseInt(s[start:], 10, 64)
		tokens = append(tokens, newLiteral(val))
	}
	return tokens
}

func step1(lines []string) int64 {
	exp := tokenizeSlice(lines)
	sum := int64(0)
	for _, e := range exp {
		sum += e.evaluate(false)
	}
	return sum
}

func step2(lines []string) int64 {
	exp := tokenizeSlice(lines)
	sum := int64(0)
	for _, e := range exp {
		sum += e.evaluate(true)
	}
	return sum
}
