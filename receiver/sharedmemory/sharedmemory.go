package sharedmemory

import (
	"net"
	"unsafe"
	"fmt"
	"bytes"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/math"
	"github.com/powerbenson/interview-homework/util"
	"github.com/powerbenson/interview-homework/sharedmemory"
)

type SharedMemoryReceiver struct {
	Connection net.Conn
	ClientInfo identity.Identification
	Memoryblock uintptr
	Producer int
	Consumer int
}

func(s *SharedMemoryReceiver) Init() {
	memKey, _ := sharedmemory.GetMemKey(s.ClientInfo.Token)
	semProKey, _ := sharedmemory.GetProducerSemKey(s.ClientInfo.Token)
	semConKey, _ := sharedmemory.GetConsumerSemKey(s.ClientInfo.Token)

	addrMem, _ := sharedmemory.AttachMemoryBlock(memKey, 4096)
	semidPro := sharedmemory.Semget(semProKey, 1)
	semidCon := sharedmemory.Semget(semConKey, 1)

	s.Memoryblock = addrMem
	s.Producer = semidPro
	s.Consumer = semidCon
}

func(s *SharedMemoryReceiver) CloseConnection() {
	sharedmemory.DetachMemoryBlock(s.Memoryblock)
	s.Connection.Close()
}

func(s *SharedMemoryReceiver) ReceiveMessage() {
	for{
		sharedmemory.SemWait(s.Producer)

		ptr := unsafe.Pointer(s.Memoryblock)
		data := make([]byte, 100)
		copy(data, (*[1 << 30]byte)(ptr)[:len(data)])
		trimmedData := bytes.TrimRight(data, "\x00")
		msg := string(trimmedData)

		if (msg == "Q") {
			break
		}
		
		if !util.ValidateInput(msg) {
			break
		}
		intArr, err := util.ParseIntArray(msg)
		if err != nil {
			break
		}
		mode := math.CalculateMode(intArr)
		result := fmt.Sprintf("%s %v", "Mode is", mode)
		fmt.Println(result)

		sharedmemory.SemPost(s.Consumer)
	}
}