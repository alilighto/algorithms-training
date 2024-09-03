package main

import "os"

func main() {
	arg := os.Args[1:]
	str := ""
	for _, v := range arg[0] {
		if string(v) == arg[1] {
			str += arg[2]
		} else {
			str += string(v)
		}
	}
	os.Stdout.WriteString(str+"\n")
}
