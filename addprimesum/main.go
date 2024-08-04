package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	n := 0
	res := ""
	if len(args) == 1 {
		pvPrime := PrePrime(args)
		for _, v := range pvPrime {
			n += itoa(pvPrime)
		}
		res = itoa(n)
		os.Stdout.WriteString(res)
	}
	os.Stdout.WriteString("0")
}
