package test

import "fmt"

func init() {
	fmt.Println("test2")
}

var Test = getAge()

func getAge() int {
	fmt.Println("get age")
	return 18
}

func Say() {
	fmt.Println(test)
	Hello()
	fmt.Println("say")
}

func main() {
	fmt.Println(test)
	Hello()
}
