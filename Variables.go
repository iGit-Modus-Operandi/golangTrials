package main

import "fmt"

// declaring variable at the package level
// cannot use short-hand syntax here
var e int = 4

func main() {
	// used when you don't want to assign a value or initialize the variable yet because of future use
	// defaults to zero
	var a int
	fmt.Println(a)

	// typical way of writing variables; convenient for other not familiar with the codebase
	var b int = 1
	fmt.Println(b)

	// short-hand way; needs initialization for the variable to determine data type
	c := 2
	fmt.Println(c)

	// shows how data type is inferred by initialized value
	d := 3
	fmt.Printf("%v, %T\n", d, d)

	// prints out variable declared at package level
	fmt.Printf("%v, %T", e, e)
}
