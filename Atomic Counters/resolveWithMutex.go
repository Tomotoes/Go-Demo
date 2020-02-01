package main

import (
	"fmt"
	"sync"
)

func main() {
	var ops int
	var m sync.Mutex

	for i := 0; i < 10000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			ops++
		}()
	}
	fmt.Println(ops)
}