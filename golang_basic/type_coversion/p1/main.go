package main

import "fmt"

func main() {

	var a float32 = 123.456
	var b int = int(a)
	var c float32 = float32(b)

	// Print out the types
	fmt.Printf("%T\t%T\t%T\n", a, b, c)

	// Print out the values
	fmt.Printf("%v\t%v\t%v\n", a, b, c)

}
