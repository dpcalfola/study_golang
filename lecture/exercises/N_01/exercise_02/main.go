package main

import "fmt"

// Demand 1
var x int
var y string
var z bool

func main() {

	// Demand 2-a
	fmt.Println(x) // 0
	fmt.Println(y) // ""
	fmt.Println(z) // false

	// Demand 2-b
	// 0, ""(null string), false are called zero values

}

/*

< Hand-on Exercise 2 >

1. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
	a. identifier “x” type int
	b. identifier “y” type string
	c. identifier “z” type bool
2. in func main
	a. print out the values for each identifier
	b. The compiler assigned values to the variables. What are these values called?

*/
