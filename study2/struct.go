package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) Error() string {
	return fmt.Sprintf("%d - %s", u.age, u.name)
}

func main() {
	u := &User{"Simon", 13}
	clone,ok := u.(*User)

	fmt.Println(name, age)
}
