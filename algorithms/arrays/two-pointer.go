package arrays

func TwoPointer[T any](array []T) []T {
	if len(array) <= 1 {
		return array
	}

	l := 0
	r := len(array) - 1

	for l <= r {
		hold := array[l]
		array[l] = array[r]
		array[r] = hold

		l += 1
		r -= 1
	}

	return array
}
