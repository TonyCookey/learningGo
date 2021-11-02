package main

import "fmt"

func main() {
	numbers := make([]int, 10)
	// different ways of incrementing a slice item
	numbers[0]++
	numbers[1] += 1
	numbers[2] += 2
	numbers[8] = numbers[8] + 10

	fmt.Println(numbers)
}
