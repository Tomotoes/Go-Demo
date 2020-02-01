package main

import (
	"fmt"
	"time"
)

func main() {
	var ops int
	ch := make(chan struct{}, 1)
	for i := 0; i < 10000; i++ {
		go func() {
			ch <- struct{}{}
			ops++
			<-ch
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(ops)
}
