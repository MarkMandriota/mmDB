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
	r := &Req{Data: make(map[string]string)}

	if err := r.Load("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	r.Data["name"] = "Mark"
	r.Data["DOB"] = "2007"

	if err := r.Unload("info"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```
