package main

import "fmt"

type tony int

func main() {
	var age tony = 12

	fmt.Printf("%T, %v \n", age, age)
}
