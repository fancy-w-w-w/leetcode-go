package main

import (
	"fmt"
	"project1/internal/wordplay"
)

func main() {
	fmt.Println(wordplay.IsPalindrome("A man, a plan, a canal: Panama"))
}

func String[T any](str ...T) {
	fmt.Println(str)
}
