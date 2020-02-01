package main

import "fmt"

type API interface {
	show()
}

type man struct {
	name string
}

func (m man) show() {
	fmt.Println(m.name)
}

type woman struct {
	name string
}

func (w woman) show() {
	fmt.Println(w.name)
}

type parent struct {
	man
	name string
}

func (p parent) show() {
	fmt.Println(p.name)
}

func say(a API) {
	a.show()
}

func main() {
	m := man{"simon"}
	w := woman{"leann"}
	p := parent{name: "parent",man:man{name:"main"}}
	for _, i := range []API{m, w, p} {
		say(i)
	}
	p.man.show()
}
