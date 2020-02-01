package main

import (
	"fmt"
	"time"
)

type method func(int) int

var cache = map[int]int{}

func fibonacci(n int) (res int) {
	var ok bool
	if res, ok = cache[n]; ok {
		return res
	}
	defer func() {
		cache[n] = res
	}()
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func Judge(f method, args int) {
	start := time.Now()
	f(args)
	end := time.Now()
	fmt.Printf("func took this amount of time: %s\n", end.Sub(start))
}

func main() {
	for i := 1; i <= 25; i++ {
		Judge(fibonacci, i)
	}
}
