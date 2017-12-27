package main

import "fmt"

type Strut struct {
	A, B int
}

var (
	z = Strut{1, 2}   // has type Strut A: 1 and B: 2
	y = Strut{A: 12}  // A: 12 and B: 0
	u = Strut{}       // A: 0 and B: 0
	i = &Strut{42, 2} //has the type *Strut
)

func main() {
	fmt.Println(z, y, u, *i) // the * infront points to the value in memory of the address in memory that we've passed - "&Strut"
}
