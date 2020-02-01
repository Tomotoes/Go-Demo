package main

import "fmt"

func main() {
	test := map[interface{}]interface{}{}
	test["a"] = 1
	test[1] = "A"
	fmt.Println(len(test))
}
