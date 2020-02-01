package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 100
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(n int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
