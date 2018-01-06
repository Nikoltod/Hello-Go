package append

import "fmt"

func printSlice(a []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}

func appendFromTo(arrayToAppendTo []int, arrayToAppendFrom []int) {

	for i := 0; i < len(arrayToAppendFrom); i++ {
		arrayToAppendTo = append(arrayToAppendTo, arrayToAppendFrom[i])
	}

	printSlice(arrayToAppendTo)

}

func Init() {
	var a []int
	// output : len=2 cap=2 [5]
	a = append(a, 5)
	printSlice(a)

	// output : len=2 cap=2 [5 55]
	a = append(a, 55)
	printSlice(a)

	// 	output : len=3 cap=4 [5 55 155]
	a = append(a, 155)
	printSlice(a)

	//	output : len=13 cap=16 [5 55 155 1 2 3 4 5 6 7 8 9 10]
	exampleArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	appendFromTo(a, exampleArray)

	// and we can add a whole sliced array inside it or we can add a whole array

	arrayExample := make([]int, 5)

	arrayExample[0] = 1
	arrayExample[1] = 1
	arrayExample[3] = 1

	// output : len=8 cap=8 [5 55 155 1 1 0 1 0]
	appendFromTo(a, arrayExample)

	arrayExample = []int{1, 2, 3, 4, 5}

	// output : len=8 cap=8 [5 55 155 1 2 3 4 5]
	appendFromTo(a, arrayExample)

	// so far we just keep adding arrays to our original one 'a' but if we print the capacity after each append we can see that the 'capacity' of this array is growing after each append - not to mention that when creating empthy arrays with 'make' we can add more elements to them cause they're dynamic

	// func append(s []T, vs ...T) []T - as per documentation
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
}
