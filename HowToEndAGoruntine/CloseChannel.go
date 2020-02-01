package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-ch:
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		time.Sleep(time.Second * 3)
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
