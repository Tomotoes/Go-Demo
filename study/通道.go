package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	nums := make(chan int, 2)
	nums <- 1
	nums <- 2

	fmt.Println(<-nums)
	fmt.Println(<-nums)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)

	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout!")
		}

		queue := make(chan string ,2)
		queue <- "one"
		queue <- "two"
		for elem := range queue{
			fmt.Println(elem)
		}
	}
}
