package Fan

import "sync"

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(cs))

	out := make(chan int)
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _,c :=range cs{
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
