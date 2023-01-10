package client

import (
	"net"

	"github.com/sina-am/queue-db/internals/command"
)

type clientConnection struct {
	net.Conn
}

type ClientConnection interface {
	net.Conn
	Enqueue(queueName string, data []byte) ([]byte, error)
	Dequeue(queueName string) ([]byte, error)
	SendMessage(cmd command.Message) ([]byte, error)
}

func (c clientConnection) SendMessage(cmd command.Message) ([]byte, error) {
	_, err := c.Write(command.EncodeMessage(cmd))
	if err != nil {
		return nil, err
	}

	response := make([]byte, 1024)
	n, err := c.Read(response)
	if err != nil {
		return nil, err
	}
	return response[:n], nil
}

func (c clientConnection) Dequeue(queueName string) ([]byte, error) {
	cmd := command.Message{
		Cmd:       command.Dequeue,
		QueueName: queueName,
	}
	return c.SendMessage(cmd)
}
func (c clientConnection) Enqueue(queueName string, data []byte) ([]byte, error) {
	cmd := command.Message{
		Cmd:       command.Enqueue,
		QueueName: queueName,
		Data:      data,
	}
	return c.SendMessage(cmd)
}

func Dial(address string) (ClientConnection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return clientConnection{}, err
	}
	return clientConnection{conn}, nil
}
