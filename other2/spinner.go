package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("game start ...")
	go spinner(time.Millisecond * 100)
	const n = 45
	fib := fibo(45)
	fmt.Printf("\rFib(%d)=%d\n", n, fib)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
