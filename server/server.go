package server

import (
	"encoding/json"
	"net"
	"log"
	"fmt"
	"time"
	"github.com/powerbenson/interview-homework/sender"
	"github.com/powerbenson/interview-homework/identity"
)


type Server struct {
	host     string
	port     string
	listener net.Listener
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server) Run() {
	lsn, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	server.listener = lsn
	if err != nil {
		log.Fatal(err)
	}
	defer server.listener.Close()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

func (server *Server) Stop() {
	for _, value := range server.GetClientSenders() {
		go value.CloseConnection()
	}
	time.Sleep(2 * time.Second)
	server.listener.Close()
}

func (server *Server) SendMessage(input string) {
	for _, value := range server.GetClientSenders() {
		go value.SendMessage(input)
	}
}

func handleRequest(conn net.Conn) {
	jsonData := make([]byte, 1024) // 假設 JSON 數據大小不超過 1024 字節

	n, err := conn.Read(jsonData)
	if err != nil {
		log.Fatal("讀取數據失敗:", err)
	}

	var clientInfo identity.Identification
	err = json.Unmarshal(jsonData[:n], &clientInfo)
	if err != nil {
		conn.Close()
		log.Fatal("JSON 解析失敗:", err)
	}

	if (!identityVerify(clientInfo)) {
		conn.Close()
		log.Fatal("身份認證失敗:", err)
	}

	sender.CreateSender(conn, clientInfo)
	fmt.Fprintf(conn, "SUCESS" + "\n")
	fmt.Printf("%s is ready\n", clientInfo.Name)
}

func identityVerify(clientInfo identity.Identification) bool {
	// 如果 Token 正確回傳 true，失敗回傳 false，這邊沒有實作，暫時都回傳 true
	return true
}

func (server *Server) GetClientSenders() map[string]sender.ClientSender {
	return sender.GetClientSenders()
}