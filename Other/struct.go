package main

import (
	"fmt"
)

type User struct {
	name string
}

func (u User) String() string {
	return fmt.Sprintf("%s", u.name)
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u User) Say(suffix string) {
	fmt.Println(u.name + suffix)
}

func (u *User) Hello(prefix string) string {
	fmt.Println(prefix + u.name)
	return prefix
}

func main() {

	say := User.Say
	say(User{name: "Simon"}, "!") //"Simon!"
	say(User{}, "!")              // "!
	u := User{name: "Leann"}
	say(u, "!") // Leann!

	hello := (*User).Hello                   //必须显式指针
	i := hello(&User{name: "Leann"}, "Hi! ") //Hi! Leann
	fmt.Println(i)
}
