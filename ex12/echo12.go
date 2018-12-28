// ex 1.2
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args[:] {
		fmt.Println(i, val)
	}
}
