package main

import "fmt"

var quit chan bool

func foo(id int) {
	fmt.Println(id)
	quit <- true
}

func main() {
	count := 1000
	quit = make(chan bool,count)
	for i := 1; i <= count; i++ {
		go foo(i)
	}
	for i := 1; i <= count; i++ {
		<-quit
	}
}
