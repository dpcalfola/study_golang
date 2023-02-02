package main

import (
	"fmt"
)

// Do not work
//x_1 := 1

// The scope of var y1 is the package scope
var y1 = 5

func main() {
	x := 1
	fmt.Println(x)

	// var Keyword
	var y = 2
	fmt.Println(y)

	// It works
	// y_1 is declared in the package scope
	fmt.Println(y1)

	// Call foo_var
	fooVar()
}

func fooVar() {
	for i := 0; i < y1; i++ {
		fmt.Println(i, ". y1 = ", y1)
	}
}

// Use short declaration operator if you can
// Do not use var keyword if you can

//
// ***** KEEP the scope of the variable as small as possible *****
//
//
//

// Var keyword variable declaration's scope is the package scope

// What is the package scope?
// The package scope is the scope of the package
// The scope of the package is the scope of the package
// The scope of the package is the scope of the package
// The scope of the package is the scope of the package
// (Sigh...)
