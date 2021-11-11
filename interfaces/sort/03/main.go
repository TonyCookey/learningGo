package main

import (
	"fmt"
	"sort"
)

func main() {
	xn1 := []int{2, 56, -4, 9, 0, 65, 109, -34, -9, 18, 105}

	sort.Ints(xn1)
	if sort.IntsAreSorted(xn1) {
		//Prints the slice of ints if they are sorted in increasing order
		fmt.Println(xn1)
	} else {
		fmt.Println("Integers are not sorted")
	}
}
