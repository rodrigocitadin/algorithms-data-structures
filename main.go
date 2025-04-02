package main

import (
	"fmt"
	"rodrigocitadin/algorithms-data-structures/algorithms/arrays"
	"strings"
)

func main() {
        text := strings.Split("love", "")
	fmt.Println(arrays.TwoPointer(text))
}
