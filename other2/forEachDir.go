package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	upperDirPattern := "./*"
	matches, err := filepath.Glob(upperDirPattern)
	if err != nil {
		panic(err)
	}
	for _, file := range matches {
		fmt.Println(file)
	}
}
