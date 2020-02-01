package main

import "sync"

func merge(cs ...chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})
	output := func(c <-chan interface{}) {
		for e := range c {
			out <- e
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
