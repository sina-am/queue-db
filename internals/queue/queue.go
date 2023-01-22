package queue

import "sync"

type PriorityQueue interface {
	Enqueue(data NodeData)
	Dequeue() NodeData
	Search(data NodeData) bool
	IsEmpty() bool
	GetSize() int
}
type queue struct {
	lock sync.RWMutex
	tail *Node
	head *Node
	size int
}

func NewQueue() *queue {
	q := &queue{
		lock: sync.RWMutex{},
		size: 0,
	}

	firstNode := NewNode(nil)
	q.head = firstNode
	q.tail = firstNode
	return q
}

func (q *queue) IsEmpty() bool {
	return q.head.data == nil
}

func (q *queue) Enqueue(data NodeData) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.tail.data = data
	new_tail := NewNode(nil)
	new_tail.next = q.tail
	q.tail.prev = new_tail
	q.tail = new_tail
	q.size++
}

func (q *queue) Dequeue() NodeData {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.IsEmpty() {
		return nil
	}
	data := q.head.data
	q.head = q.head.prev
	q.head.next = nil
	q.size--
	return data
}

func (q *queue) Search(data NodeData) bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	for cur := q.head; cur != nil; cur = cur.prev {
		if cur.Contains(data) {
			return true
		}
	}
	return false
}

func (q *queue) GetSize() int {
	return q.size
}
