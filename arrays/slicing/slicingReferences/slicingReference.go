package main

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

func main() {

	var names = [4]string{
		"Paul",
		"George",
		"John",
		"William",
	}

	// both examples will work - but if you remove the assignment of the array on line #41 outside the 'if' statement with 'recover' the program will begin to 'panic' and you will get a 'Index out of range exception'
	brokenSlicedArray(names)

	workingSlicedArray(names)

}
