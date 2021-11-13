package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go dataRacer("First")
	go dataRacer("Second")
	wg.Wait()
	fmt.Println("Final Count:", counter)
}

func dataRacer(s string) {
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, atomic.LoadInt64(&counter))
	}
	wg.Done()
}
