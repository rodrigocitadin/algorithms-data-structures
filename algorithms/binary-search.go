package algorithms

func BinarySearch(array []int, target int) int {
	l := 0
	r := len(array)

	for l < r {
		m := (l + r) / 2

		switch {
		case array[m] == target:
			return m
		case array[m] < target:
			l = m + 1
		default:
			r = m
		}
	}

	return -1
}
