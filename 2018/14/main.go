package main

import (
	"fmt"
	"log"
	"time"
)

const debug = false

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Digits splits the digits of a number into a slice
func digits(i int) []int {
	if i == 0 {
		return []int{0}
	}
	d := make([]int, 0)
	for i > 0 {
		d = append(d, i%10)
		i = int(i / 10)
	}

	// Reverse the order
	for j := (len(d) / 2) - 1; j >= 0; j-- {
		k := len(d) - (j + 1)
		d[j], d[k] = d[k], d[j]
	}

	return d
}

// NextTenRecipes gives the next 10 recipes posted after a certain number of preceding recipes
func NextTenRecipes(count int) []int {
	r := recipes(count + 10)
	return r[len(r)-10:]
}

func recipes(count int) []int {
	r := []int{3, 7}
	elfA := 0
	elfB := 1
	if debug {
		printRecipes(r, elfA, elfB)
	}
	for len(r) < count {
		r = append(r, digits(r[elfA]+r[elfB])...)
		elfA = (elfA + r[elfA] + 1) % len(r)
		elfB = (elfB + r[elfB] + 1) % len(r)
		if debug {
			printRecipes(r, elfA, elfB)
		}
	}
	return r
}

func printRecipes(r []int, a, b int) {
	for i := 0; i < len(r); i++ {
		if i == a {
			fmt.Printf("(%d)", r[i])
		} else if i == b {
			fmt.Printf("[%d]", r[i])
		} else {
			fmt.Printf(" %d ", r[i])
		}
	}
	fmt.Printf("\n")
}

// ScoreIndex reports how many recipes appear on the board before the first instance of the provided pattern
func ScoreIndex(scoreDigits []int) int {
	r := []int{3, 7}
	elfA := 0
	elfB := 1

	// Note the location of any digit that matches the first digit from our search.
	// We can check for an exact match if there are enough recipes for the entire length.
	searchStarts := make([]int, 0)
	digitsLen := len(scoreDigits)

	for {
	SEARCH:
		for i, s := range searchStarts {
			if s+digitsLen <= len(r) {
				for j := 0; j < digitsLen; j++ {
					if scoreDigits[j] != r[s+j] {
						continue SEARCH
					}
				}
				// If all digits matched, we've found an instance of the pattern.
				return s
			} else {
				// Trim off all digits before this one since they were searched.
				searchStarts = searchStarts[i:]
				break
			}
		}

		d := digits(r[elfA] + r[elfB])
		for i := 0; i < len(d); i++ {
			if d[i] == scoreDigits[0] {
				searchStarts = append(searchStarts, len(r)+i)
			}
		}
		r = append(r, d...)
		elfA = (elfA + r[elfA] + 1) % len(r)
		elfB = (elfB + r[elfB] + 1) % len(r)
		if debug {
			printRecipes(r, elfA, elfB)
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println(NextTenRecipes(209231))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(ScoreIndex([]int{2, 0, 9, 2, 3, 1}))
	fmt.Println(time.Since(start))
}
