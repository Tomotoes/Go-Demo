package main

import "time"

func main() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		println("你好")
	}
}
