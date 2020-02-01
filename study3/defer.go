package main

import "fmt"

// 1. defer 会延迟到 return 之后执行

// 2. 一个函数中 多个defer 的执行顺序模型是 栈模型

// 3. defer 函数中的所引用的变量状态 可以按照函数参数传递方式理解
//    除了 map slice chan pointer 其他类型都是值类型 , copy 值传递
//         map slice chan pointer 都属于引用类型 , 引用同一地址传递
//    所以, 针对值类型来说: defer 函数中的值类型变量的值, 都是defer 定义时 , 该值变量所处的值
//          针对引用类型来说: defer 函数中的引用类型变量的值, 都是最终值

// 4. defer 可以修改已经返回的 引用类型变量的值 , 可以修改函数返回值类型的申明变量 func()(r int) 此例中是 r
// 5. panic recover 有奇效 , panic后如果不调用recover , 会使程序退出. 所以 可以把 recover 放到 defer 函数中

// 语法 defer func(){}()

func example1() int {
	i := 0
	defer fmt.Println("in defer :", i)
	i = 1000
	defer fmt.Println("in foo :", i)
	return i + 24
}

func example2() (i int) {
	i = 0
	defer func() {
		i = 996
	}()
	defer func() {
		i--
		fmt.Println("第一个defer", i)
	}()

	i++
	fmt.Println("+1后的i：", i)

	defer func() {
		i--
		fmt.Println("第二个defer", i)
	}()

	i++
	fmt.Println("再+1后的i：", i)

	defer func() {
		i--
		fmt.Println("第三个defer", i)
	}()

	i++
	fmt.Println("再+1后的i：", i)

	return
}

func example3() {
	a := []int{1, 2, 3, 4}
	defer func() {
		fmt.Println("d1:", a)
	}()
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	defer func() {
		fmt.Println("d2:", a)
	}()
}

func example4() map[string]int {
	m := map[string]int{"age": 18}
	defer func() {
		m["age"]++
	}()

	return m
}
/*
在函数执行过程中，有可能在很多地方都会出现panic，panic后如果不调用recover，程序会退出，
为了不让程序退出，我们需要在panic后调用recover，
但，panic后的代码不会执行，recover是不可能在panic后调用，
然而panic所在的函数内defer指定的函数可以执行，所以recover只能在defer指定的函数中被调用，并且只需要在1个defer指定的函数中处理。

*/
func panicFunc() {
	defer func() {
		fmt.Println("before recover")
		recover()
		fmt.Println("after recover")
	}()

	fmt.Println("before panic")
	panic(0)
	fmt.Println("after panic")
}

func main() {
	fmt.Println("in main :", example1())
	/*
		in foo: 1000
		in defer : 0
		in main func: 1024
	*/

	println(example2())
	/*
		+1后的i： 1
		再+1后的i： 2
		再+1后的i： 3
		第三个defer 2
		第二个defer 1
		第一个defer 0
	*/
	example3()

	fmt.Println(example4())
	//	map[age:19]

	panicFunc()
}
