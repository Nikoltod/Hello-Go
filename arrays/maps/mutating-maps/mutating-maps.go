package mutatingMaps

import "fmt"

func Init() {
	y := make(map[string]int)

	y["asd"] = 44

	//creating a newline for sake of formatting
	fmt.Println()
	//shows the whole map
	fmt.Printf("Value of element with index 'dsa' is : %v \n", y["dsa"])
	fmt.Printf("Whole map : %v \n", y) //output : Whole map : map[asd:44]

	y["dsa"] = 42

	//shows just the value of that given element in the map
	fmt.Printf("Value of element with index 'dsa' is : %v \n", y["dsa"])
	//showcase of the whole map - there are two elements in the map currently
	fmt.Printf("Whole map : %v \n", y) // //output : Whole map : map[asd:44 dsa:42]

	//remove the created element in the map
	delete(y, "asd")

	//shows the whole map so you can see that the element with index "asd" is no more.
	fmt.Printf("Whole map : %v \n", y)
}
