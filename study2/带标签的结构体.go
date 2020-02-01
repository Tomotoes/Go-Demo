package main

import (
	"fmt"
	"reflect"
)

type Simon struct {
	name string "simon's name"
	age  int    "simon's age"
	sex  bool   "simon's bool"
}

func main() {
	simon := Simon{"simon", 18, true}
	reflectSimon := reflect.TypeOf(simon)
	fmt.Printf("%#v\n",reflectSimon.Field(0))
	fmt.Printf("%#v\n",reflectSimon.Field(1))
	fmt.Printf("%#v\n",reflectSimon.Field(2))
}
