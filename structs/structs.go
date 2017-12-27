package main

import "fmt"

const pi = 3.14159

// showcasing normal structure in GoLang.
type MyStruct struct {
	a int
	b int
}

// showcasing a structure that contains mixed data types.
type MyMixedStruct struct {
	a int
	b string
	c bool
	d float32
}

func main() {
	// Printing MyStruct while assigning values to it.
	fmt.Println(MyStruct{1, 2})

	fmt.Println(MyMixedStruct{1, "alphabet", true, float32(pi)})

	// Declaring new struct of certain type "MyStruct" with given values
	declaredStruct := MyStruct{32, 42}

	// Assigning new value of the 'a' field in my new structure "declaredStruct" from "MyStruct" type.
	declaredStruct.a = 1

	// Accessing the 'a' field from 'declaredStruct'
	fmt.Println(declaredStruct.a)
}
