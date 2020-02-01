package test

import "fmt"

func init() {
	fmt.Println("test1")
}

func getName() string {
	fmt.Println("getName")
	return "name"
}

var (
	s = getName()
)

const test = 1

func Hello() {
	fmt.Println("Hello")
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		fmt.Println("defer2")
	}()
}

//func main() {
//	Hello()
//}
