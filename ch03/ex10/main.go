package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	slice := []byte(s)
	n := len(s)

	if n <= 3 {
		return s
	}

	number := 0
	for i := 0; i <= n-1; i++ {
		if number == 3 {
			buf.WriteString(",")
			number = 0
		}
		buf.WriteByte(slice[i])
		number++
	}

	return buf.String()
}
