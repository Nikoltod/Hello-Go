package main

import (
	"fmt"
	"strconv"
)

//declare global vars
var (
	i int     = 42
	f float64 = float64(i)
	u uint    = uint(f)
)

func main() {
	printNums(i, f, u)
}

func printNums(x int, y float64, z uint) {

	//Int to string
	var a = strconv.Itoa(x)

	//Float to string
	var b = strconv.FormatFloat(y, 'E', -1, 64)

	//Uint to string
	var c = strconv.FormatUint(uint64(z), 10)

	//assign all of them with "short variable declaration" to a array
	s := []string{a, b, c}

	//Loop through the array and printout each and every one of the strings inside it.
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
