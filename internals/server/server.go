package server

import (
	"log"
	"net"

	"github.com/sina-am/queue-db/internals/queue"

	"github.com/sina-am/queue-db/internals/command"
)

func DecodeRecvBytes(conn net.Conn) (command.Command, error) {
	commandData := make([]byte, 1024)
	n, err := conn.Read(commandData)
	if err != nil {
		return command.Command{}, err
	}
	return command.DecodeCommand(commandData[:n])
}

func HandleConnection(conn TCPConnection) {
	for {
		cmd, err := DecodeRecvBytes(conn)
		if err != nil {
			if err == command.ErrorInvalidOpration {
				conn.WriteError("Sorry, don't know how to read this.")
				continue
			} else {
				log.Printf("Connection closed %v", conn.RemoteAddr())
				conn.Close()
				break
			}
		}
		response, err := command.RunCommand(cmd)
		if err != nil {
			conn.WriteError(err.Error())
		}
		conn.WriteData(response)
	}
}

func WaitForConnections(listener net.Listener) {
	// TODO: Should set a max_connection parameter
	for {
		conn, err := listener.Accept()
		log.Printf("Connection accepted from %v", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		tcpConn := tcpConnection{conn}
		go HandleConnection(tcpConn)
	}
}

func RunServer(address string) {
	queue.InitiateMemoryQueues()
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	WaitForConnections(listener)
}
