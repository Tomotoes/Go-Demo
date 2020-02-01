package main

import (
	"fmt"
	"other/goruntinePool/namePrinter"
	"other/goruntinePool/worker"
	"sync"
)

var (
	names = []string{
		"simon",
		"leann",
	}
	wg = sync.WaitGroup{}
	N  = 10
)

func main() {
	p := worker.New(2)
	wg.Add(N * len(names))
	for i := 0; i < N; i++ {
		for _, name := range names {
			np := namePrinter.NamePrinter{Name: name,}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
	fmt.Println("finish")
}
