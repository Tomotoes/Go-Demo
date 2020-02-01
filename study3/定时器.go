package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool {
	done := make(chan bool)

	go func() {
		time.Sleep(duration)
		done <- true
	}()

	return done
}

func main() {
	timeout := timer(time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("already 1s!")
			return
		}
	}
}
