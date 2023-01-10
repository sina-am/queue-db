package server

import (
	"fmt"
	"log"
	"net"

	"github.com/sina-am/queue-db/internals/queue"

	"github.com/sina-am/queue-db/internals/command"
)

type TCPServerOps struct {
	Addr string
}

type TCPServer struct {
	TCPServerOps
	queues queue.QueueList
}

func NewTCPServer(ops TCPServerOps, queues queue.QueueList) *TCPServer {
	return &TCPServer{
		TCPServerOps: ops,
		queues:       queues,
	}
}

func (s *TCPServer) Run() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Printf("Server is running on %s", s.Addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error %s", conn)
			continue
		}
		log.Printf("Connection accepted from %v", conn.RemoteAddr())
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// TODO: better error handling
		if err := s.handleCommand(conn); err != nil {
			if err == command.ErrInvalidMessage || err == queue.ErrorQueueIsEmpty {
				conn.Write([]byte(err.Error()))
				continue
			}
			log.Printf("%s\nConnection closed %s", err.Error(), conn.RemoteAddr())
			break
		}
	}
}

func (s *TCPServer) handleCommand(conn net.Conn) error {
	rowMessage := make([]byte, 1024)
	n, err := conn.Read(rowMessage)
	if err != nil {
		return err
	}
	msg, err := command.DecodeMessage(rowMessage[:n])
	if err != nil {
		return err
	}
	switch msg.Cmd {
	case command.Enqueue:
		err := s.queues.Enqueue(msg.QueueName, msg.Data)
		if err != nil {
			return err
		}
		_, err = conn.Write([]byte("ok"))
		return err
	case command.Dequeue:
		data, err := s.queues.Dequeue(msg.QueueName)
		if err != nil {
			return err
		}
		_, err = conn.Write(data)
		return err
	default:
		return fmt.Errorf("invalid messsage command %s", err)
	}
}
