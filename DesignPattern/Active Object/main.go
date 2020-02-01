package main

import "fmt"

type MethodRequest int

const (
	Incr MethodRequest = iota
	Decr
)

type Service struct {
	queue chan MethodRequest
	v     int
}

func (s *Service) Incr() {
	s.queue <- Incr
}

func (s *Service) Decr() {
	s.queue <- Decr
}

func (s *Service) schedule() {
	for r := range s.queue {
		if r == Incr {
			s.v++
		} else if r == Decr {
			s.v--
		}
	}
}

func New(buffer int) *Service {
	s := &Service{
		queue: make(chan MethodRequest, buffer),
	}
	go s.schedule()
	return s
}

func main() {
	s := New(0)
	s.Incr()
	s.Decr()
	s.Decr()
	fmt.Println(s.v)
}
