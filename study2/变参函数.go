package main

import "fmt"

func Sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for num := range nums {
		total += num
	}
	return total
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	r := Sum(nums...)
	var result int = r
	fmt.Println(result)
}
