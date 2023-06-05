package fifo

import (
	"os"
	"syscall"
)

func CreateFIFO(key string) error {
	err := syscall.Mkfifo("/tmp/conn-file/fifo." + key, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFIFO(key string) error {
	err := os.Remove("/tmp/conn-file/fifo." + key)
	if err != nil {
		return err
	}
	return nil
}

func OpenWriteFIFO(key string) (*os.File, error) {
	fifo, err := os.OpenFile("/tmp/conn-file/fifo." + key, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}
	return fifo, nil
}

func OpenReadFIFO(key string) (*os.File, error) {
	fifo, err := os.OpenFile("/tmp/conn-file/fifo." + key, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}
	return fifo, nil
}
