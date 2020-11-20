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

func (r *Requester) Load(name string) error {
	file, err := os.Open(name)
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

func (r *Requester) Unload(name string) error {
	path, _ := syscall.UTF16PtrFromString(name)
	syscall.DeleteFile(path)

	file, err := syscall.Open(name, syscall.O_CREAT | syscall.O_WRONLY | syscall.O_CLOEXEC | syscall.O_ASYNC, 0)
	if err != nil {
		return err
	}

	for k, v := range r.Data {
		syscall.Write(file, []byte(k + "\000" + v + "\000"))
	}

	return nil
}
