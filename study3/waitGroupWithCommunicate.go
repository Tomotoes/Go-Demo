package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go Go1(&wg)
	}
	wg.Wait()
}

func Go1(wg *sync.WaitGroup) {
	total := 0
	for i := 1; i <= 100000000; i++ {
		total += i
	}
	wg.Done()
	fmt.Println(total)
}
