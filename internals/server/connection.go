package server

import (
	"net"
)

type tcpConnection struct {
	net.Conn
}

type TCPConnection interface {
	net.Conn
	WriteData(data []byte) (int, error)
	WriteError(message string) (int, error)
}

func (c tcpConnection) WriteData(data []byte) (int, error) {
	return c.Write(data)
}

func (c tcpConnection) WriteError(message string) (int, error) {
	return c.Write([]byte("{\"error\": " + message + "}"))
}
