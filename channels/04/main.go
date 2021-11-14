package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 30; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for k := range c {
				fmt.Println(k)
			}
			done <- true
		}()
	}
	for i := 0; i < n; i++ {
		<-done
	}
	close(done)
}
