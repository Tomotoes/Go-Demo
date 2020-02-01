package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(1)
	mostLeft := make(chan string)
	left := mostLeft
	var N int = 10
	for i := 0; i < N; i++ {
		right := make(chan string)
		go gopher(i, left, right)
		left = right
	}
	left <- "I am hungry"
	fmt.Println(<-mostLeft)
	fmt.Println(time.Since(start))
}

func gopher(id int, left, right chan string) {
	fmt.Printf("%d wait...\n", id)
	left <- <-right
	fmt.Printf("%d receive...\n", id)
}
