package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	data map[string]int
	mux  sync.Mutex
}

func (c *SafeCounter) Set(key string) {
	c.mux.Lock()
	c.data[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Get(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data[key]
}

func main() {
	c := SafeCounter{data: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Set("key")
	}
	time.Sleep(time.Second)

	fmt.Println(c.Get("key"))
}
