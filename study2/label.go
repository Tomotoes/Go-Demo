package main

import "fmt"

func main() {
LABEL:
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 10; j++ {
				if j == 2{
					break LABEL
				}
				fmt.Println(1)
		}
	}
	fmt.Println(2)
}
