package server

import (
	"bytes"
	"testing"

	"github.com/sina-am/queue-db/internals/queue"
)

func TestHandleCommand(t *testing.T) {
	queues := queue.NewmemoryQueueStorage()
	server := NewTCPServer(
		TCPServerOps{
			Addr:         ":8080",
			MaxQueueSize: 1000,
		},
		queues,
	)

	buffer := bytes.NewBuffer([]byte("enqueue q1 test"))
	if err := server.handleCommand(buffer); err != nil {
		t.Error(err)
	}

	data := make([]byte, 100)
	n, err := buffer.Read(data)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data[:n], []byte("ok")) {
		t.Errorf("expected %s got %s\n", "ok", data[:n])
	}

	buffer.Write([]byte("dequeue q1"))
	if err := server.handleCommand(buffer); err != nil {
		t.Error(err)
	}

	n, err = buffer.Read(data)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data[:n], []byte("test")) {
		t.Errorf("expected %s got %s\n", "test", data[:n])
	}
}
