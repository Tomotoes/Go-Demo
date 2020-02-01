package main

import (
	"fmt"
)

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func Random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 1:
			case c <- 0:
			}
		}
	}()
	return c
}

func main() {
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	fmt.Println(max)

	for i := range Random(100) {
		fmt.Println(i)
	}

	///* 禁止main退出的三种方法 */
	//defer func() {for {}}()
	//defer func() { select {} }()
	//defer func() {<-make(chan struct{})}()
}
