package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		time.Sleep(time.Second)
		result <- job * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go Worker(i, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		r := <-results
		fmt.Println("worker received", r)
	}
}
