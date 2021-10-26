package main

import "fmt"

func main() {
	fmt.Println(funcDoubleReturns(14))
}

func funcDoubleReturns(x int) (int, bool) {
	firstReturn := x / 2
	secondReturn := false

	if x%2 == 0 {
		secondReturn = true
	}
	return firstReturn, secondReturn
}
