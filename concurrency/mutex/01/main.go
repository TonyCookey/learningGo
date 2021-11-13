package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go dataRacer("First")
	go dataRacer("Second")
	wg.Wait()
	fmt.Println("Final Count:", counter)
}

func dataRacer(s string) {
	for i := 0; i <= 10; i++ {
		mutex.Lock()
		counter++
		time.Sleep(time.Duration(100 * time.Millisecond))
		fmt.Println(s, counter)
		mutex.Unlock()

	}
	wg.Done()
}
