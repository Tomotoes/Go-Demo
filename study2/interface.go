package main

import "fmt"

type Quacker interface {
	Quack()
}
// 组合接口
type Duck interface {
	Quacker
}

type Man struct {
}

func (m Man) Quack() {
	fmt.Println("这个人在模仿鸭子叫~")
}

func InvokeQuack(duck Duck) {
	duck.Quack()
}

func main() {
	InvokeQuack(Man{})
}
