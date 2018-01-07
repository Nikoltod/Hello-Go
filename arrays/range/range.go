package arrRange

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//passing the index and value by pointers - within the memory
func printIndexOrValue(index *int, value *int) {

	// if we value of 'index' is not 'nil' the index is printer out by value in memory
	if index != nil {
		fmt.Printf("%d\n", *index)
	}

	// if we value of 'value' is not 'nil' the value is printer out by value stored in memory
	if value != nil {
		fmt.Printf("%d\n", *value)
	}

	if value == nil && index == nil {
		fmt.Print("Provide either a index or a value of type 'int'.")
	}
}
func rangeNormalSyntax() {
	fmt.Println("\n Range for arrays \n")

	fmt.Println("\n Range for arrays - normal syntax \n")
	// Showcase of the range form of the for loop in GoLang

	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	// The range form of the for loop iterates over a slice or map. - IMPORTANT
	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
}

func rangeShortSyntax() {
	fmt.Println("\n Range for arrays - short syntax \n")

	// We can also skip the index or value by assigning to _.
	// If we only want the index, we can drop the value in this case "i".

	fmt.Printf("\n Only by value - if we omit the index \n")

	//example with only the value
	for _, value := range pow {
		printIndexOrValue(nil, &value)
	}

	fmt.Printf("\n Only by index - if we omit the value \n")

	//example with only the index
	for index := range pow {
		printIndexOrValue(&index, nil)
	}

}
func Init() {
	rangeNormalSyntax()

	rangeShortSyntax()
}
