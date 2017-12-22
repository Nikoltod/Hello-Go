//In this example we will also take a look at the "Panic" and "Recovery" statements

// Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

// Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

package main

import "fmt"

func a(i int) {
	if i > 3 {
		fmt.Println("Panic !")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer from a", i)
	fmt.Println("State of a", i)
	a(i + 1)
}

func b() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in b", r)
		}
	}()

	fmt.Println("Calling a")
	a(0)
	fmt.Println("Recover from b")
}

func main() {
	b()
	fmt.Println("Rerutned from b")
}

//Output :
// Calling a
// State of a 0
// State of a 1
// State of a 2
// State of a 3
// Panic!
// Defer in a 3
// Defer in a 2
// Defer in a 1
// Defer in a 0
// Recover from b 4
// Returned normally from b
