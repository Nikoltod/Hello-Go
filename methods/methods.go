package main

import (
	"github.com/Nikoltod/Hello-Go/methods/methodPointers"
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//reciever - the reciever is a special argument for methods.
//here we added a reciever argument of type Vertex which is a struct.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	fmt.Println("Methods & Recievers")
	v := Vertex{3, 4}

	fmt.Print(v.Abs())

	fmt.Println()

	fmt.Println("Pointer recievers")

	methodPointers.Init()

	fmt.Println()

}


//methods are just functions with a reciever argument.

//There are several reasons to choose pointer recievers over value recievers
//You can modify the value that its reciever points to.
//You can avoid copying the value on each method call. This can be more efficient if the reciever is a large struct.


