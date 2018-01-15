package main

import (
	"fmt"

	"github.com/Nikoltod/Hello-Go/functions/closures"
	"github.com/Nikoltod/Hello-Go/functions/factorial"
	"github.com/Nikoltod/Hello-Go/functions/fibonacci"
)

func returnFunc() func() string {
	return func() string {
		return "Cuckoo"
	}
}

func birdCall(song func() string) {
	birdcall := song
	////or we can assign it to another variable and use it as a delegate... - if we call it simply as "returnFunc"
	//birdsong := birdcall()
	fmt.Printf("The bird goes %s", birdcall())
}

func main() {

	fmt.Println("Funtions")

	closures.Init()

	fmt.Println()

	fmt.Println("Functions which accept other functions as arguments :")

	birdCall(returnFunc())

	fmt.Println()

	fmt.Println("Fibonacci sequence with functions :")

	fibonacci.Init()

	fmt.Println()

	fmt.Println("Factorial calculation with functions :")

	factorial.Init(5)

	fmt.Println()
}
