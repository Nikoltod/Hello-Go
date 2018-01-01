package declaration

import "fmt"

func print(arr [2]int) {
	fmt.Println(arr)
}

func implicitArray() {
	// implicit declaration of array - create a variable and then assign it values
	var arr [2]int

	arr[0] = 1
	arr[1] = 2

	print(arr)
}

func explicitArray() {
	// explicit declaration of array - create a variable with the shorthand declaration and assign it values in the declaration
	var arr = [2]int{3, 4}

	// explicit declaration with shorthand syntax
	// arr := [2]int{1, 2}

	print(arr)
}

func Init() {
	fmt.Println("\n")
	fmt.Println("Array declaration : \n")
	//implicit
	implicitArray()

	//explicit
	explicitArray()
}
