package database

import (
	"bufio"
	"io"
	"os"
	. "syscall"
	"unsafe"
)

type Requester struct {
	Data map[string]string
}

func (r *Requester) Load(pass string) {
	file, err := os.OpenFile(pass, O_CREAT|O_RDONLY|O_CLOEXEC, 0777)
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
	file, err := os.OpenFile(pass, O_CREAT|O_TRUNC|O_WRONLY|O_CLOEXEC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for k, v := range r.Data {
		writer.WriteString(k + "\000" + v + "\000")
	}
	writer.Flush()
}
