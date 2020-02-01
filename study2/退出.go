package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Exit 退出时, defer 函数不会执行
	defer fmt.Println("nikjbudcwo")

	//退出并且退出状态为 3。
	os.Exit(3)
	// 终端可以使用 $? 来捕获这个3
}