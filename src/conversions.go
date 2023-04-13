package main

import (
	"fmt"
)

func main() {
	var length int = 4
	var width float64 = 1.7

	//converting int into float
	fmt.Println(width * float64(length))

	fmt.Println(length * int(width))
}
