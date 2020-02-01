package main

import "fmt"

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	q, r := div(3, 4)
	fmt.Println(q, r)
}
