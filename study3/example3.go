package main

import (
	"fmt"
	"runtime"
)

var quita chan int = make(chan int)

func loop2(id int) { // id: 该goroutine的标号
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quita <- 0
}

func main() {
	runtime.GOMAXPROCS(2) // 最多同时使用2个核

	for i := 0; i < 3; i++ { //开三个goroutine
		go loop2(i)
	}

	for i := 0; i < 3; i++ {
		<-quita
	}
}
