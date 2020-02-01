package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	ch := make(chan int)
	var v atomic.Value
	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				if v.Load().(bool) {
					close(ch)
					return
				}
				i++
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		v.Store(false)
		time.Sleep(time.Second * 3)
		v.Store(true)
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
