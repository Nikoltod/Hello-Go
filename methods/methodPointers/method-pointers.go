package methodPointers

import (
	"fmt"
	"math"
)

type Matrix struct {
	X, Y float64
}

func (m Matrix) Abs() float64 {
	return math.Sqrt(m.X * m.X + m.Y * m.Y)
}

func (m *Matrix) Scale(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

func Init() {
	m := Matrix{3, 4}

	m.Scale(10)

	fmt.Println(m.Abs())
}

// You can declare methods with pointer receivers. 
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.) 

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers. 

// If we remove the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes. 

// With a value receiver, the Scale method operates on a copy of the original Matrix value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Matrix value declared in the main function. 

//Which means if we're using pointers that we can overrite the reciver argument or if we don't use pointers we can operate on a copy of that argument in this case "Matrix" copy and use it like a map - but in that scence we make just make a empthy map like so : make(map[] int) and be on with our business...