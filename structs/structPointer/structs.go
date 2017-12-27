package main

import "fmt"

type Strut struct {
	A int
	B int
}

func main() {
	s := Strut{1, 2}
	p := &s // pointing to the address in memory of the 's' struct
	p.A = 3 // assinging new value to the 'A' field of the 's' struct via pointer
	fmt.Println(s)
}
