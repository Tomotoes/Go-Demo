package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		os.RemoveAll(".\\"+filepath.Join(file.Name(), ".git"))
	}
}
