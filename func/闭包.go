package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"

	c := func(str string) string {
		t += " " + str
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()

	fmt.Println(a("Simon"))
	fmt.Println(b("Leann"))

	fmt.Println(a("!!!"))
	fmt.Println(b(":D"))
}
