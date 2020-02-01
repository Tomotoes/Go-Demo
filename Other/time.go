package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	time.AfterFunc(time.Second, func() {
		fmt.Println("Hello World")
		done <- true
	})
	<-done
}
