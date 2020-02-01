package main

import "fmt"

const (
	a = 1
	b
	c
	d
	e = iota
	f
)

func main() {

	fmt.Println(a, b, c, d, e, f)
}
