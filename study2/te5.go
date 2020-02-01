package main

import "fmt"

func IAdd() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	a:=IAdd()
	fmt.Println(a())
	fmt.Println(a())
}
