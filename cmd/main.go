package main

import (
	"fmt"
	"project1/internal/tree"
)

func main() {
	fmt.Println(tree.GenerateTrees(3))
}

func String[T any](str ...T) {
	fmt.Println(str)
}
