package main

func main() {
	ch := make(chan string, 1)
	ch <- "test"
	t := make(chan string, 1)
	p := "ASd"
	t <- p
}
