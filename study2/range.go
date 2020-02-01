package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}

	sum := 0

	for idx, num := range nums {
		fmt.Printf("%d-%d\n", idx, num)
		sum += num
	}

	fmt.Printf("sum is %d\n", sum)

	kvs := map[string]string{"name": "Simon"}

	for key, value := range kvs {
		fmt.Printf("key:%d value:%d\n", key, value)
	}

	for key := range kvs {
		fmt.Println(key)
	}

}
