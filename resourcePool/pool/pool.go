package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	lock     sync.Mutex
	resource chan io.Closer
	factory  func() (closer io.Closer, err error)
	closed   bool
}

func New(fn func() (closer io.Closer, err error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size too small!")
	}
	return &Pool{
		resource: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		if !ok {
			return nil, errors.New("poll already closed!")
		}
		fmt.Println("acquire resource!")
		return r, nil
	default:
		fmt.Println("create resource!")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resource <- r:
	default:
		r.Close()
	}
}

func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed {
		return
	}
	close(p.resource)
	for r := range p.resource {
		r.Close()
	}
}
