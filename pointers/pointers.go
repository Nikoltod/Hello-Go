package main

import "fmt"

//global pointer of type int without value.
var pointer *int

func basicPointers() {

	//nil value by default - pointers have no value to point to in the memory
	fmt.Printf("pointer = %v", pointer)

	//assigning value to the pointer - will have a address to point to.
	pointer := 42

	//will have a address in memory
	fmt.Printf("pointer = %v", pointer)
}

func advancedPointers() {
	a, b := 1, 2

	c := &a
	d := &b

	//c is pointing to the value a through the pointer - using it's address in memory
	fmt.Printf("c = %v \n", *c)

	//d is pointing to the value of 'b' through the pointer - using it's address in memory
	fmt.Printf("d = %v \n", *d)
}

func main() {

	basicPointers()

	advancedPointers()

}
