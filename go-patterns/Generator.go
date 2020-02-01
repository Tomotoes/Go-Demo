package main

import "fmt"

func Iter(start, end int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func main() {
	for num := range Iter(0, 99) {
		fmt.Println(num)
	}
}
