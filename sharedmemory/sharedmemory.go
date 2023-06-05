package sharedmemory

import (
	"fmt"
	"os"
	"syscall"
	"log"
	"unsafe"
	"github.com/hslam/ftok"
)

/*
#include <sys/sem.h>
typedef struct sembuf sembuf;
typedef union semun semun;
*/
import "C"

func CreateMemKeyFile(key string) {
	file, err := os.Create("/tmp/conn-file/mem." + key)
	if err != nil {
		fmt.Printf("無法創建檔案：%v\n", err)
		return
	}
	defer file.Close()
}

func RemoveMemKeyFile(key string) {
	err := os.Remove("/tmp/conn-file/mem." + key)
	if err != nil {
		fmt.Printf("無法刪除檔案：%v\n", err)
		return
	}
}

func GetMemKey(path string) (int, error) {
	key, err := ftok.Ftok("/tmp/conn-file/mem." + path, 0x22)
	if err != nil {
		return 0, fmt.Errorf("無法生成鍵值：%v", err)
	}
	return key, nil
}

func CreateProducerSemKeyFile(key string) {
	file, err := os.Create("/tmp/conn-file/pro." + key)
	if err != nil {
		fmt.Printf("無法創建檔案：%v\n", err)
		return
	}
	defer file.Close()
}

func RemoveProducerSemKeyFile(key string) {
	err := os.Remove("/tmp/conn-file/pro." + key)
	if err != nil {
		fmt.Printf("無法刪除檔案：%v\n", err)
		return
	}
}

func GetProducerSemKey(path string) (int, error) {
	key, err := ftok.Ftok("/tmp/conn-file/pro." + path, 0x22)
	if err != nil {
		return 0, fmt.Errorf("無法生成鍵值：%v", err)
	}
	return key, nil
}

func CreateConsumerSemKeyFile(key string) {
	file, err := os.Create("/tmp/conn-file/con." + key)
	if err != nil {
		fmt.Printf("無法創建檔案：%v\n", err)
		return
	}
	defer file.Close()
}

func RemoveConsumerSemKeyFile(key string) {
	err := os.Remove("/tmp/conn-file/con." + key)
	if err != nil {
		fmt.Printf("無法刪除檔案：%v\n", err)
		return
	}
}

func GetConsumerSemKey(path string) (int, error) {
	key, err := ftok.Ftok("/tmp/conn-file/con." + path, 0x22)
	if err != nil {
		return 0, fmt.Errorf("無法生成鍵值：%v", err)
	}
	return key, nil
}

func GetSharedBlock(key, size int) (uintptr, error) {
	// 暫且固定
	perm := 0666

	// 使用 SYS_SHMGET 系統呼叫獲取共享記憶體區域
	shmid, _, err := syscall.Syscall(
		syscall.SYS_SHMGET,
		uintptr(key),
		uintptr(size),
		uintptr(C.IPC_CREAT|perm),
	)
	if err != 0 {
		return 0, fmt.Errorf("無法獲取共享記憶體區域：%v", err)
	}

	return shmid, nil
}

func AttachMemoryBlock(key, size int) (uintptr, error) {
	shmid, err := GetSharedBlock(key, size)
	if err != nil {
		return 0, err
	}

	// 使用 SYS_SHMAT 系統呼叫附加到共享記憶體區域
	addr, _, err := syscall.Syscall(
		syscall.SYS_SHMAT,
		shmid,
		0,
		0,
	)
	if addr == ^uintptr(0) {
		return 0, fmt.Errorf("無法附加到共享記憶體區域：%v", err)
	}

	return addr, nil
}

func DetachMemoryBlock(addr uintptr) error {
	_, _, err := syscall.Syscall(
		syscall.SYS_SHMDT,
		addr,
		0,
		0,
	)
	if err != 0 {
		return fmt.Errorf("無法分離共享記憶體區域：%v", err)
	}

	return nil
}

func DestroyMemoryBlock(key, size int) error {
	shmid, err := GetSharedBlock(key, size)
	if err != nil {
		return err
	}

	_, _, err = syscall.Syscall(
		syscall.SYS_SHMCTL,
		shmid,
		uintptr(C.IPC_RMID),
		0,
	)
	if err != nil {
		return fmt.Errorf("無法銷毀共享記憶體區域：%v", err)
	}

	return nil
}

func ClearMemoryBlock(addr uintptr, size int) {
	data := make([]byte, size)
	copy((*[1 << 30]byte)(unsafe.Pointer(addr))[:size], data)
}

func Semget(key int, initCount int) int {
	r1, r2, err := syscall.Syscall(syscall.SYS_SEMGET, uintptr(key),
		uintptr(initCount), uintptr(00666))
	if int(r1) < 0 {
		r1, r2, err = syscall.Syscall(syscall.SYS_SEMGET, uintptr(key),
			uintptr(initCount), uintptr(C.IPC_CREAT|C.IPC_EXCL|00666))
		if int(r1) < 0 {
			log.Printf("error:semget error is %v\n", err)
		}
	} else {
		log.Printf("success :semget is %v,%v,%v\n", r1, r2, err)
	}
	return int(r1)
}

func SemWait(semid int) int {

	stSemBuf := C.sembuf{
		sem_num: 0,
		sem_op:  -1,
		sem_flg: 0,
	}

	r1, r2, err := syscall.Syscall(syscall.SYS_SEMOP, uintptr(semid), uintptr(unsafe.Pointer(&stSemBuf)), 1)
	if int(r1) < 0 {
		log.Printf("error:semget error is %v,%v,%v\n", r1, r2, err)
	}
	return int(r1)
}

func SemPost(semid int) int {

	stSemBuf := C.sembuf{
		sem_num: 0,
		sem_op:  1,
		sem_flg: 0,
	}

	r1, r2, err := syscall.Syscall(syscall.SYS_SEMOP, uintptr(semid), uintptr(unsafe.Pointer(&stSemBuf)), 1)
	if int(r1) < 0 {
		log.Printf("error:semget error is %v,%v,%v\n", r1, r2, err)
	}
	return int(r1)
}

func SemDestroy(semid int) error {
	_, _, err := syscall.Syscall(syscall.SYS_SEMCTL, uintptr(semid), 0, uintptr(C.IPC_RMID))
	if err != 0 {
		return err
	}
	return nil
}

func GetSemaphoreValue(semid int) (int, error) {
	arg := C.IPC_STAT
	value, _, err := syscall.Syscall6(syscall.SYS_SEMCTL, uintptr(semid), 0, uintptr(arg), 0, 0, 0)
	if err != 0 {
		return 0, err
	}
	return int(value), nil
}