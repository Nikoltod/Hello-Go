package maps

import "fmt"

type Square struct {
	x, y float64
}

func mapSample() {
	var a = make(map[string]Square)

	a["Alpha"] = Square{
		50.5162, 23.5561,
	}

	fmt.Println(a["Alpha"])

	// A map maps keys to values.

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	// The make function returns a map of the given type, initialized and ready for use.
}

func mapLiterals() {
	var m = map[string]Square{
		"Alpha": Square{
			50.5162,
			23.5561,
		},
		"Beta": Square{
			52.5127,
			23.5561,
		},
	}

	fmt.Println(m)
}

func Init() {

	fmt.Printf("\n Maps example in GoLang \n")

	fmt.Printf("\n Map declaration \n")
	mapSample()

	fmt.Printf("\n Map literal declaration \n")
	mapLiterals()

}
