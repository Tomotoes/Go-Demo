package main

import "fmt"

func random() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}

func main() {
	r := random()
	for i := 1; i <= 10; i++ {
		fmt.Println(<-r)
	}
}
