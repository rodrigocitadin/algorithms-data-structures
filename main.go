package main

import (
	"fmt"
	"rodrigocitadin/algorithms-data-structures/algorithms"
	"strings"
)

func main() {
	text := strings.Split("love", "")
	fmt.Println(algorithms.TwoPointer(text))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        fmt.Println(algorithms.BinarySearch(numbers, 10))
}
