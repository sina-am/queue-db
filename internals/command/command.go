package command

import (
	"bytes"
	"errors"
)

var ErrInvalidMessage = errors.New("invalid opration")

type Cmd string

const (
	Enqueue Cmd = "enqueue"
	Dequeue Cmd = "dequeue"
)

type Message struct {
	Cmd       Cmd
	QueueName string
	Data      []byte
}

func EncodeMessage(cmd Message) []byte {
	return bytes.Join(
		[][]byte{[]byte(cmd.Cmd), []byte(cmd.QueueName), cmd.Data},
		[]byte(" "))
}

func DecodeMessage(rowMessage []byte) (Message, error) {
	list := bytes.SplitN(rowMessage, []byte(" "), 3)
	if len(list) < 2 {
		return Message{}, ErrInvalidMessage
	}

	command := Message{
		Cmd:       Cmd(string(list[0])),
		QueueName: string(list[1]),
	}
	if command.Cmd == Dequeue {
		return command, nil
	}
	if command.Cmd == Enqueue && len(list) == 3 {
		command.Data = list[2]
		return command, nil
	}
	return command, ErrInvalidMessage
}
