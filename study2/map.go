package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["simon"] = 18
	ages["leann"] = 19
	ages["john"] = 20

	delete(ages, "john")

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	if age, ok := ages["simon"]; ok {
		fmt.Println("my age is ", age)
	}
}
