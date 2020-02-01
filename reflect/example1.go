package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 1997.1109
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
	fmt.Println("CanSet: ", v.CanSet())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	p = p.Elem()
	fmt.Println("The Elem of p is: ", p)
	fmt.Println("settability of p:", p.CanSet())
	p.SetFloat(3.1415) // this works!
	fmt.Println(p.Interface())
	fmt.Println(p)
}
