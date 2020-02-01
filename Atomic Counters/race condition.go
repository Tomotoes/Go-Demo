package main

import (
	"fmt"
	"time"
)

func main() {
	var ops int
	for i := 0; i < 10000; i++ {
		go func() {
			ops++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(ops)
	// 9976
}


