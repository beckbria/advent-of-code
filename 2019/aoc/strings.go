package aoc

import (
	"log"
)

// IsUpper returns true if a letter is an upper-case ASCII letter
func IsUpper(i byte) bool {
	return (i >= 'A' && i <= 'Z')
}

// IsLower returns true if a letter is a lower-case ASCII letter
func IsLower(i byte) bool {
	return (i >= 'a' && i <= 'z')
}

// Ordinal returns the ordinal (0-25) value of an ASCII letter
func Ordinal(i byte) byte {
	if IsUpper(i) {
		return i - 'A'
	} else if IsLower(i) {
		return i - 'a'
	} else {
		log.Fatalf("Non-letter passed to Ordinal: '%c'")
		return 0
	}
}

// ToUpper converts a character to upper case
func ToUpper(i byte) byte {
	if IsLower(i) {
		return i - 'a' + 'A'
	}
	return i
}

// ToLower converts a character to lower case
func ToLower(i byte) byte {
	if IsUpper(i) {
		return i - 'A' + 'a'
	}
	return i
}

// Apparently sort.Sort doesn't understand byte arrays, and you can't implement
// Len/Less/Swap on []byte directly.  Go REALLY needs generics.
type ByteSlice []byte

func (a ByteSlice) Len() int           { return len(a) }
func (a ByteSlice) Less(i, j int) bool { return a[i] < a[j] }
func (a ByteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// StringSetIntersection finds the intersection of two sets of strings

type StringSet map[string]bool

func StringSetIntersection(a, b StringSet) StringSet {
	seen := make(StringSet)
	for k := range a {
		if _, present := b[k]; present {
			seen[k] = true
		}
	}
	return seen
}