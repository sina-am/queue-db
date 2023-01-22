package queue

import (
	"testing"
)

func TestCreateNewNode(t *testing.T) {
	data := []byte("node1")
	firstNode := NewNode(data)
	if !firstNode.Contains(data) {
		t.Errorf("Expected %s got %s", data, firstNode.data)
	}
}
