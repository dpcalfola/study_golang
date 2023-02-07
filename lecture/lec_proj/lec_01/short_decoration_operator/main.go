package main

import "fmt"

func main() {

	x := 2
	// Short declaration operator :=
	// Declare a variable and Assign a value
	fmt.Println(x)
	x = 99
	fmt.Println(x)

	// Anonymous function
	funcFoo := func() {
		fmt.Println("I'm a func")
	}

	funcFoo()

}
