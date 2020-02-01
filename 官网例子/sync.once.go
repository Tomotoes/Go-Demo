package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	fn := func() {
		fmt.Println("hello")
	}
	done := make(chan bool)
	for i := 1; i <= 10; i++ {
		go func() {
			once.Do(fn)
			done <- true
		}()
	}
	for i := 1; i <= 10; i++ {
		<-done
	}
}
