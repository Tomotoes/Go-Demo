package main

import (
	"fmt"
	"math"
)

func Iter() <-chan int {
	iter := make(chan int)
	go func() {
		i := 0
		for i < math.MaxInt64 {
			iter <- i
			i++
		}
		close(iter)
	}()
	return iter
}

func main() {
	for i := range Iter() {
		fmt.Println(i)
		if i == 100 {
			break
		}
	}
}
