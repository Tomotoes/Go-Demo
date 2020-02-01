package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func Producer(id int, item chan<- int) {
	for i := 0; i < 10; i++ {
		//wg.Add(1)
		item <- i
		fmt.Printf("Producer %d produces data: %d \n", id, i)
		time.Sleep(time.Second)
	}
}

func Consumer(id int, item <-chan int) {
	for {
		i := <-item
		fmt.Printf("Consumer %d get data: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func main() {

	item := make(chan int, 1)
	go Producer(1, item)
	go Consumer(1, item)

	//wg.Wait()

	time.Sleep(time.Second * 20)
}
