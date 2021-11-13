package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go dataRacer()
	go dataRacer()
	wg.Wait()
}

func dataRacer() {
	for i := 0; i <= 10; i++ {
		counter++
		time.Sleep(100 * time.Millisecond)
		fmt.Println(counter)
	}
	wg.Done()
}
