package main

import "fmt"

func generator() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(input <-chan int, target int) chan int {
	output := make(chan<- int)
	go func() {
		for {
			i := <-input
			if i%target != 0 {
				output <- i
			}
		}
	}()
	return output
}

func main() {
	const MAX = 100
	xrange := generator()

	number := <-xrange

	for number <= MAX {
		fmt.Println(number)
		xrange = filter(xrange, number)
		number = <-xrange
	}

}
