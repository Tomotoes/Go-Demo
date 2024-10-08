package main

import (
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	log.SetFlags(log.Llongfile)
	log.Print("asd")
}
