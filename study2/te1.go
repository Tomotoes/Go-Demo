package main

import "fmt"

var f = func(int){}

func main() {
	f = func(i int) {
		fmt.Println(i)
	}
	f(2)
	f = func(i int) {
		fmt.Println(i*i*i)
	}
	f(2)
}