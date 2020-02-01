package main

import (
	"fmt"
	"sync"
	"autoinc"
)

const (
	Start = 0
	End   = 100
	Step  = 2
)

func main() {
	var wg sync.WaitGroup

	// get the instance
	ai := autoinc.New(Start, Step)
	// close the counter
	defer ai.Close()
	for i := Start; i < End; i++ {
		wg.Add(1)
		go func() {
			// get an integer and print it
			fmt.Println(ai.Id())
			wg.Done()
		}()
	}
	wg.Wait()

	ai2 := autoinc.New(1,1)
	defer ai2.Close()
	for i := 0; i < 10; i++ {
		fmt.Println(ai2.Id())
	}
}
