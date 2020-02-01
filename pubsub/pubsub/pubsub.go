package pubsub

import (
	"sync"
	"time"
)

type (
	Topic interface{}
	Subscribe chan Topic
	TopicFilter func(Topic) bool
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[Subscribe]TopicFilter
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[Subscribe]TopicFilter),
	}
}

func (p *Publisher) SubscribeTopic(topicFilter TopicFilter) Subscribe {
	ch := make(Subscribe, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topicFilter
	p.m.Unlock()
	return ch
}

func (p *Publisher) SubscribeAll() Subscribe {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) Evict(sub Subscribe) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) Publish(topic Topic) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, filter := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, filter, topic, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(sub Subscribe, filter TopicFilter, topic Topic, wg *sync.WaitGroup) {
	defer wg.Done()
	if filter != nil && !filter(topic) {
		return
	}

	select {
	case sub <- topic:
	case <-time.After(p.timeout):
	}
}
