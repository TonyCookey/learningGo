package functions

func FuncDoubleReturns(x int) (int, bool) {
	firstReturn := x / 2
	secondReturn := false

	if x%2 == 0 {
		secondReturn = true
	}
	return firstReturn, secondReturn
}
