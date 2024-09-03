package main

import (
	"fmt"
	"os"
)

func TrimSpaces(s string) string {
	if s == "" {
		return s
	}
	for _, v := range s {
		if v != ' ' {

		}
	}
}

func main() {
	args := os.Args[:1]
	var res []string
	if len(args) != 2 {

		res = Split(args)
		fmt.Println(res)

	}
}
