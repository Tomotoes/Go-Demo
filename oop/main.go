package main

import (
	"./employee"
	"fmt"
)

func main() {
  e := employee.New("Simon","Ma")
  fmt.Println(e.Fullname())
}
