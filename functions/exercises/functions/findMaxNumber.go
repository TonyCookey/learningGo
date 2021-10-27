package functions

func FindMaxNumber(x ...int) int {
	var maxNumber int
	for i, v := range x {
		if v > maxNumber || i == 0 {
			maxNumber = v
		}
	}
	return maxNumber
}
