package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma2("123456789"))
}

func comma2(s string) string {
	buf := &bytes.Buffer{}
	l := len(s)
	if l < 3 {
		return s
	}
	pre := l % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < l; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i:(i + 3)])

	}
	return buf.String()
}
