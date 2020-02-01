package main

import "fmt"

func main() {
	// 一个 值为nil的channel 读取值时 , 会永远阻塞
	<-(chan int)(nil)
	fmt.Println("done")
}
