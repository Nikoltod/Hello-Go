package indirection

import (
	"fmt"
)

type Matrix struct {
	X, Y float64
}

func (m *Matrix) Scale(f float64) {
	m.X = m.X*f
	m.Y = m.Y*f
}

func ScaleFunc(m *Matrix, f float64) {
	m.X = m.X*f
	m.Y = m.Y*f
}

func Init() {
	m := Matrix{3, 4}

	m.Scale(10)

	ScaleFunc(&m, 10)

	//we're declaring a struct of type Matrix that with as a pointer
	a := &Matrix{4, 3}

	a.Scale(8)

	ScaleFunc(a, 8)

	fmt.Println(m, a)
}

//Here we're comparing two approaches - one with a function pointer (Scale) and another with a method reciever (ScaleFunc).

//As we can see they are quite different in regards to implementation 
//functions with a pointer argument must take a pointer "Explicitly"
//while methods with pointer recievers can take either a value of a pointer as the reciever when they are called

// For the statement m.Scale(10), even though m is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement m.Scale(10) as (&m).Scale(10) since the Scale method has a pointer receiver. 