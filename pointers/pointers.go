package main

import "fmt"

// func name (arguments) returns
// if it returns anything there should be in func declaration
func double(num int) int {
	num *= 2
	return num
}

// Now to rewrite the function above
// so now this function get a pointer
func pointerDouble(number *int) {
	// update the value at the pointer
	*number *= 2
}

func main() {
	amount := 6
	amount = double(amount)
	fmt.Println(amount)
	//pointer to get the address
	fmt.Println(&amount)

	// anything that has * in front of it is pointer variable
	integer := 5
	// getting the address where it is stored
	address_to_integer := &integer
	// when you give & in front of variable it will read as pointer to the variable
	fmt.Printf("%T\n", &integer)
	// when you do *address_to_integer it means value at the pointer
	fmt.Println(*address_to_integer)

	// Using Pointer
	a := 6
	//Pass the pointer
	pointerDouble(&a)
	// You will notice we changed the values using pointer and not having a return type
	fmt.Println(a)
}
