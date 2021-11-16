package main

import (
	"fmt"
)

func main() {
	c := incrementor(2)
	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				c <- fmt.Sprint("Process: ", i, " Printing: ", j)
			}
			done <- true
		}(i)
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
}
