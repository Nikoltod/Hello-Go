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

////Quick showcase of how to create, update and delete elements in maps.

// Insert or update an element in map m:
// m[key] = elem

// Retrieve an element:
// elem = m[key]

// Delete an element:
// delete(m, key)

// Test that a key is present with a two-value assignment:
// elem, ok = m[key]

// If key is in m, ok is true. If not, ok is false.
// If key is not in the map, then elem is the zero value for the map's element type.
// Note: if elem or ok have not yet been declared you could use a short declaration form:
// elem, ok := m[key]
