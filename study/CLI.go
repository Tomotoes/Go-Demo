package main

import (
	"flag"
	"fmt"
	"os"
)

var str string
var num int
var help bool

func main() {
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num,"num",0,"text description")
	flag.BoolVar(&help,"help",false,"Display Help")

	flag.Parse()

	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String: ",str)
	fmt.Println(">> Number: ",num)
	fmt.Println(os.Args)

	fmt.Println(flag.Args())
}
