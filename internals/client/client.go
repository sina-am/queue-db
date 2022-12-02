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
	SendCommand(cmd command.Command) ([]byte, error)
}

func (c clientConnection) SendCommand(cmd command.Command) ([]byte, error) {
	_, err := c.Write(command.EncodeCommand(cmd))
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
	cmd := command.Command{
		Opration:  command.Dequeue,
		QueueName: queueName,
	}
	return c.SendCommand(cmd)
}
func (c clientConnection) Enqueue(queueName string, data []byte) ([]byte, error) {
	cmd := command.Command{
		Opration:  command.Enqueue,
		QueueName: queueName,
		Data:      data,
	}
	return c.SendCommand(cmd)
}

func Dial(address string) (ClientConnection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return clientConnection{}, err
	}
	return clientConnection{conn}, nil
}
