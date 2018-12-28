// ex13
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func m1() {
	var s, sep string
	for _, val := range os.Args[:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
}

func m2() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

func main() {
	start := time.Now()
	m1()
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	m2()
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())
}
