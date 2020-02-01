package main

import (
	"fmt"
	"time"
)

func printStr(str string) {
	fmt.Println(str)
}

func add(x, y int) (result int) {
	result = x + y
	return
}

func main() {
	defer printStr("defer 1")
	defer printStr("defer 2")
	defer printStr("defer 3")

	fmt.Println("main 1")

	time.Sleep(time.Microsecond)

	add(1, 3)

}
