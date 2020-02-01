package main

import (
	"fmt"
)

type IPhone interface {
	call(name string) string
}

type Person struct {
	Name string
	Sex  bool
}

func (p Person) call(name string) string {
	return p.Name + "is Call " + name
}

func main() {
	p1 := Person{"Simon", false}
	str := p1.call("Leann")

	fmt.Println(str)
}
