package client

import (
	"encoding/json"
	"net"
	"fmt"
	"log"
	"bufio"
	"strings"
	"encoding/hex"
	"crypto/rand"
	"github.com/powerbenson/interview-homework/receiver"
	"github.com/powerbenson/interview-homework/identity"
)

type Client struct {
	host     string
	port     string
	conn     net.Conn
	receiver receiver.ClientReceiver
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) *Client {
	return &Client{
		host: config.Host,
		port: config.Port,
	}
}

func (client *Client) Run(clientName string, sendType string) {
	connection, _ := net.Dial("tcp", fmt.Sprintf("%s:%s", client.host, client.port))
	client.conn = connection

	person := identity.Identification {
		Name:     clientName,
		Token:    generateToken(),
		SendType: sendType,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("轉換成 JSON 失敗:", err)
		return
	}

	_, err = client.conn.Write(jsonData)
	if err != nil {
		log.Fatal("發送資料失敗:", err)
	}
	// fmt.Println("已向客戶端發送 JSON 數據")

	message, err := bufio.NewReader(client.conn).ReadString('\n')
	if (strings.TrimSpace(message) != "SUCESS") {
		log.Fatal("連線失敗:", err)
	}

	client.receiver = receiver.CreateReceiver(client.conn, person)
}

func (client *Client) Stop() {
	client.receiver.CloseConnection()
}

func (client *Client) Receive() {
	client.receiver.ReceiveMessage()
}

func generateToken() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}

	uuidStr := hex.EncodeToString(uuid)

	return uuidStr
}