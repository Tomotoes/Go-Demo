package queue

import "sync"

type Queue struct {
	head *Element
	tail *Element
	size int
	sync.Mutex
}

type Element struct {
	data []byte
	next *Element
}

func New() *Queue {
	return new(Queue)
}

func (Q *Queue) Push(data []byte) {
	Q.Lock()
	defer Q.Unlock()
	newElement := &Element{data: data}
	if Q.tail == nil {
		Q.head = newElement
	} else {
		Q.tail.next = newElement
	}

	Q.tail = newElement
	Q.size++
}

func (Q *Queue) Pop() []byte {
	Q.Lock()
	defer Q.Unlock()
	if Q.head == nil {
		return nil
	}
	data := Q.head.data
	Q.head = Q.head.next
	Q.size--
	return data
}

func (Q *Queue) Size() int {
	return Q.size

}
