package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go myPrinter()
	wg.Wait()
}

func myPrinter() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(time.Second))
	}
	wg.Done()
}
