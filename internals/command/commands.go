package command

import (
	"errors"
	"log"

	"github.com/sina-am/queue-db/internals/queue"
)

func RunEnqueueCommand(cmd Command) ([]byte, error) {
	queue := queue.GetQueueOrCreate(cmd.QueueName)
	queue.Enqueue(cmd.Data)
	return []byte("ok"), nil
}

func RunDequeueCommand(cmd Command) ([]byte, error) {
	queue, err := queue.GetQueue(cmd.QueueName)
	if err != nil {
		return nil, err
	}
	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return queue.Dequeue().([]byte), nil
}

func RunCommand(cmd Command) ([]byte, error) {
	if cmd.Opration == Enqueue {
		return RunEnqueueCommand(cmd)
	}
	if cmd.Opration == Dequeue {
		return RunDequeueCommand(cmd)
	}
	log.Printf("unexpected error")
	return nil, errors.New("unxpected error")
}
