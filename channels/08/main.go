package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 6, 7, 8, 9, 11, 23, 23, 45, 22, 12, 12, 34, 2, 1, 3, 133, 1, 3, 1, 22, 5, 6, 3, 66, 3)
	// FAN OUT: distribute the square work across two go routines that both read from

	//fixed the error of a deadlock
	// used var - zero value then appended to the slice rather than make the channels with their zero values already
	// make , already created the channels inside the slice
	var xch []<-chan int
	for i := 0; i < 5; i++ {
		xch = append(xch, square(in))
	}

	// FAN IN: consuming the merged output from multiple channels
	for n := range merge(xch...) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	outChan := make(chan int)

	go func() {
		for _, n := range nums {
			outChan <- n
		}
		close(outChan)
	}()
	return outChan
}
func square(ch <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for n := range ch {
			outChan <- n * n
		}
		close(outChan)
	}()
	return outChan
}
func merge(xch ...<-chan int) <-chan int {
	outChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(xch))
	for _, ch := range xch {
		go func(c <-chan int) {
			for n := range c {
				outChan <- n
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(outChan)
	}()
	return outChan
}
