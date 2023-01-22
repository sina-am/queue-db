package queue

import (
	"bytes"
	"testing"
)

func generateList() [][]byte {
	return [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	data_list := generateList()
	for _, item := range data_list {
		q.Enqueue(item)
	}

	for i, item := range data_list {
		data := q.Dequeue()
		if !bytes.Equal(data, data_list[i]) {
			t.Errorf("Expected %s, got %s", item, data)
		}
	}
}

func TestSearch(t *testing.T) {
	q := NewQueue()
	data_list := generateList()
	for _, item := range data_list {
		q.Enqueue(item)
	}

	for i := range data_list {
		if !q.Search(data_list[i]) {
			t.Errorf("Didn't find %s", data_list[i])
		}
	}
}
func TestGetSize(t *testing.T) {
	q := NewQueue()
	if q.GetSize() != 0 {
		t.Errorf("expected 0")
	}
	data_list := generateList()
	for _, item := range data_list {
		q.Enqueue(item)
	}

	if q.GetSize() != 4 {
		t.Errorf("excepted 4")
	}
}

func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue()
	data := []byte("data")

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			q.Enqueue(data)
		}
	}
}
