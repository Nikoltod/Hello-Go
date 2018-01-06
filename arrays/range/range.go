package arrRange

import "fmt"

func Init() {

	fmt.Println("\n Range for arrays \n")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// Showcase of the range form of the for loop in GoLang

	for i, z := range pow {
		fmt.Printf("2**%d = %d\n", i, z)
	}

	// The range form of the for loop iterates over a slice or map. - IMPORTANT
	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
}
