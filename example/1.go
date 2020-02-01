package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//state := 0

	for {
		c := readChar()

		fmt.Println(c)
	}
}

func readChar() rune {
	c,_,_ := reader.ReadRune()
	return c
}
