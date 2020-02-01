package worker

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	wg       sync.WaitGroup
	taskPool chan Worker
}

func New(maxGoruntineNum int) *Pool {
	p := Pool{
		taskPool: make(chan Worker),
	}
	p.wg.Add(maxGoruntineNum)
	for i := 0; i < maxGoruntineNum; i++ {
		go func() {
			for task := range p.taskPool {
				task.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.taskPool <- w
}

func (p *Pool) Shutdown() {
	close(p.taskPool)
	p.wg.Wait()
}
