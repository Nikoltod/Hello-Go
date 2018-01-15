package factorial

import "fmt"

func factorial() func(int) int {
	return func(i int) int {
		return i * i
	}
}

//i - the number from which the factorial will be fabricated
func Init(x int) {
	f := factorial()
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println(f(x + i))
	}
}
