package main

import (
	"fmt"
	binarysearch "project1/internal/binary_search"
)

func main() {
	fmt.Println(binarysearch.SearchNumV2([]int{1, 0, 1, 1, 1}, 0))
}

func String[T any](str ...T) {
	fmt.Println(str)
}
