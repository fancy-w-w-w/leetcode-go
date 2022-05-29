package main

import (
	"fmt"
	"project1/internal/wordplay"
)

func main() {
	fmt.Println(wordplay.LargestNumber([]int{10, 2, 9, 39, 17}))
}

func String[T any](str ...T) {
	fmt.Println(str)
}
