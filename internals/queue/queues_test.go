package queue

import (
	"testing"
)

func TestGetQueue(t *testing.T) {
	q := NewmemoryQueueList()
	_, err := q.GetQueue("test")
	if err == nil {
		t.Error("empty queue should return error")
	}

	q.Enqueue("test", []byte("data"))
	_, err = q.GetQueue("test")
	if err != nil {
		t.Error(err)
	}
}

func TestGetQueueOrCreate(t *testing.T) {
	q := NewmemoryQueueList()
	queue := q.GetQueueOrCreate("test")
	if queue == nil {
		t.Error("should create a queue")
	}

	err := q.Enqueue("test", []byte("data"))
	if err != nil {
		t.Error(err)
	}

	queue = q.GetQueueOrCreate("test")
	if queue.GetSize() != 1 {
		t.Error("wrong queue")
	}
}

func TestCreateNewQueue(t *testing.T) {
	q := NewmemoryQueueList()
	_, err := q.CreateNewQueue("test")
	if err != nil {
		t.Error(err)
	}

	_, err = q.CreateNewQueue("test")
	if err == nil {
		t.Error("should not override a queue")
	}
}

func TestEnqueueQueueList(t *testing.T) {
	q := NewmemoryQueueList()
	if err := q.Enqueue("test", []byte("data")); err != nil {
		t.Error(err)
	}

	_, err := q.Dequeue("test")
	if err != nil {
		t.Error(err)
	}
}

// func Dequeue(t *testing.T) {}
// func GetSize(t *testing.T) {}
