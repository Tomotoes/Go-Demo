package main

import (
	"fmt"
)

type add func(a int, b int) int

var FuncValue = func() string {
	return "I am FuncValue"
}

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	a := func(person string) string {
		fmt.Println("Hello ", person)
		return person
	}("SimonMa")

	fmt.Println(FuncValue())

	fmt.Println()
	fmt.Println("%T", a)

	var judge add = func(a int, b int) int {
		return a + b
	}

	t := judge(5, 6)
	fmt.Println(t)

}