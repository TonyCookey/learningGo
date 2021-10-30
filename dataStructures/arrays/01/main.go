package main

import "fmt"

func main() {
	var x [50]int
	fmt.Println(x, len(x))
	for i := 0; i < 50; i++ {
		x[i] = i
	}
	fmt.Println(x)

}
