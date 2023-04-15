package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter a number")
	// Setup a buffer reader to get a value
	reader := bufio.NewReader(os.Stdin) // now our reader becomes like a function
	// input, err := reader.ReadString('\n') // \n means everything until new line will be taken, Go returns 2 values most of time

	// cant just declare a variable err and not use it sincce go doesnt allow it
	// so to solve this we use '_', hence line 13 becomes

	// '_' means blank identifier and works as a placeholder
	input, _ := reader.ReadString('\n')
	fmt.Println("The input is", input)

	// The other way is to handle the error
	fmt.Println("Enter a number")
	reader2 := bufio.NewReader(os.Stdin)
	input2, error := reader2.ReadString('\n')
	if true {
		fmt.Println("The input is" + input2) // Note if i use '+' the output will not have the space after is
	} else { // Note that else cant start in new line and will throw an error
		log.Fatal(error)
	}
	/*
		The if else loops look like
		if condition {
			....
		} else if condition {
			....
		} else {
			....
		}
	*/
}
