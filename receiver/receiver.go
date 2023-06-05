package receiver

import (
	"net"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/receiver/socket"
	"github.com/powerbenson/interview-homework/receiver/pipe"
	"github.com/powerbenson/interview-homework/receiver/sharedmemory"
)

type ClientReceiver interface {
	Init()
	CloseConnection()
	ReceiveMessage()
}

func CreateReceiver(conn net.Conn, clientInfo identity.Identification) ClientReceiver {
	switch clientInfo.SendType {
		case "socket":
			receiver := &socket.SocketReceiver {Connection: conn, ClientInfo: clientInfo}
			receiver.Init()
			return receiver
		case "pipe":
			receiver := &pipe.PipeReceiver {Connection: conn, ClientInfo: clientInfo}
			receiver.Init()
			return receiver
		case "shared-memory":
			receiver := &sharedmemory.SharedMemoryReceiver {Connection: conn, ClientInfo: clientInfo}
			receiver.Init()
			return receiver
	}
	return nil
}