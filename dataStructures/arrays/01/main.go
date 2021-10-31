package main

import "fmt"

func main() {
	var x [50]int
	var y [50]string
	fmt.Println(x, len(x))
	for i := 65; i < 100; i++ {
		x[i-65] = i
		y[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(y)

}
