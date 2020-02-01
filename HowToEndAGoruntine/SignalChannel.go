package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan os.Signal, 2)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

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

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
