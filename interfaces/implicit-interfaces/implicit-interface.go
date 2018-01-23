package implicitInterfaces

import (
	"fmt"
)

type myInterface interface {
	PrintStruct()
}

type myStruct struct {
	data string
}

// This method means type "myStruct" implements the interface "myInterface",
// but we don't need to explicitly declare that it does so.
func (s myStruct) PrintStruct() {
	fmt.Print(s.data)
}

func Init() {
	var I myInterface = myStruct{"Greetings, mortals !"}

	I.PrintStruct()
}
