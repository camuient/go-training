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
	n := len(s)

	if n < 3 {
		return s
	}

	var number int
	for i, v := n s {
		if number == 3 {
			buf.WriteString(",")
			number = 0
		}
		fmt.Fprintf(&buf, "%s", v)
		number++
	}

	return buf.String()
}
