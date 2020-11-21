package mm_database

import (
	"bufio"
	"io"
	"os"
	"syscall"
)

type (
	Requester interface {
		Load(name string) error
		Unload(name string) error
	}
	
	Req struct {
		Data map[string]string
	}
)	

func (r *Req) Load(name string) error {
	file, err := os.OpenFile(name, os.O_CREATE | os.O_RDONLY | syscall.O_CLOEXEC, 0777)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for {
		k, _ := reader.ReadBytes(0)
		v, err := reader.ReadBytes(0)
		if err == io.EOF {
			break
		}

		r.Data[string(k[:len(k)-1])] = string(v[:len(v)-1])
	}

	return nil
}

func (r *Req) Unload(name string) error {
	file, err := syscall.Open(name, os.O_CREATE | os.O_TRUNC | syscall.O_WRONLY | syscall.O_CLOEXEC, 0777)
	if err != nil {
		return err
	}

	buff := make([]byte, 0, 1024)
	for k, v := range r.Data {
		buff = append(buff, []byte(k + "\000" + v + "\000")...)
	}
	syscall.Write(file, buff)

	return nil
}
