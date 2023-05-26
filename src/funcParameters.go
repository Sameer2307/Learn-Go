package main

import "fmt"

func repeat(line string, number int) {
	for i := 0; i < number; i++ {
		fmt.Println(line)
	}
}

func main() {
	repeat("I am repeated", 3)
}
