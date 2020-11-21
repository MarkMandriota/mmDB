package mm_database

import (
	"bufio"
	"io"
	"os"
	"syscall"
)

type Requester struct {
	Data map[string]string
}

func (r *Requester) Load(pass string) {
	file, err := os.OpenFile(pass, os.O_CREATE|os.O_RDONLY|syscall.O_CLOEXEC, 0777)
	if err != nil {
		panic(err)
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
}

func (r *Requester) Unload(pass string) {
	file, err := syscall.Open(pass, os.O_CREATE | os.O_TRUNC | syscall.O_WRONLY | syscall.O_CLOEXEC, 0777)
	if err != nil {
		panic(err)
	}

	buff := make([]byte, 0, 1024)
	for k, v := range r.Data {
		buff = append(buff, []byte(k + "\000" + v + "\000")...)
	}
	syscall.Write(file, buff)
}
