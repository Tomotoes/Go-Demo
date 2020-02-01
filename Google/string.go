package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我是Simon"
	fmt.Printf("len :%d\n", len(s))
	fmt.Printf("[]byte: %s\n", []byte(s))
	for _,v := range []byte(s) {
		fmt.Printf("%X ",v)
	}
	fmt.Println()
	for _,v := range s {
		fmt.Printf("%X ",v)
	}
	fmt.Println()
	fmt.Println("Rune Count:",utf8.RuneCountInString(s))
	for _,v := range []rune(s) {
		fmt.Printf("(%c) ",v)
	}
}
