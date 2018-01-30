package interfaceValues

import (
	"fmt"
	"math"
)

//interface I
type I interface {
	M()
}

//structure T
type T struct {
	S string
}

//method implements interface I's method M()
func (t *T) M() {
	fmt.Println(t.S)
}

//Interface value of F of type float64
type F float64

//method implements interface I's method M()
func (f F) M() {
	fmt.Println(f)
}

//the 'describe' method prints out the value and type[T] of 'i' which is of type interface 'I'
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Init() {
	//variable 'i' of type interface - I is assigned pointer to a structure of type T
	var i I = &T{"Hello"}

	//describe uses the variable 'i' of type interface - I and returns a result
	describe(i)

	//variable 'i' uses the method M() of with pointer reciever - at line #16 - (t *T) M() - should return string from the structure of type T.
	i.M()

	//variable 'i' is assigned a new value of interface holding a value of F of data type float64
	i = F(math.Pi)

	//describe uses the new value of 'i' printing out the value and type F of interface that's been passed as an argument.
	describe(i)

	//newly assigned variable of 'i' uses the method of M() with value reciever - at line #23 - (f F) M() - should return float64 from interface F of type float64 - which is the value that it holds.
	i.M()
}
