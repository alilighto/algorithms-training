package main

import "os"

func main() {
	str := ""
	c := false
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	if args == nil {
		os.Stdout.WriteString("\n")
	}
	for i, v := range args[0] {
		if v != ' ' {
			c = true
			str += string(v)
		}
		if v == ' ' && c && i < len(args[0])-1 && args[0][i+1] != ' ' {
			str += "   "
			c = false
		}
	}
	os.Stdout.WriteString(str+"\n")
}
