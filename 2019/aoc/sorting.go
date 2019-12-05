package aoc

// Go quirks: sort.Sort doesn't understand byte arrays, and you can't implement
// Len/Less/Swap on []byte (or []int64, or.... you get the point) directly.  So
// we need to provide manual comparator functions for an inbuilt type.  I would
// say something like "who thought that was a good idea", but the much likelier
// answer is that no one thought of it at all.  It's not like sorting is common
// in computer programs or anything...
// 
// Anyway, we have two choices: One, every single time we want to sort, write a
// custom lambda comparison function.  We cannot even write a reusable one here
// because it has to capture the slice in question:
//
// ints := []int64{1,5,2,-3}
// sort.Slice(ints, func(i,j int) bool { return ints[i] < ints[j]})
//
// The closest we get to a reusable way is to implement a custom slice type for
// sorting purposes, cast a native slice to our alias type (which is completely
// identical), and sort it that way.

// ByteSlice is a []byte alias necessary for sorting
type ByteSlice []byte

func (a ByteSlice) Len() int           { return len(a) }
func (a ByteSlice) Less(i, j int) bool { return a[i] < a[j] }
func (a ByteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Int64Slice is a []int64 alias necessary for sorting
type Int64Slice []int64

func (a Int64Slice) Len() int           { return len(a) }
func (a Int64Slice) Less(i, j int) bool { return a[i] < a[j] }
func (a Int64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }