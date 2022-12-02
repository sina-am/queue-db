package queue

import (
	"testing"
)

func generateList() []string {
	return []string{"1", "2", "3", "4"}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	data_list := generateList()
	for _, item := range data_list {
		q.Enqueue(item)
	}

	for i, item := range data_list {
		data := q.Dequeue()
		if data != data_list[i] {
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
