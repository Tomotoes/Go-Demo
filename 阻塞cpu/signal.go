package main

import (
	"os"
	"os/signal"
	"syscall"
)

// BUG(astaxie):This divides by zero.
func add(){

}

func main() {

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

}
