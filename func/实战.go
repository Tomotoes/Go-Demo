package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Grade string
}

func filter(s []Student, f func(Student) bool) []Student {
	var r []Student
  
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {

	s1 := Student{Name: "Simon", Grade: "A"}

	s2 := Student{Name: "John", Grade: "B"}

	s3 := Student{Name: "Leann", Grade: "A"}

	students := []Student{s1, s2, s3}

  s := filter(students,func(s Student) bool{
    return s.Grade != "B"
  })
  
  fmt.Println(s)
}
