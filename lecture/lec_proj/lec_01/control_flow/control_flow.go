package main

import "fmt"

func main() {

	// Basic syntax
	fmt.Println("1st Println")

	// Call function
	foo()

	// fop-loop
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd number:", i)
			// There's one space between string and variable
		}
	}
}

func foo() {
	fmt.Println("Print in func foo")
}

// Control Flow:
// 1. Sequence
// 2. Loop; iterative
// 3. Conditional
