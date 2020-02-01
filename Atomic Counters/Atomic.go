package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops int64
	for i := 0; i < 10000; i++ {
		go func() {
			atomic.AddInt64(&ops,1)
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println(opsFinal)
}
