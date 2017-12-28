package main

import "fmt"

// will print the exact value
func printPointer(x *int) {
	fmt.Printf("Value of 'b' is : %v\n", *x)
}

func printValue(y int) {
	fmt.Printf("Value of 'c' is : %v\n", y)
}

func main() {
	var (
		a = 1
		b = &a
		c = a
	)

	a = 15

	printPointer(b) // by pointer

	printValue(c) // by value

	fmt.Printf("Value of 'a' is : %v\n", a) // original value of 'a'

}

// Simple showcase why it's a good practice to use a pointer - output

// Value of 'b' is : 15
// Value of 'c' is : 1
// Value of 'a' is : 15

// The value of 'a' changes and yet the value by pointer 'b' doesn't change unlike the value of 'c'
