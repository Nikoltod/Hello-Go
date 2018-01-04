package sliceLiterals

import "fmt"

func Init() {
	//slice literal
	fmt.Println("\n Slicing with Slice literals \n")
	arrS := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	//When you leave the default one - for low bound it's ommitted as you can see in the 'aS' example that the low bound is ommitted - 3,4,5...11 instead of 2,3...11
	aS := arrS[2:] // 3 - 11
	fmt.Println(aS)

	//As for the high bound you won't get any omitted number - 1,2,3,4,5
	bS := arrS[:5] // 1 - 5
	fmt.Println(bS)

	fmt.Println("\n Slicing with Normal array \n")
	//normal array
	arr := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	a := arr[:5]
	fmt.Println(a)

	b := arr[7:]
	fmt.Println(b)

	c := arr[3:5]
	fmt.Println(c)
}
