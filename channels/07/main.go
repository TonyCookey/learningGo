package main

import "fmt"

func main() {
	c := gen(2, 3, 4, 5, 6, 7, 8, 9)
	for n := range square(c) {
		fmt.Println(n)
	}
}
func gen(num ...int) chan int {
	outchan := make(chan int)
	go func() {
		for _, n := range num {
			outchan <- n
		}
		close(outchan)

	}()
	return outchan
}

func square(x chan int) chan int {
	outChan := make(chan int)
	go func() {
		for n := range x {
			outChan <- n * n
		}
		close(outChan)
	}()
	return outChan
}
