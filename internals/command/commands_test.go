package command

import (
	"bytes"
	"testing"

	"github.com/sina-am/queue-db/internals/queue"
)

func TestRunEnqueueCommand(t *testing.T) {
	queue.InitiateMemoryQueues()

	cmd := Command{
		Opration:  Enqueue,
		QueueName: "test",
		Data:      []byte("data"),
	}
	ok, err := RunEnqueueCommand(cmd)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(ok, []byte("ok")) {
		t.Error("invalid response")
	}
}

func TestRunDequeueCommand(t *testing.T) {
	queue.InitiateMemoryQueues()

	cmd1 := Command{
		Opration:  Enqueue,
		QueueName: "test",
		Data:      []byte("data"),
	}
	RunEnqueueCommand(cmd1)

	cmd2 := Command{
		Opration:  Dequeue,
		QueueName: "test",
	}

	data, err := RunDequeueCommand(cmd2)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, cmd1.Data) {
		t.Error("invalid")
	}
}
