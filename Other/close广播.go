package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello(cancel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-cancel:
			fmt.Println("cancel")
			return
		default:
			fmt.Println("hello")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printHello(cancel, &wg)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
