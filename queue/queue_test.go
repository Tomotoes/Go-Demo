package queue_test

import (
	"fmt"
	"queue"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	Q := queue.New()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(i, j int) {
				Q.Push([]byte(fmt.Sprintf("%d %d", i, j)))
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait()

	size := Q.Size()

	for i := 0; i < size; i++ {
		fmt.Println(string(Q.Pop()))
	}

	fmt.Println("size: ", size)
}
