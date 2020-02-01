package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	log(s)
	fmt.Println(s[0])

	s = s[:0]
	log(s)

	s = s[:4]
	log(s)
}

func log(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
