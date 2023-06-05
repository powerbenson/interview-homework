package sender

import (
	"net"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/sender/socket"
	"github.com/powerbenson/interview-homework/sender/pipe"
	"github.com/powerbenson/interview-homework/sender/sharedmemory"
)

var (
	clientSenders = make(map[string]ClientSender)
	// clientSendersLock sync.Mutex
)

type ClientSender interface {
	Init()
	CloseConnection()
	SendMessage(msg string)
}

func CreateSender(conn net.Conn, clientInfo identity.Identification) ClientSender {
	switch clientInfo.SendType {
		case "socket":
			sender := &socket.SocketSender {Connection: conn, ClientInfo: clientInfo}
			sender.Init()
			clientSenders[clientInfo.Token] = sender
			return sender
		case "pipe":
			sender := &pipe.PipieSender {Connection: conn, ClientInfo: clientInfo}
			sender.Init()
			clientSenders[clientInfo.Token] = sender
			return sender
		case "shared-memory":
			sender := &sharedmemory.SharedMemorySender {Connection: conn, ClientInfo: clientInfo}
			sender.Init()
			clientSenders[clientInfo.Token] = sender
			return sender
	}
	return nil
}

func GetClientSenders() map[string]ClientSender {
	return clientSenders
}