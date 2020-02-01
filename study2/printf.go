package main

import "fmt"

type People struct {
	name string
}

func main() {
	p := People{"simon"}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	/*
	{simon}
	{name:simon}
	main.People{name:"simon"}
	main.People
	*/
}
