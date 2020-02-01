package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond * 200)

	for r := range request {
		<-limiter
		fmt.Println("request", r, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	// 刚开始的三个直接执行 , 后面的两个每隔 200 毫秒执行
	for r := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", r, time.Now())
	}

}
