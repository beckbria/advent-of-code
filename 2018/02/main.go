package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Checksum computes a very basic checksum.  It counts the number of strings in the input
// which have a letter which appears exactly twice, and the number of strings in the input
// which have a leter which appears exactly three time.  It then multiplies these two counts
// to produce the checksum.
func Checksum(input []string) int {
	twice := 0
	thrice := 0
	for _, s := range input {
		counts := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			counts[s[i]]++
		}
		seenTwice := false
		seenThrice := false
		for _, c := range counts {
			if c == 2 {
				seenTwice = true
			} else if c == 3 {
				seenThrice = true
			}
		}
		if seenTwice {
			twice++
		}
		if seenThrice {
			thrice++
		}
	}
	return twice * thrice
}

// EditDistanceOne finds two strings in the input array which differ by only a single character
func EditDistanceOne(input []string) (string, string, error) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			d, err := differentCharacterCount(input[i], input[j])
			if err != nil {
				return "", "", err
			} else if d == 1 {
				return input[i], input[j], nil
			}
		}
	}
	return "", "", fmt.Errorf("No strings with edit distance one")
}

func differentCharacterCount(a string, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("String lengths differ: %s, %s", a, b)
	}
	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}

// DifferingCharacter returns the first character that is different between two strings
func CommonLetters(a string, b string) (string, error) {
	if len(a) != len(b) {
		return "", fmt.Errorf("String lengths differ: %s, %s", a, b)
	}

	c := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			c = c + string(a[i])
		}
	}

	return c, nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
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
	fmt.Printf("Checksum: %d\n", Checksum(input))
	a, b, err := EditDistanceOne(input)
	check(err)
	c, err := CommonLetters(a, b)
	check(err)
	fmt.Printf("Common Letters: %s\n", c)
}
