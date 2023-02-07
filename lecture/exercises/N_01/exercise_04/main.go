package main

import (
	"fmt"
	"reflect"
)

// Demand 1 - Create my own type
// ** This is called underlying type
type myOwnType int

// Demand 2 - Declare a variable of my own type
var x myOwnType

func main() {

	// Demand 3-a
	fmt.Println(x) // 0 => zero value
	// Demand 3-b
	fmt.Println(reflect.TypeOf(x))
	// Demand 3-c
	x = 42
	// Demand 3-d
	fmt.Println(reflect.TypeOf(x))

	// make a separating line
	fmt.Println("=====================================")

	// Answer of lecture
	answerOfLecture()

}

func answerOfLecture() {
	// Demand 3-a
	fmt.Println(x)

	// Demand 3-b - Print out the type of the variable "x"
	// This point is what I didn't know
	fmt.Printf("%T\n", x)

	// Demand 3-c
	x = 42

	// Demand 3-d
	fmt.Printf("%T\n", x)
}

/*

< Exercise 4 >

1. Create your own type. Have the underlying type be an int.

2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

3. in func main
	a. print out the value of the variable “x”
	b. print out the type of the variable “x”
	c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	d. print out the value of the variable “x”


*/
