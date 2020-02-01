package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("create: ", path)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File, data string) {
	fmt.Println("write")
	_, err := fmt.Fprintf(file, data)
	if err != nil {
		panic(err)
	}
}

func closeFile(file *os.File) {
	fmt.Println("close")
	_ = file.Close()
}

func main() {
	file := createFile("C:/Users/jinma/Desktop/test.txt")
	defer closeFile(file)
	writeFile(file, "HelloWorld")
}
