package main

import (
	"fmt"
	"math"

	"github.com/Nikoltod/Hello-Go/interfaces/interface-values"

	"github.com/Nikoltod/Hello-Go/interfaces/implicit-interfaces"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Printf("Method 1 with float value reciever : %v", a.Abs())

	fmt.Println()

	a = &v // a *Vertex implements Abser
	fmt.Printf("Method 2 with pointer reciever : %v", a.Abs())

	fmt.Println()
	fmt.Println()

	fmt.Println("Declaring implicit interfaces")

	implicitInterfaces.Init()

	fmt.Println()
	fmt.Println()

	fmt.Println("Interface values")

	interfaceValues.Init()

	fmt.Println()
	fmt.Println()

}

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
