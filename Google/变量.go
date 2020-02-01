package main

import "fmt"

func main() {
	var s string = `%d xxx`
	fmt.Printf("%q\n",s)
	fmt.Printf("%s\n",s)
	fmt.Println(s)

	one := 0
	one, two := 1, 2	// two 是新变量，允许 one 的重复声明。比如 error 处理经常用同名变量 err

	fmt.Println(one,two)
}