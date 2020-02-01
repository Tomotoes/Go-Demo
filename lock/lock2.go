package main

import (
	"fmt"
	"sync"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter2 int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				fmt.Println("lock failed")
				return
			}
			counter2++
			fmt.Println("current counter2: ", counter2)
			l.Unlock()
		}()
	}
	wg.Wait()

}
