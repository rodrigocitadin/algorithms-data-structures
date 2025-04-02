package algorithms

import "math"

func ExponentialSearch(array []int, target int) int {
	if array[0] == target {
		return 0
	}

	n := len(array)
	i := 1

	for i < n && array[i] < target {
		i *= 2
	}

	if array[i] == target {
		return i
	}

	return BinarySearchLogic(
		array,
		target,
		i/2,
		int(math.Min(float64(i), float64(n-1))))
}
