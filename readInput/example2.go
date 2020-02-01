package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input2: ")
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input2 was: %s\n", input2)
	}
}
