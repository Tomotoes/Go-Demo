package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	arr := []int{1, 2, 3,}
	sum(arr...)

	ints := [...]int{1, 2, 3, 5, 6,}
	fmt.Println(len(ints))
	sum(ints[:]...)
}
