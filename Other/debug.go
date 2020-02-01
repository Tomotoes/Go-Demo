package main

import "fmt"

func forEach(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)

	}
}

func main() {
	name := "Simon"
	fmt.Println(name)
	forEach(5)
}
