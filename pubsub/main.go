package main

import (
	"fmt"
	"other/pubsub/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	all := p.SubscribeAll()
	goland := p.SubscribeTopic(func(topic pubsub.Topic) bool {
		if s, ok := topic.(string); ok {
			return strings.Contains(s, "goland")
		}
		return false
	})

	p.Publish("hello, world")
	p.Publish("hello, goland")

	go func() {
		for topic := range all {
			fmt.Println(topic)
		}
	}()

	go func() {
		for topic := range goland {
			fmt.Println(topic)
		}
	}()

	time.Sleep(time.Second * 3)
}
