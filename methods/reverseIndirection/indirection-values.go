package reverseIndirection

import (
	"math"
	"fmt"
)

type Matrix struct {
	X, Y float64
}

func (m Matrix) Abs() float64{
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}


func AbsFunc(m Matrix) float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

func Init() {
	//functions must take a value of that specifit type T
	m := Matrix{3, 4}
	fmt.Println(m.Abs())
	fmt.Println(AbsFunc(m))

	//method with value recievers take either a value or a pointer as the reciever when they're called
	a := &Matrix{4, 3}
	fmt.Println(a.Abs())
	fmt.Println(AbsFunc(*a))
}

// Functions that take a value argument must take a value of that specific type: 

// // var m Matrix
// // fmt.Println(AbsFunc(m))  // OK
// // fmt.Println(AbsFunc(&m)) // Compile error!

// while methods with value receivers take either a value or a pointer as the receiver when they are called: 

// // var m Matrix
// // fmt.Println(m.Abs()) // OK
// // a := &m
// // fmt.Println(a.Abs()) // OK

// In this case, the method call a.Abs() is interpreted as (*a).Abs(). 