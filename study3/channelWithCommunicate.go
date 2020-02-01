package main

import "fmt"

func main() {
	finish := make(chan bool, 10)
	for i := 1; i <= 10; i++ {
		go Go(finish)
	}
	for i := 1; i <= 10; i++ {
		<-finish
	}
}

func Go(finish chan<- bool) {
	total := 0
	for i := 1; i <= 100000000; i++ {
		total += i
	}
	finish <- true
	fmt.Println(total)
}
