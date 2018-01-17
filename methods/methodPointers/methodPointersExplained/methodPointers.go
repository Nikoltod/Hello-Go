package methodPointersExplained

import (
	"math"
	"fmt"
)

type Matrix struct {
	X, Y float64
}

func Abs(m Matrix) float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

//as arguments we're expecting the value of T in this case "Matrix" in memory
func Scale(m *Matrix, f float64) {
	m.X = m.X *f
	m.Y = m.Y *f
}

func Init() {
	m := Matrix{3, 4}

	//we're passing the address in memory of the type Matrix struct 
	Scale(&m, 10)

	fmt.Println(Abs(m))

}

//If we remove the "*" symbol from the argument in the "Scale" function we're going to get a compiler error telling us that - cannot use &m (type *Matrix) as type Matrix in argument to Scale - which is normal because we're passing the address in memory and not the value of T in memory
