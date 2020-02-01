package main

import "fmt"

func main() {
	arr := [5]int{1: 2, 4: 5}
	//[0 2 0 0 5]
	fmt.Println(arr)

	var t string = "hello"
	fmt.Println(t)
	fmt.Println(t[1:])
}
