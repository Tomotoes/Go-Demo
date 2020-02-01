package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-ctx.Done():
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("finish")
}
