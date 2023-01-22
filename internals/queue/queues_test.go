package queue

import (
	"testing"
)

func TestGetQueue(t *testing.T) {
	q := NewmemoryQueueStorage()
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
	q := NewmemoryQueueStorage()
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
	q := NewmemoryQueueStorage()
	_, err := q.CreateNewQueue("test")
	if err != nil {
		t.Error(err)
	}

	_, err = q.CreateNewQueue("test")
	if err == nil {
		t.Error("should not override a queue")
	}
}

func TestEnqueueQueueStorage(t *testing.T) {
	q := NewmemoryQueueStorage()
	if err := q.Enqueue("test", []byte("data")); err != nil {
		t.Error(err)
	}

	_, err := q.Dequeue("test")
	if err != nil {
		t.Error(err)
	}
}
func BenchmarkGetOrCreate(b *testing.B) {
	q := NewmemoryQueueStorage()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			q.GetQueueOrCreate("queue")
		}
	}
}

func BenchmarkEnqueueMemoryStorage(b *testing.B) {
	q := NewmemoryQueueStorage()
	data := []byte("data")

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			q.Enqueue("queue1", data)
		}
	}
}

func BenchmarkDequeueMemoryStorage(b *testing.B) {
	// Initialazation
	q := NewmemoryQueueStorage()
	n := 100
	data := []byte("data")
	for i := 0; i < n; i++ {
		q.Enqueue("queue1", data)
	}
	//
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			q.Dequeue("queue1")
		}
	}
}

// func TestPersistentGetQueue(t *testing.T) {
// 	defer func() { os.Remove("./test") }()
// 	q, err := NewPersistentQueueStorage("./test")
// 	assert.Nil(t, err)

// 	_, err = q.GetQueue("test")
// 	assert.NotNil(t, err)

// 	q.Enqueue("test", []byte("data"))
// 	_, err = q.GetQueue("test")
// 	assert.Nil(t, err)
// }

// func TestPersistentEnqueue(t *testing.T) {
// 	defer func() { os.Remove("./test") }()
// 	q, err := NewPersistentQueueStorage("./test")
// 	assert.Nil(t, err)

// 	err = q.Enqueue("test", []byte("data"))
// 	assert.Nil(t, err)

// 	_, err = q.Dequeue("test")
// 	assert.Nil(t, err)
// }

// func TestLoadFromMemory(t *testing.T) {
// 	defer func() { os.Remove("./test") }()
// 	q, err := NewPersistentQueueStorage("./test")
// 	assert.Nil(t, err)

// 	for i := 0; i < 10; i++ {
// 		err = q.Enqueue("test", []byte("data"))
// 		assert.Nil(t, err)
// 	}
// 	for i := 0; i < 4; i++ {
// 		_, err = q.Dequeue("test")
// 		assert.Nil(t, err)
// 	}
// 	q.walFd.Close()

// 	q, err = NewPersistentQueueStorage("./test")
// 	assert.Nil(t, err)
// 	data, err := q.Dequeue("test")
// 	assert.Nil(t, err)
// 	assert.Equal(t, data, []byte("data"))
// }
