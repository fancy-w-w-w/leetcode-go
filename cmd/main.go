package main

import (
	"fmt"
	arrayhandle "project1/internal/array_handle"
)

func main() {
	fmt.Println(arrayhandle.MaxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func String[T any](str ...T) {
	fmt.Println(str)
}
