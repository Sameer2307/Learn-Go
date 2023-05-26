package main

import (
	"fmt"
)

func main() {
	//use sprint to print
	a := 23
	name := "Sameer"
	// use %t to print value
	printed := fmt.Sprintf("%v\n %v\n", name, a)
	fmt.Printf(printed)
}
