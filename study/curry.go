package main

import "fmt"

func curry(f func(...int) interface{}, length int, args ...int) interface{} {
	if (len(args) >= length) {
		return f(args...)
	}
	return func() interface{} {
		return curry(f, length, args...)
	}
}

func _sum(nums ...int) interface{} {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	a := curry(_sum, 3, 1, 2)
	fmt.Println(a)
	r := a(3)
	fmt.Println(r)
}
