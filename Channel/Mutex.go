package main

import "time"

type Mutex struct {
	lock chan struct{}
}

func New() *Mutex {
	mu := &Mutex{lock: make(chan struct{}, 1)}
	mu.lock <- struct{}{}
	return mu
}

func (mu *Mutex) Lock() {
	<-mu.lock
}

func (mu *Mutex) Unlock() {
	select {
	case mu.lock <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (mu *Mutex) TryLock() bool {
	select {
	case <-mu.lock:
		return true
	default:
	}
	return false
}

func (mu *Mutex) TryLockWithTimeout(timeout time.Duration) bool {
	select {
	case <-mu.lock:
		return true
	case <-time.After(timeout):
	}
	return false
}

func (mu *Mutex) IsLocked() bool {
	return len(mu.lock) == 0
}

func main() {

}
