package main

import (
	"fmt"
	"reflect"
)

const (
	a int = iota
	b
	c
	d string = "1"
	e
)

func main() {
	fmt.Println(a, b, c, d, e)

	t := reflect.ValueOf(e).Kind()
	fmt.Println(t)
}
