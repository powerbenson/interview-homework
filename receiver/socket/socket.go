package socket

import (
	"net"
	"bufio"
	"fmt"
	"github.com/powerbenson/interview-homework/math"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/util"
)

type SocketReceiver struct {
	Connection net.Conn
	ClientInfo identity.Identification
}

func(s *SocketReceiver) Init() {}

func(s *SocketReceiver) CloseConnection() {
	s.Connection.Close()
}

func(s *SocketReceiver) ReceiveMessage() {
	scanner := bufio.NewScanner(s.Connection)
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
		mean := math.CalculateMean(intArr)
		result := fmt.Sprintf("%s %g", "Mean is", mean)
		fmt.Println(result)
	}
}