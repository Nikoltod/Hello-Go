package closures

import "fmt"

func closure() func(int) int {
	return func(i int) int {
		return i + 1
	}
}

func Init() {

	fmt.Println()

	fmt.Println("Closures")

	//making two shorthand variables and assinging them a function that will return a function which will return a data of type int - in other words a simple number
	a := closure()
	b := closure()

	for i := 0; i < 10; i++ {
		//Each of the functions "closures" will return a simple number which will be printer out
		fmt.Println(
			//the first one will be positive
			a(i),
			//the second one will be negative
			b(-i),
		)
	}
}
