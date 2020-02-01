package main

import (
	"fmt"
	"time"
)

func sum(arr []int, result chan int) {
	total := 0
	for _, v := range arr {
		total += v
	}
	result <- total
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	length := len(arr)

	result := make(chan int, 2)
	defer close(result)

	go sum(arr[length/2:], result)
	go sum(arr[:length/2], result)

	//x, y := <-result, <-result
	//fmt.Println(x, y)

	for {
		select {
		case v := <-result:
			fmt.Println(v)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout!")
			return
		}
	}

}
