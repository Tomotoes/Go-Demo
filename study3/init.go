package main

import "fmt"

func init() {
	fmt.Println("1")
}

func init() {
	fmt.Println("2")
}

func init() {
	fmt.Println("3")
}
type myInt int
func (i myInt) String() string{
	return "asd"
}

func main() {
	fmt.Println("4")
	fmt.Println(1)

	var a myInt = 2
	fmt.Println(a)
}

