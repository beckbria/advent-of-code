package aoc

import "math"

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

// Gcd uses Euclid's algorithm to find the greatest common denominator of two numbers
func Gcd(x, y int64) int64 {
	if x == 0 || y == 0 {
		return 1
	}
	x = Abs(x)
	y = Abs(y)
	// Ensure that x is the greater number
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		oldY := y
		y = x % y
		x = oldY
	}
	return x
}

// Fraction represents a fraction with integer numerator and denominator
type Fraction struct {
	Numerator, Denominator int64
}

// NewFraction creates a new reduced fraction
func NewFraction(num, den int64) Fraction {
	f := Fraction{Numerator: num, Denominator: den}
	f.Reduce()
	return f
}

// Reduce reduces a fraction to its lowest terms
func (f *Fraction) Reduce() {
	gcd := Gcd(f.Numerator, f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd
}

// Equals returns true if two fractions are equal
func (f *Fraction) Equals(g *Fraction) bool {
	if f.Numerator == g.Numerator {
		return f.Numerator == 0 || f.Denominator == g.Denominator
	}
	return false
}

// PiOver2 is a constant equal to Pi/2
const PiOver2 = float64(math.Pi / 2)
