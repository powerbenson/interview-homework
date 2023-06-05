package pipe

import (
	"net"
	"log"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/fifo"
)

type PipieSender struct {
	Connection net.Conn
	ClientInfo identity.Identification
	// Fifo       *os.File
}

func(s *PipieSender) Init() {
	token := s.ClientInfo.Token

	fifo.DeleteFIFO(token)
	err := fifo.CreateFIFO(token)
	if err != nil {
		log.Fatal("創建命名管道失敗:", err)
	}
}

func(s *PipieSender) CloseConnection() {
	token := s.ClientInfo.Token

	pipe, err := fifo.OpenWriteFIFO(token)
	if err != nil {
		log.Fatal("無法開啟 FIFO:", err)
	}
	_, err = pipe.WriteString("Q" + "\n")
	if err != nil {
		log.Println("寫入 FIFO 失敗:", err)
	}

	fifo.DeleteFIFO(token)
	s.Connection.Close()
}

func(s *PipieSender) SendMessage(msg string) {
	token := s.ClientInfo.Token
	pipe, err := fifo.OpenWriteFIFO(token)
	if err != nil {
		log.Fatal("無法開啟 FIFO:", err)
	}
	_, err = pipe.WriteString(msg + "\n")
	if err != nil {
		log.Println("寫入 FIFO 失敗:", err)
	}
}