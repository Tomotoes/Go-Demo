package main

import "fmt"

var Cache = map[int]int{0: 0, 1: 1}

func Fusc(n int) int {
	v, had := Cache[n]
	if had {
		return v
	}
	var r int
	if n%2 == 0 {
		r = Fusc(n / 2)
	} else {
		r = Fusc(n/2) + Fusc(n/2+1)
	}
	Cache[n] = r
	return r
}

func main() {
	fmt.Println(Fusc(4))
	fmt.Println(Fusc(4545642138))
	fmt.Println(Fusc(8))
}
