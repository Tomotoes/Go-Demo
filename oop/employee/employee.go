package employee

import (
	"fmt"
)

type IEmployee interface {
	Fullname() string
}

type Employee struct {
	Firstname string
	Lastname  string
}

func (e Employee) Fullname() string {
	return fmt.Sprintf("%s %s", e.Firstname, e.Lastname)
}

func New(Firstname string, Lastname string) Employee {
	e := Employee{Firstname, Lastname}
	return e
}
