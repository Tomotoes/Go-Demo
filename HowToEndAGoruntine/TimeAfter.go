package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := time.After(time.Second * 3)
	go func(t <-chan time.Time) {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-t:
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(timeout)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
