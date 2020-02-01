package main

import "other/test"
import (
	"fmt"
)

func getName() string {
	fmt.Println("func name1")
	return "Simon"
}

func getAge() int {
	fmt.Println("func age1")
	return 19
}

var (
	name = getName()
	age  = getAge()
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

var (
	name1 = getName()
	age1  = getAge()
)

func main() {
	fmt.Println("main")
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		fmt.Println("defer2")
	}()
	test.Hello()
	fmt.Println("main exit")

	//mai vs vs
	//main exit
	//	defer2
	//	defer1
}
