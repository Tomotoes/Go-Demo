package autoinc

type Autoinc struct {
	start, step int
	queue       chan int
	quit        chan bool
}

func New(start, step int) *Autoinc {
	ai := &Autoinc{
		start: start,
		step:  step,
		queue: make(chan int, 4),
		quit:  make(chan bool),
	}
	go ai.process()
	return ai
}

func (ai *Autoinc) process() {
	defer func() {
		recover()
		close(ai.queue)
	}()

	for i := ai.start; ; i += ai.step {
		select {
		case ai.queue <- i:
		case <-ai.quit:
			break
		}
	}
}

func (ai *Autoinc) Id() int {
	return <-ai.queue
}

func (ai *Autoinc) Close() {
	close(ai.quit)
}
