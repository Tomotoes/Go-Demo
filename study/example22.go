package main

import (
	"fmt"
)

type user struct {
	name    string
	sex     bool
	address struct {
		city string
		no   int
	}
	inter interface {
		say()
	}
}

func main() {
	u := &user{name: "Simon", address: struct {
		city string
		no   int
	}{city: "QHD", no: 1}, inter: nil}

	m := map[user]string{user{}: "1"}
	fmt.Println(m)
	fmt.Println(u.address.city)

	for range []int{1,2,3}{
		fmt.Println(1)
	}
}
