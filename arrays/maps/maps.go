package maps

import "fmt"

type Square struct {
	x, y float64
}

func Init() {

	fmt.Printf("\n Maps example in GoLang \n")

	var a = make(map[string]Square)

	a["Alpha"] = Square{
		50.5162, 23.5561,
	}

	fmt.Println(a["Alpha"])

	// A map maps keys to values.

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	// The make function returns a map of the given type, initialized and ready for use.
}
