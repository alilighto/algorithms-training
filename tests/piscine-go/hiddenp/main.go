package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	if s1 == "" {
		os.Stdout.WriteString("1")
		return
	}

	pos := 0

	for _, ch := range s2 {
		if pos < len(s1) && s1[pos] == byte(ch) {
			pos++
		}
	}

	if pos == len(s1) {
		os.Stdout.WriteString("1")
	} else {
		os.Stdout.WriteString("0")
	}
}
