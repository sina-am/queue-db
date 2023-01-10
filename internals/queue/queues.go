package queue

import (
	"errors"
	"fmt"
)

var ErrorQueueIsEmpty = errors.New("empty queue")

type QueueList interface {
	GetQueue(queueName string) (PriorityQueue, error)
	GetQueueOrCreate(name string) PriorityQueue
	CreateNewQueue(name string) (PriorityQueue, error)
	Enqueue(queueName string, data []byte) error
	Dequeue(queueName string) ([]byte, error)
	GetSize(queueName string) int
	// SearchInQueue(queueName string, data []byte) bool
}

type memoryQueueList struct {
	queues map[string]PriorityQueue
}

func NewmemoryQueueList() *memoryQueueList {
	return &memoryQueueList{
		queues: map[string]PriorityQueue{},
	}
}

func (q *memoryQueueList) GetQueue(name string) (PriorityQueue, error) {
	if queue, found := q.queues[name]; found {
		return queue, nil
	}
	return nil, ErrorQueueIsEmpty
}

func (q *memoryQueueList) GetQueueOrCreate(name string) PriorityQueue {
	queue, err := q.GetQueue(name)
	if err != nil {
		queue, err = q.CreateNewQueue(name)
		if err != nil {
			panic(err)
		}
		return queue
	}
	return queue
}

func (q *memoryQueueList) CreateNewQueue(name string) (PriorityQueue, error) {
	if _, found := q.queues[name]; found {
		return nil, fmt.Errorf("queue %s already exits", name)
	}
	q.queues[name] = NewQueue()
	return q.queues[name], nil
}

func (q *memoryQueueList) Enqueue(queueName string, data []byte) error {
	queue := q.GetQueueOrCreate(queueName)
	queue.Enqueue(data)
	return nil
}

func (q *memoryQueueList) Dequeue(queueName string) ([]byte, error) {
	queue, err := q.GetQueue(queueName)
	if err != nil {
		return nil, err
	}
	if queue.IsEmpty() {
		return nil, ErrorQueueIsEmpty
	}
	return queue.Dequeue().([]byte), nil
}

func (q *memoryQueueList) GetSize(queueName string) int {
	queue, err := q.GetQueue(queueName)
	if err != nil {
		return 0
	}
	return queue.GetSize()
}
