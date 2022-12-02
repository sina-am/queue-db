package queue

import "errors"

var ErrorQueueNotFound = errors.New("no such queue")
var queues map[string]PriorityQueue

func GetQueue(name string) (PriorityQueue, error) {
	if queue, found := queues[name]; found {
		return queue, nil
	}
	return nil, ErrorQueueNotFound
}

func GetQueueOrCreate(name string) PriorityQueue {
	queue, err := GetQueue(name)
	if err != nil {
		return CreateNewQueue(name)
	}
	return queue
}

func CreateNewQueue(name string) PriorityQueue {
	queues[name] = NewQueue()
	return queues[name]
}

func InitiateMemoryQueues() {
	queues = make(map[string]PriorityQueue)
}
