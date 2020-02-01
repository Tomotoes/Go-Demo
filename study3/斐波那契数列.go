package main

import "fmt"

func fibClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibChannel(n int) <-chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			ch <- a
		}
		close(ch)
	}()
	return ch
}

func main() {
	nextFib := fibClosure()
	for i := 0; i < 20; i++ {
		fmt.Println(nextFib())
	}

	for c := range fibChannel(20) {
		fmt.Println(c)
	}
}
