package main

import "fmt"

/* PI: 定义常量，建议变量名全部大写 */
const PI float64 = 3.14159

type myType int

func init(){
	fmt.Println("Init1")
}

func init() {
	fmt.Println("Init2")
}

func main()  {
	fmt.Println("Hello World")
}
// /* 定义结构体 */
// type Person struct {
// }

// /* 定义接口 */
// type IPerson interface {
// }

// /* 程序的入口函数 */
// func main() {

// 	/* var 变量名 变量类型 */
// 	var age myType = 18

// 	/* 平行赋值 */
// 	var name, address string = "Simon", "China"

// 	/* 多行赋值 */
// 	var (
// 		num1                int    = 1
// 		firstName, lastName string = "Simon", "Ma"
// 	)

// 	fmt.Println(name)
// 	fmt.Println("Hello World")

// }
