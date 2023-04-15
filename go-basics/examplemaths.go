package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("learning go for first time"))
	//a string is always in "" whereas the runes(char) are always in ''.
	// if I run the line below instead of printing the A it will print the unicode character code.
	fmt.Println('A')
}
