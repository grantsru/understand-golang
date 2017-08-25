package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // Prints value
	fmt.Println(&a) // Prints location in memory (address)

	var b = &a      // Points to location in memory for a
	fmt.Println(b)  // Prints location in memory (address)
	fmt.Println(*b) // Prints value

	*b = 42        // b says, "The value at this address, change it to 42."
	fmt.Println(a) // 42

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add an * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
