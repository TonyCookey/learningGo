package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()
	// FAN OUT: distribute the square work across two go routines that both read from
	ch1 := factorial(in)
	ch2 := factorial(in)
	ch3 := factorial(in)
	ch4 := factorial(in)
	ch5 := factorial(in)
	ch6 := factorial(in)
	ch7 := factorial(in)
	ch8 := factorial(in)
	ch9 := factorial(in)
	ch10 := factorial(in)

	// FAN IN: consuming the merged output from multiple channels
	var y int
	for n := range merge(ch1, ch2, ch3, ch4, ch5, ch6, ch7, ch8, ch9, ch10) {
		y++
		fmt.Println("Loop Counter: ", y, "\t", n)
	}
}

func gen() <-chan int {
	outChan := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				outChan <- j
			}
		}
		close(outChan)
	}()
	return outChan
}
func facts(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
func factorial(ch <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for n := range ch {
			outChan <- facts(n)
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
