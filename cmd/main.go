package main

import (
	"fmt"
)

func main() {

}

func String[T any](str ...T) {
	fmt.Println(str)
}
