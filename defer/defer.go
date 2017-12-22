package main

import "fmt"

func main() {

	//The last "defer" statement will execute first before the other "defer" statements

	//The output will be :
	//3, 2, 1

	// This code will execute last.
	defer fmt.Println("1")

	// This code will execute second.
	defer fmt.Println("2")

	// This code will execute first.
	fmt.Println("3")

}

//https://github.com/Nikoltod/Hello-Go/issues/1
