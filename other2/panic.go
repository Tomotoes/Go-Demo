package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(1)
	//panic("error")
}