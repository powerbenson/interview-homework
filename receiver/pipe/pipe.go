package pipe

import (
	"net"
	"os"
	"log"
	"fmt"
	"bufio"
	"github.com/powerbenson/interview-homework/math"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/util"
	"github.com/powerbenson/interview-homework/fifo"
)

type PipeReceiver struct {
	Connection net.Conn
	ClientInfo identity.Identification
	Fifo       *os.File
}

func(s *PipeReceiver) Init() {
	token := s.ClientInfo.Token
	fifo, err := fifo.OpenReadFIFO(token)
	if err != nil {
		log.Fatal("無法開啟 FIFO:", err)
	}
	s.Fifo = fifo
}

func(s *PipeReceiver) CloseConnection() {
	s.Connection.Close()
}

func(s *PipeReceiver) ReceiveMessage() {
	scanner := bufio.NewScanner(s.Fifo)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "Q" {
			return
		}

		if !util.ValidateInput(input) {
			continue
		}

		intArr, err := util.ParseIntArray(input)
		if err != nil {
			continue
		}

		median := math.CalculateMedian(intArr)
		result := fmt.Sprintf("%s %g", "Median is", median)
		fmt.Println(result)
	}
}