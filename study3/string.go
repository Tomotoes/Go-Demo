package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	word := "I am Simon"
	words := strings.Fields(word)
	fmt.Println(len(words))
	fields := strings.FieldsFunc(word, func(r rune) bool {
		fmt.Println(r)
		return false
	})
	fmt.Println(fields)

	var wg sync.WaitGroup
}