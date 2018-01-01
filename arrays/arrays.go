package main

import (
	"github.com/Nikoltod/Hello-Go/arrays/declaration"
	"github.com/Nikoltod/Hello-Go/arrays/slicing"
)

func main() {
	//main program for arrays

	declaration.Init()

	slicing.Init()
}

// Output :

// Array declaration :

// [1 2]
// [3 4]

// Sliced arrays :

// slicedArr = [2 3 4]

// Slicing references :

// [Paul George John William]
// [Paul George] [George John]
// [Paul XXX] [XXX John]
// [Paul XXX John William]
