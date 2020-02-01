package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	name := make(chan string,2)
	go func() {
		name <- "simon"
		fmt.Println("Simon")
	}()

	// 管道的输入输出会阻塞当前协程!

	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table
}
