# mmDB

Example for using:
```go
package main

import (
	. "./MM_database"
	"fmt"
	"os"
)

func main() {
	r := Requester{Data: make(map[string]string)}

	if err := r.Load("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	r.Data["name"] = "Mark"
	r.Data["age"] = "13"

	if err := r.Unload("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```
