package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Reduce removes any adjacent pairs of upper/lowercase letters (Aa/aA) until no such pairs remain
func Reduce(polymer string) string {
	// By moving right-to-left, we can find all opposite pairs in a single pass without recursing
	// or looking back at our work.  Take the polymer abcCBA.  We walk through it until i=2:
	// abcCBA
	//   ^
	// After removing the matching pair and decrementing i, we're now looking at the following:
	// abBA
	//  ^
	// At every point, we've guaranteed that everything to the right of i can't be reduced any further.
	// The only thing that could be reduced is the leftmost character of the right chunk - which we're
	// comparing to the character to its left presently.  Thus, we can do the reduction in one pass.
	for i := len(polymer) - 2; i >= 0; i-- {
		if (len(polymer) > (i + 1)) && shouldCancel(polymer[i], polymer[i+1]) {
			polymer = polymer[:i] + polymer[i+2:]
		}
	}
	return polymer
}

func shouldCancel(i byte, j byte) bool {
	return ((isUpper(i) && isLower(j)) || (isLower(i) && isUpper(j))) && (ordinal(i) == ordinal(j))
}

func isUpper(i byte) bool {
	return (i >= 'A' && i <= 'Z')
}

func isLower(i byte) bool {
	return (i >= 'a' && i <= 'z')
}

func ordinal(i byte) byte {
	if isUpper(i) {
		return i - 'A'
	} else {
		return i - 'a'
	}
}

func toUpper(i byte) byte {
	if isLower(i) {
		return i - 'a' + 'A'
	}
	return i
}

func toLower(i byte) byte {
	if isUpper(i) {
		return i - 'A' + 'a'
	}
	return i
}

func uniqueUpperChars(s string) map[byte]bool {
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		seen[toUpper(s[i])] = true
	}
	return seen
}

// ShortestRemoveOne tests all combinations of the polymer with one letter pair removed (Aa, for instance).  It returns
// the shortest polymer which could be created if you reduced the resulting redacted polymer.
func ShortestRemoveOne(input string) string {
	seen := uniqueUpperChars(input)
	shortest := input
	for c := range seen {
		toReplace := fmt.Sprintf("%c", c)
		test := strings.Replace(input, toReplace, "", -1)
		test = strings.Replace(test, strings.ToLower(toReplace), "", -1)
		r := Reduce(test)
		if len(r) < len(shortest) {
			shortest = r
		}
	}
	return shortest
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
	r := Reduce(input[0])
	fmt.Println(len(r))
	fmt.Println(time.Since(start))
	start = time.Now()
	r = ShortestRemoveOne(input[0])
	fmt.Println(len(r))
	fmt.Println(time.Since(start))
}
