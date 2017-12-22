package main

import "fmt"

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax. - shorthand variable declaration.
const pi = 3.14159

func main() {

	// pi = 3.14159
	fmt.Println(pi)

	// // this will not work - cause constans can't be assigned new values
	// pi = 10

	// // this will throw a error - "cannot assign to pi"
	// fmt.Println(pi)
}
