package main

import (
	"fmt"
)

func main() {
	fmt.Println("You are learning for loops")
	fmt.Println("We are looping 10 times from 0-9")
	for x := 0; x < 10; x++ {
		fmt.Println("X value is", x)
	}
}
