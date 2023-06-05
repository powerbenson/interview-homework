package sharedmemory

import (
	"net"
	"unsafe"
	"github.com/powerbenson/interview-homework/identity"
	"github.com/powerbenson/interview-homework/sharedmemory"
)

type SharedMemorySender struct {
	Connection  net.Conn
	ClientInfo  identity.Identification
	Memoryblock uintptr
	Producer	int
	Consumer	int
}

func(s *SharedMemorySender) Init() {
	sharedmemory.CreateMemKeyFile(s.ClientInfo.Token)
	sharedmemory.CreateProducerSemKeyFile(s.ClientInfo.Token)
	sharedmemory.CreateConsumerSemKeyFile(s.ClientInfo.Token)

	memKey, _ := sharedmemory.GetMemKey(s.ClientInfo.Token)
	semProKey, _ := sharedmemory.GetProducerSemKey(s.ClientInfo.Token)
	semConKey, _ := sharedmemory.GetConsumerSemKey(s.ClientInfo.Token)

	addrMem, _ := sharedmemory.AttachMemoryBlock(memKey, 4096)
	semidPro := sharedmemory.Semget(semProKey, 1)
	semidCon := sharedmemory.Semget(semConKey, 1)

	sharedmemory.SemPost(semidCon)

	s.Memoryblock = addrMem
	s.Producer = semidPro
	s.Consumer = semidCon
}

func(s *SharedMemorySender) CloseConnection() {
	sharedmemory.SemWait(s.Consumer)

	sharedmemory.ClearMemoryBlock(s.Memoryblock, 4096)
	data := []byte("Q")
	ptr := unsafe.Pointer(s.Memoryblock)
	copy((*[1 << 30]byte)(ptr)[:len(data)], data)

	sharedmemory.SemPost(s.Producer)

	sharedmemory.DetachMemoryBlock(s.Memoryblock)
	memKey, _ := sharedmemory.GetMemKey(s.ClientInfo.Token)
	sharedmemory.DestroyMemoryBlock(memKey, 4096)
	sharedmemory.RemoveMemKeyFile(s.ClientInfo.Token)

	sharedmemory.SemDestroy(s.Producer)
	sharedmemory.SemDestroy(s.Consumer)

	sharedmemory.RemoveProducerSemKeyFile(s.ClientInfo.Token)
	sharedmemory.RemoveConsumerSemKeyFile(s.ClientInfo.Token)
	
	s.Connection.Close()
}

func(s *SharedMemorySender) SendMessage(msg string) {
	sharedmemory.SemWait(s.Consumer)

	sharedmemory.ClearMemoryBlock(s.Memoryblock, 4096)
	data := []byte(msg)
	ptr := unsafe.Pointer(s.Memoryblock)
	copy((*[1 << 30]byte)(ptr)[:len(data)], data)

	sharedmemory.SemPost(s.Producer)
}