package main

import "fmt"

func A(){
	fmt.Println("start AAA")
	defer func() {
		fmt.Println("end AAA")
	}()
	B()
}

func B(){
	fmt.Println("start BBB")
	defer func() {
		fmt.Println("end BBB")
	}()
	C()
}

func C(){
	fmt.Println("start CCC")
	defer func() {
		fmt.Println("end CCC")
	}()
}

func main() {
	A()
}
