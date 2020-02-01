package main

import (
	"fmt"
	"log"
	"reflect"
)

func Reflect(arr interface{}) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		log.Fatal("type error")
		return
	}
	n := v.Len()
	for i := 0; i < n; i++ {
		fmt.Println(v.Index(i)) // 输出值
	}
}

func main() {
	s := []int{1, 3, 5, 7, 9}
	b := []float64{1.2, 3.4, 5.6, 7.8}
	Reflect(s)
	Reflect(b)
}
