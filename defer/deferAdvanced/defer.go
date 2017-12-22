package main

import "fmt"

//1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func printZero() {
	i := 0
	defer fmt.Println(i)
	// i will be '0' because the value of 'i' will be incremented after the execution of fmt.Println(i) will occur
	i++
	return
}

//2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func execCount() {
	//This loop will print 43210 instead of 01234 - because the "defer" execution is LIFO - Last In First Out
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//This can be quite dangerous - meaning that values can be changed in compile time.
//3. Deferred functions may read and assign to the returning function's named return values.
func reasignReturn() {
	i := 0
	defer func() { i++ }()
	return
}

func main() {
	//I will call the functions in sequence 3,2,1
	defer printZero()
	defer execCount()
	defer reasignReturn()
}
