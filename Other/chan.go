package main

func main() {
	ch4 := make(chan chan chan chan int)
	ch3 := make(chan chan chan int)
	ch2 := make(chan chan int)
	ch1 := make(chan int)
	ch1 <- 0
	ch2 <- ch1
	ch3 <- ch2
	ch4 <- ch3

	ch1 = <- ch2
	ch2 = <- ch3
	ch3 = <-ch4
}
