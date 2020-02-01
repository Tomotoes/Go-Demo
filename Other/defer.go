package main

import "fmt"

func double(x int) (result int) {
	x = 2
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	x += x
	return x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	double(4)
	fmt.Println(triple(4))
}
