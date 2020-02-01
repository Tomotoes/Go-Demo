package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	g
)

func main() {
	fmt.Println(a == d)
	fmt.Println(b == e)
}
