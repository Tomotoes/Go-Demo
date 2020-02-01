package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doStuff(id int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return 100 - id
}

func branch(id int) chan int {
	ch := make(chan int)
	go func() {
		ch <- doStuff(id)
	}()
	return ch
}

func fanln(branches ...chan int) chan int {
	ch := make(chan int)
	for _, c := range branches {
		go func(c chan int) {
			ch <- <-c
		}(c)
	}
	return ch
}

func main() {
	r := fanln(branch(1), branch(2), branch(3))
	for i := 1; i <= 3; i++ {
		fmt.Println(<-r)
	}
}
