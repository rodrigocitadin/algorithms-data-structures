package arrays

func TwoPointer(text string) string {
	runes := []rune(text)
	if len(runes) <= 1 {
		return text
	}

	l := 0
	r := len(runes) - 1

	for l <= r {
		hold := runes[l]
		runes[l] = runes[r]
		runes[r] = hold

		l += 1
		r -= 1
	}

	return string(runes)
}
