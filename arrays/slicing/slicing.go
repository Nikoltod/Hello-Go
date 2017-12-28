package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	// give me a slice of 'arr's values from index 1 through index 3
	var slicedArr = arr[1:4]

	// output - {2, 3, 4}
	fmt.Printf("slicedArr = %v", slicedArr)
}
