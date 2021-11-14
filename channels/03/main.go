package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
