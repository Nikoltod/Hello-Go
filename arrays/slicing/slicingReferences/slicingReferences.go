package slicingReferences

import "fmt"

func workingSlicedArray(names [4]string) {
	fmt.Println("/////////////////////////////////////// Normal working example ///////////////////////////////////////")
	fmt.Println(names)

	slicedNamesOne := names[0:2]

	fmt.Println(slicedNamesOne)

	slicedNamesTwo := names[0:2]

	slicedNamesTwo[1] = "XXX Nicholas XXX"

	fmt.Println(slicedNamesTwo)
	fmt.Println(names)

	fmt.Println("///////////////////////////////////////")
}

// This showcases how when assigning value to the index that outside of the range of the array which will blow up with a "Index out of range exception" - hopefully we can handle this with the : 'recover' statement
func brokenSlicedArray(names [4]string) {

	//The following code will be invalid because we're slicing a array from position 1 to 2 which means that we will get only 1 index from the array and we will assign value to the Index 1 which will be out of range for the array.
	fmt.Println("/////////////////////////////////////// Index out of range with recover ///////////////////////////////////////")
	fmt.Println(names)

	slicedNamesOne := names[0:2]

	fmt.Println(slicedNamesOne)

	// Correct will be // slicedNamesTwo := names[0:2]
	// this will 'panic' must recover
	slicedNamesTwo := names[1:2]

	// The program is going to 'panic' and enter the if statement.
	if r := recover(); r != nil {
		slicedNamesTwo[1] = "XXX Nicholas XXX"
	}

	fmt.Println(slicedNamesTwo)
	fmt.Println(names)

	fmt.Println("///////////////////////////////////////")
}

func arrayReferences(arr [4]string) {
	fmt.Println(arr)

	a := arr[0:2]
	b := arr[1:3]

	fmt.Println(a, b)

	// Assigning new value to index '0' from the incoming array
	b[0] = "XXX"

	// The values in 'b' have changed
	fmt.Println(a, b)

	// The value in 'arr' has changed also - slices are like references to arrays
	fmt.Println(arr)

	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 	A slice does not store any data, it just describes a section of an underlying array.
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Other slices that share the same underlying array will see those changes.
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
}

func Init() {

	var names = [4]string{
		"Paul",
		"George",
		"John",
		"William",
	}

	//The following two functions 'brokenSlicedArray' and 'workingSlicedArray' are just for fun when we're trying to access Index out of range in a array - that's handled with 'recover' statement
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//// both examples will work - but if you remove the assignment of the array on line #41 outside the 'if' statement with 'recover' the program will begin to 'panic' and you will get a 'Index out of range exception'
	//brokenSlicedArray(names)

	//workingSlicedArray(names)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	arrayReferences(names)

}
