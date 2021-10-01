package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{32, 18, 13, 12}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}
