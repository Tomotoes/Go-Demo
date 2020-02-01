package main

import (
	"fmt"
	"runtime"
)

var quit = make(chan bool)

func loop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	quit <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go loop()
	go loop()
	for i := 1; i <= 2; i++ {
		<- quit
	}
}
