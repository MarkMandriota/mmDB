# mmDB

**Example:**
```go
package main

import (
	. "./MM_database"
	"fmt"
	"os"
)

func main() {
	r := &Requester{Data: make(map[string]string)}

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	r.Load("info")

	r.Data["name"] = "Mark"
	r.Data["DOB"] = "11.01.2007"

	r.Unload("info")

	fmt.Printf(read("info"))
}

func read(pass string) string {
	r := &Requester{Data: make(map[string]string)}
	r.Load(pass)

	var str string
	for k, v := range r.Data {
		str += fmt.Sprintf("Key: %s, Value: %s\n", k, v)
	}

	return str
}
```
