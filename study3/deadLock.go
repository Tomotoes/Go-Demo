package main

import "fmt"

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
)

func say(s string) {
	fmt.Println(s)
	ch1 <- <-ch2
}

func main() {
	say("s")
	<-ch1
}
