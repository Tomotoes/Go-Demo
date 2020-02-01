package main

import (
	"fmt"
	"time"
)

func fo(id int) chan int {
	c := make(chan int)
	go func() { c <- id }()
	return c
}

func main() {
	c1, c2, c3 := fo(1), fo(2), fo(3)

	ch := make(chan int)
	go func() {
		isTimeout := false
		for !isTimeout {
			select {
			case v1 := <-c1:
				ch <- v1
			case v2 := <-c2:
				ch <- v2
			case v3 := <-c3:
				ch <- v3
			case <-time.After(1 * time.Second):
				isTimeout = true
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
