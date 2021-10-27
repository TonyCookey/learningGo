package main

import (
	"fmt"
	"learningGo/functions/exercises/functions"
)

func main() {
	// implementing funcDoubleReturns using func expressions
	y, z := func(x int) (int, bool) {
		firstReturn := x / 2
		secondReturn := false

		if x%2 == 0 {
			secondReturn = true
		}
		return firstReturn, secondReturn
	}(27)
	fmt.Println(y, z)

	fmt.Println(functions.FuncDoubleReturns(14))
	x := []int{1, 2, 5, 12, 45, 18, 0}
	fmt.Println(functions.FindMaxNumber(x...))
}
