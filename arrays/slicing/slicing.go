package slicing

import (
	"fmt"
	"strings"

	"github.com/Nikoltod/Hello-Go/arrays/slicing/append"
	"github.com/Nikoltod/Hello-Go/arrays/slicing/capacity"
	"github.com/Nikoltod/Hello-Go/arrays/slicing/sliceLiterals"
	"github.com/Nikoltod/Hello-Go/arrays/slicing/slicingReferences"
)

func slicingWithMake() {
	//this will result in a empthy array - [0,0,0,0,0]
	a := make([]int, 5)

	//this will result in a sliced array of the zeroed out array we declared earlier
	a = a[0:2]

	//and later we can assign values to our new sliced array
	a[0] = 1
	a[1] = 2
	fmt.Printf("Array 'a'=%v", a)
	fmt.Println()

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	// a := make([]int, 5)  // len(a)=5
	// To specify a capacity, pass a third argument to make:
	// a := make([]int, 0, 5) // len(a)=0, cap(a)=5

	// a = a[:cap(a)] // len(a)=5, cap(a)=5
	// a = a[1:]      // len(a)=4, cap(a)=4
}

func slicingSlices() {
	// We'll create a simple print of a tic-tac-toe game

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// We'll hardcode the players choices just so we can print them out

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][0] = "X"
	board[1][0] = "O"
	board[1][2] = "X"
	board[0][1] = "O"
	board[2][1] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Well it's a draw - we can also print each move of each player so it's more interactive...
	// Here we showcased that we can create jagged arrays from slices - like in any other language - firstly we created a slice literal "board" and then we gave it values of other '3' slice literals inside it. - the rest is quite simple, we assign values to each of the slice literals by index inside the "board" jagged array or 'slice literal'.
}

func Init() {
	fmt.Println("\n")
	fmt.Println("Sliced arrays : \n")

	var arr = [5]int{1, 2, 3, 4, 5}

	// give me a slice of 'arr's values from index 1 through index 3
	var slicedArr = arr[1:4]

	// output - {2, 3, 4}
	fmt.Printf("slicedArr = %v \n", slicedArr)

	fmt.Println("\n Slicing references : \n")

	slicingReferences.Init()

	fmt.Println("\n Slicing literals : \n")

	sliceLiterals.Init()

	fmt.Println("\n Arrays slice and capacity \n")

	capacity.Init()

	fmt.Println("\n Slicing arrays with the make function \n")

	slicingWithMake()

	fmt.Println("\n Slicing arrays with the make function \n")

	slicingSlices()

	fmt.Println("\n Appending slices with built in 'append' function \n")

	append.Init()

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
/////////////////////////////////////////////////////////////////////////////////////////////////////////
// The type []T is a slice with elements of type T.
// A slice is formed by specifying two indices, a low and high bound, separated by a colon:

// a[low : high]
/////////////////////////////////////////////////////////////////////////////////////////////////////////
// This selects a half-open range which includes the first element, but excludes the last one.
// The following expression creates a slice which includes elements 1 through 3 of a:

// a[1:4]
/////////////////////////////////////////////////////////////////////////////////////////////////////////
