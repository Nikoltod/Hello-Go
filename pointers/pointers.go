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

func simplePointers(p *int) {
	fmt.Printf("value of the passed pointer is : %v\n", *p)             //leads to the value of type T in memory - for example : 12
	fmt.Printf("address in memory of the passed pointer is : %v\n", &p) //leads to the address in memory of the passed pointer - for example : 0x1040c130
}

func main() {

	a := 12
	p := &a // points to the address in memory

	// '&' address in memory
	// '*' value in memory
	simplePointers(p) // passing the address in memory

	basicPointers()

	advancedPointers()

}
