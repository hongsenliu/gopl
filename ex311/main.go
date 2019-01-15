package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func comma(s string) string {
	buf := &bytes.Buffer{}
	var suffix string
	var ii string
	l := len(s)
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		suffix = s[dot:]
		l = l - (l - dot)
		ii = s[:dot]
	}
	var p int
	if strings.HasPrefix(ii, "-") || strings.HasPrefix(ii, "+") {
		p = 1
		l--
		buf.WriteString(ii[:p])
		ii = s[p:len(ii)]
	}

	iil := len(ii)
	if iil < 3 {
		return s
	}
	pre := iil % 3

	if pre == 0 {
		pre = 3
	}
	buf.WriteString(ii[:pre])
	for i := pre; i < iil; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(ii[i : i+3])
	}
	buf.WriteString(suffix)
	return buf.String()
}
