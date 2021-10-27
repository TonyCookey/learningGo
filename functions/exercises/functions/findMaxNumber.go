package functions

func FindMaxNumber(x ...int) int {
	var maxNumber int
	for _, i := range x {
		if i > maxNumber {
			maxNumber = i
		}
	}
	return maxNumber
}
