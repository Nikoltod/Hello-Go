package maps

import (
	"fmt"

	"github.com/Nikoltod/Hello-Go/arrays/maps/mutating-maps"
)

type Coords struct {
	x, y float64
}

func mapSample() {
	var a = make(map[string]Coords)

	a["Alpha"] = Coords{
		50.5162, 23.5561,
	}

	fmt.Println(a["Alpha"])

	// A map maps keys to values.

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	// The make function returns a map of the given type, initialized and ready for use.
}

func mapLiterals() {
	var m = map[string]Coords{
		"Alpha": Coords{
			50.5162,
			23.5561,
		},
		"Beta": Coords{
			52.5127,
			23.5561,
		},
	}

	fmt.Println(m)

	fmt.Println("\n We can also declare it with a shorter syntax \n")

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var n = map[string]Coords{
		"SAP":    {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}

	fmt.Println(n)

	//Map literals are like struct literals, but the keys are required.
}

func Init() {

	fmt.Printf("\n Maps example in GoLang \n")

	fmt.Printf("\n Map declaration \n")
	mapSample()

	fmt.Printf("\n Map literal declaration \n")
	mapLiterals()

	fmt.Printf("\n Mutating maps showcase \n")
	mutatingMaps.Init()

}
