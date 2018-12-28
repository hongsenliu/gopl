// ex 1.1
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, e := range os.Args[:] {
		s += sep + e
		sep = " "
	}
	fmt.Println(s)
}
