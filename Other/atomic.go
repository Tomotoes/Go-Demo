package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func worker(name string) {
	defer wg.Done()
	for {
		time.Sleep(time.Millisecond * 250)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("showDown: ", name)
			break
		}
	}
}

func main() {
	wg.Add(2)
	go worker("A")
	go worker("B")
	time.Sleep(time.Second)
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}
