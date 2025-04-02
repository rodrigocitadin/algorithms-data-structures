package algorithms

import "math"

func ExponentialSearch(arr []int, target int) int {
	if arr[0] == target {
		return 0
	}

	n := len(arr)
	i := 1

	for i < n && arr[i] < target {
		i *= 2
	}

	if arr[i] == target {
		return i
	}

	return BinarySearchLogic(
		arr,
		target,
		i/2,
		int(math.Min(float64(i), float64(n-1))))
}
