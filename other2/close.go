package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool,1)
	go func() {
		for {
			select {
			case r := <-done:
				fmt.Println(r)
				return
			}
		}
	}()
	//done <- true

	close(done)
	time.Sleep(time.Second)

	fmt.Println("exit")
}
