package main

import "fmt"

// Declare there is a Variable with the IDENTIFIER "z"
// and that the VARIABLE wit the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"

// ZERO VALUE is:
// 0 for numeric types
// false for the boolean type
// 0.0 for float types
// "" for string types
// nil for pointers, functions, interfaces, slices, channels, and maps

// z assigned the zero value of type int -> 0
var z int

// 1. Declare the variable "a"
// 2. Assign the value 1029387
// 3. Declare & Assign = Initialization
var a = 1029387

func main() {
	fmt.Println(z)
	// Result: 0
	// Because the variable "z" is of type int,
	// The zero value of int is 0

	fmt.Println(a)
}
