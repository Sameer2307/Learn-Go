// This is for showing how variables are declared

package main

import "fmt"

func main() {
	// var variable_name value_it_holds
	// this is variable declaration
	var value int
	var name string
	// once you have declared them its variablename = value
	// this is value assignment
	value = 7
	name = "Sameer Saxena"
	// using the variables
	fmt.Println(name)
	fmt.Println("has the lucky number", value)

	//second way of declaring values
	var bday int = 23
	var month string = "July"

	fmt.Println("His Bday is on", month, bday)

	//Third and easiest way to declare variables called short variable declaration
	year := 1996
	fmt.Println("The year of birth is", year)

	//Things to Note:
	//1. you cant assign same value to multiple var like a,b:=2 the numbers should match
	//2. Variables can only be assigned same typed values
	//3. All declared variables must be used in the program or it shows error.
}
