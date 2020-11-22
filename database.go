package database

import (
	"bufio"
	"io"
	"os"
	"syscall"
	"unsafe"
)

type Requester struct {
	Data map[string]string
}

func (r *Requester) Load(pass string) {
	file, err := os.OpenFile(pass, syscall.O_CREAT|syscall.O_RDONLY|syscall.O_CLOEXEC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		k, _ := reader.ReadBytes(0)
		v, err := reader.ReadBytes(0)
		if err == io.EOF {
			break
		}
		r.Data[(*(*string)(unsafe.Pointer(&k)))[:len(k)-1]] = (*(*string)(unsafe.Pointer(&v)))[:len(v)-1]
	}
}

func (r *Requester) Unload(pass string) {
	file, err := syscall.Open(pass, syscall.O_CREAT|syscall.O_TRUNC|syscall.O_WRONLY|syscall.O_CLOEXEC, 0777)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(file)

	buff := make([]byte, 0, 4096)
	for k, v := range r.Data {
		buff = append(buff, k + "\000" + v + "\000"...)
	}
	syscall.Write(file, buff)
}
