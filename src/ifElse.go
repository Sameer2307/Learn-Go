package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your marks between 0 to 100")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	if input == "100" {
		fmt.Println("Perfect Score")
	} else if input >= "50" {
		fmt.Println("You passed")
	} else {
		fmt.Println("You failed")
	}
}
