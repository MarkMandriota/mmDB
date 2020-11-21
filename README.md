# mmDB

**Example:**
```go
package main

import (
	. "MM_database"
	"fmt"
	"os"
)

func main() {
	r := &Req{Data: make(map[string]string)}

	if err := r.Load("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	r.Data["name"] = "Mark"
	r.Data["DOB"] = "2007"

	if err := r.Unload("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf(read("info"))
}

func read(pass string) string {
	r := &Req{Data: make(map[string]string)}
	r.Load(pass)

	var str string
	for k, v := range r.Data {
		str += fmt.Sprintf("Key: %s, Value: %s\n", k, v)
	}

	return str
}
```
