package main

import (
	"fmt"

	"../aoc"
)

func main() {
	sw := aoc.NewStopwatch()
	start := 145852
	end := 616942
	fmt.Println(len(validPasswords(start, end, false)))
	fmt.Println(len(validPasswords(start, end, true)))
	fmt.Println(sw.Elapsed())
}

func validPasswords(start, end int, part2 bool) []int {
	valid := []int{}
	for pw := start; pw <= end; pw++ {
		if validPassword(pw, part2) {
			valid = append(valid, pw)
		}
	}
	return valid
}

func validPassword(pw int, part2 bool) bool {
	d := aoc.DigitsInt(pw)
	if d[0] > d[1] || d[1] > d[2] || d[2] > d[3] || d[3] > d[4] || d[4] > d[5] {
		return false
	}

	// Fun fact: If you replace == with != in the below clauses, you end up with a correct answer for one of the other inputs on part 2
	ab := d[0] == d[1]
	bc := d[1] == d[2]
	cd := d[2] == d[3]
	de := d[3] == d[4]
	ef := d[4] == d[5]

	if !ab && !bc && !cd && !de && !ef {
		return false
	}

	if (part2) {
		if !((ab && !bc) || (!ab && bc && !cd) || (!bc && cd && !de) || (!cd && de && !ef) || (!de && ef)) {
			return false
		}
	}

	return true
}