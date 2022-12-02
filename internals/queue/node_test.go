package queue

import (
	"testing"
)

func TestCreateNewNode(t *testing.T) {
	data := "node1"
	firstNode := CreateNewNode(data)
	if firstNode.data != data {
		t.Errorf("Expected %s got %s", data, firstNode.data)
	}
}
