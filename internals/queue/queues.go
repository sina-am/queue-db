package queue

import (
	"errors"
	"fmt"
)

var ErrorQueueIsEmpty = errors.New("empty queue")

type QueueStorage interface {
	GetQueue(queueName string) (PriorityQueue, error)
	GetQueueOrCreate(name string) PriorityQueue
	CreateNewQueue(name string) (PriorityQueue, error)
	Enqueue(queueName string, data []byte) error
	Dequeue(queueName string) ([]byte, error)
	GetQueueSize(queueName string) int
	GetMemorySize() int
	// SearchInQueue(queueName string, data []byte) bool
}

type memoryQueueStorage struct {
	queues     map[string]PriorityQueue
	memorySize int
}

func NewmemoryQueueStorage() *memoryQueueStorage {
	return &memoryQueueStorage{
		queues: map[string]PriorityQueue{},
	}
}

func (q *memoryQueueStorage) GetQueue(name string) (PriorityQueue, error) {
	if queue, found := q.queues[name]; found {
		return queue, nil
	}
	return nil, ErrorQueueIsEmpty
}

func (q *memoryQueueStorage) GetQueueOrCreate(name string) PriorityQueue {
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

func (q *memoryQueueStorage) CreateNewQueue(name string) (PriorityQueue, error) {
	if _, found := q.queues[name]; found {
		return nil, fmt.Errorf("queue %s already exits", name)
	}
	q.queues[name] = NewQueue()
	return q.queues[name], nil
}

func (q *memoryQueueStorage) Enqueue(queueName string, data []byte) error {
	queue := q.GetQueueOrCreate(queueName)
	queue.Enqueue(data)
	q.memorySize += len(data)
	return nil
}

func (q *memoryQueueStorage) Dequeue(queueName string) ([]byte, error) {
	queue, err := q.GetQueue(queueName)
	if err != nil {
		return nil, err
	}
	if queue.IsEmpty() {
		return nil, ErrorQueueIsEmpty
	}
	if data, ok := queue.Dequeue().([]byte); ok {
		q.memorySize -= len(data)
		return data, nil
	}
	return nil, errors.New("can't")
}

func (q *memoryQueueStorage) GetQueueSize(queueName string) int {
	queue, err := q.GetQueue(queueName)
	if err != nil {
		return 0
	}
	return queue.GetSize()
}

func (q *memoryQueueStorage) GetMemorySize() int {
	return q.memorySize
}
