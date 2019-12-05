package aoc

// Max returns the smaller of two numbers
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// MaxInt returns the smaller of two numbers
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

// MinInt returns the smaller of two numberrs
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

// AbsInt returns the absolute value of a number
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

// SumInt returns the sum of a slice of integers
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
	d := []int64{}
	for i > int64(0) {
		d = append([]int64{i % 10}, d...)
		i = int64(i / 10)
	}

	return d
}

// DigitsInt splits the digits of a number into a slice
func DigitsInt(i int) []int {
	if i == 0 {
		return []int{i}
	}
	d := []int{}
	for i > 0 {
		d = append([]int{i % 10}, d...)
		i = int(i / 10)
	}

	return d
}
