package main

import "fmt"

func sum(nums ...int)  int{
	fmt.Println(nums)
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
	return total
}

func main() {
	sum(1,2,3)
	sum(1,2,3,4)
	sum(1)

	arr := []int{1,2,3}
	// 必须是切片才能使用
	sum(arr...)
}