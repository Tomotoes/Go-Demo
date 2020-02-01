package main

import "fmt"

func main() {
	// 定义数组
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	a[0] = -1
	fmt.Println(a, b)

	// 定义切片
	s := []int{5}
	y := a[:]
	y[1] = 3
	y = append(y, 9)
	x := append(s, 1)
	z := append(y, 2)
	z[2] = 100
	fmt.Println(x, y, x, z)

	fmt.Println(a)

}
