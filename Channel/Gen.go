package main

func gen() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			out <- i
			i++
		}
	}()

	return out
}
