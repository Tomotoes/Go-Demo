package main

type F func(a, b int) func()

func FP(arg F, a, b int) func() {
	return arg(a, b)
}

func main() {

}
