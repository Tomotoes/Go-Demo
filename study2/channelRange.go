package main

import "fmt"

func main() {
	queue := make(chan string,2)
	queue <- "Simon"
	queue <- "Leann"
	close(queue)

	for name := range queue{
		fmt.Println(name)
	}
}