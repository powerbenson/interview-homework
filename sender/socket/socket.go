package socket

import (
	"net"
	"fmt"
	"github.com/powerbenson/interview-homework/identity"
)

type SocketSender struct {
	Connection net.Conn
	ClientInfo identity.Identification
}

func(s *SocketSender) Init() {}

func(s *SocketSender) CloseConnection() {
	fmt.Fprintf(s.Connection, "Q") 
	s.Connection.Close()
}

func(s *SocketSender) SendMessage(msg string) {
	fmt.Fprintf(s.Connection, msg + "\n") 
}