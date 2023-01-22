package queue

import (
	"errors"
	"fmt"
	"sync"
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
	memorySize int // Used Memory
	lock       sync.RWMutex
}

func NewMemoryQueueStorage() *memoryQueueStorage {
	return &memoryQueueStorage{
		queues:     map[string]PriorityQueue{},
		memorySize: 0,
		lock:       sync.RWMutex{},
	}
}

func (q *memoryQueueStorage) GetQueue(name string) (PriorityQueue, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

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
	q.lock.Lock()
	defer q.lock.Unlock()

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
	data := queue.Dequeue()
	q.memorySize -= len(data)
	return data, nil
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

// type persistentQueueStorage struct {
// 	walFd      *os.File // Write ahead log
// 	memStorage *memoryQueueStorage
// }

// func fileExists(filePath string) (bool, error) {
// 	info, err := os.Stat(filePath)
// 	if err == nil {
// 		return !info.IsDir(), nil
// 	}
// 	if errors.Is(err, os.ErrNotExist) {
// 		return false, nil
// 	}
// 	return false, err
// }

// func readIntoMemory(filePath string, memStorage *memoryQueueStorage) error {
// 	fd, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer fd.Close()

// 	fileScanner := bufio.NewScanner(fd)
// 	fileScanner.Split(bufio.ScanLines)

// 	for fileScanner.Scan() {
// 		msg, err := command.DecodeMessage(fileScanner.Bytes())
// 		if err != nil {
// 			return err
// 		}
// 		switch msg.Cmd {
// 		case command.Enqueue:
// 			err = memStorage.Enqueue(msg.QueueName, msg.Data)
// 			if err != nil {
// 				return err
// 			}
// 		case command.Dequeue:
// 			_, err = memStorage.Dequeue(msg.QueueName)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func NewPersistentQueueStorage(filePath string) (*persistentQueueStorage, error) {
// 	memStorage := NewMemoryQueueStorage()
// 	exist, err := fileExists(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Read from wal
// 	if exist {
// 		if err := readIntoMemory(filePath, memStorage); err != nil {
// 			return nil, err
// 		}
// 	}

// 	fd, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &persistentQueueStorage{
// 		walFd:      fd,
// 		memStorage: memStorage,
// 	}, nil
// }

// func (q *persistentQueueStorage) GetQueue(name string) (PriorityQueue, error) {
// 	return q.memStorage.GetQueue(name)
// }

// func (q *persistentQueueStorage) GetQueueOrCreate(name string) PriorityQueue {
// 	return q.memStorage.GetQueueOrCreate(name)
// }

// func (q *persistentQueueStorage) CreateNewQueue(name string) (PriorityQueue, error) {
// 	return q.memStorage.CreateNewQueue(name)
// }

// func (q *persistentQueueStorage) Enqueue(queueName string, data []byte) error {
// 	byteMsg := command.EncodeMessage(
// 		command.Message{
// 			Cmd:       command.Enqueue,
// 			QueueName: queueName,
// 			Data:      data,
// 		},
// 	)
// 	_, err := q.walFd.Write(byteMsg)
// 	if err != nil {
// 		return err
// 	}
// 	return q.memStorage.Enqueue(queueName, data)
// }

// func (q *persistentQueueStorage) Dequeue(queueName string) ([]byte, error) {
// 	byteMsg := command.EncodeMessage(
// 		command.Message{
// 			Cmd:       command.Dequeue,
// 			QueueName: queueName,
// 		},
// 	)
// 	_, err := q.walFd.Write(byteMsg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return q.memStorage.Dequeue(queueName)
// }

// func (q *persistentQueueStorage) GetQueueSize(queueName string) int {
// 	return q.memStorage.GetQueueSize(queueName)
// }

// func (q *persistentQueueStorage) GetMemorySize() int {
// 	return q.memStorage.GetMemorySize()
// }
