package slicing

import (
	"fmt"

	"github.com/Nikoltod/Hello-Go/arrays/slicing/slicingReferences"
)

func Init() {
	fmt.Println("\n")
	fmt.Println("Sliced arrays : \n")

	var arr = [5]int{1, 2, 3, 4, 5}

	// give me a slice of 'arr's values from index 1 through index 3
	var slicedArr = arr[1:4]

	// output - {2, 3, 4}
	fmt.Printf("slicedArr = %v \n", slicedArr)

	fmt.Println("\n")
	fmt.Println("Slicing references : \n")

	slicingReferences.Init()
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
