package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-done:
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		done <- struct{}{}
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
