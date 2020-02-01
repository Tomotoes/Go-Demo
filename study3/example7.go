package main

import (
	"fmt"
	"sync"
)

func gen(num ...int) chan int {
	out := make(chan int, len(num))
	for _, n := range num {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(chs ...<-chan int) chan int {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	out := make(chan int,1)

	output := func(ch <-chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	c := gen(2, 3)
	for n := range sq(c) {
		fmt.Println(n)
	}
	for n := range merge(sq(gen(2, 3)), sq(gen(2, 3))) {
		fmt.Println(n)
	}
}
