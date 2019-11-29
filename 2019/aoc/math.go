package aoc

// Max returns the smaller of two numbers
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// Max returns the smaller of two numbers
func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of two numberrs
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Min returns the smaller of two numberrs
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Abs returns the absolute value of a number
func Abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// Abs returns the absolute value of a number
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Sum returns the sum of a slice of integers
func Sum(i []int64) int64 {
	s := int64(0)
	for _, v := range i {
		s += v
	}
	return s
}

// Sum returns the sum of a slice of integers
func SumInt(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

// Digits splits the digits of a number into a slice
func Digits(i int64) []int64 {
	if i == int64(0) {
		return []int64{i}
	}
	d := make([]int64, 0)
	for i > int64(0) {
		d = append(d, i%10)
		i = int64(i / 10)
	}

	// Reverse the order
	for j := (len(d) / 2) - 1; j >= 0; j-- {
		k := len(d) - (j + 1)
		d[j], d[k] = d[k], d[j]
	}

	return d
}

// Digits splits the digits of a number into a slice
func DigitsInt(i int) []int {
	if i == 0 {
		return []int{i}
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