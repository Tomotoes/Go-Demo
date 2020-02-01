package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age  int
	sex  bool
}

func main() {
	simon := User{name: "Simon", age: 12, sex: true}
	t := reflect.TypeOf(simon)
	n := t.NumField()
	for i := 0; i < n; i++ {
		fmt.Println(t.Field(i).Name) // 键名
		fmt.Println(reflect.ValueOf(simon).Field(i)) // 值
	}
}
