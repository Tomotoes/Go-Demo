package main

func test(...bool) {

}

func main() {
	test(true, false)
	a := false
	b := true
	a = true
	b = true
	test(a, b)
	test([]bool{a, b}...)
}
