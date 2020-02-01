package main

import "fmt"

func main() {
	type myInt int
	m := map[string]int{}
	fmt.Println(m == nil)
	fmt.Println(m["a"])
	m["b"]++
	fmt.Println(m["b"])

	var n map[string]int
	fmt.Println(n == nil)
	fmt.Println(n["a"])
	n["b"]++
	fmt.Println(n["b"])

}
