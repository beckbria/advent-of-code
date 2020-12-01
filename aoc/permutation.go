package aoc

// NextPermutation sets a slice to its next lexographic permutation.  Returns true if it found another permutation
func NextPermutation(arr []int64) bool {
	// Find longest non-increasing suffix
	i := len(arr) - 1
	for ; i > 0 && arr[i-1] > arr[i]; i-- {
	}

	if i <= 0 {
		// We've found the last permutation
		return false
	}

	// Find the rightmost element that exceeds the pivot.  Guaranteed to find something because arr[i+1] > arr[i]
	j := len(arr) - 1
	for ; arr[j] <= arr[i-1]; j-- {
	}

	// Swap the pivot with j
	arr[i-1], arr[j] = arr[j], arr[i-1]

	// Reverse the suffix
	for j = len(arr) - 1; i < j; j-- {
		arr[i], arr[j] = arr[j], arr[i]
		i++
	}

	return true
}
